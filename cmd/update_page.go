package cmd

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"

	"github.com/alexhokl/confluence-cli/swagger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
)

var updatePageID int64
var updatePageMessage string

var updatePageCmd = &cobra.Command{
	Use:   "page",
	Short: "Update page content",
	Long: `Update a Confluence page by its ID using your default editor.

This command fetches the page content, converts it to GitHub-flavored Markdown,
opens it in your default editor (specified by $EDITOR environment variable),
and upon saving, converts the markdown back to Confluence storage format and
updates the page if changes were made.

The editor is determined by the $EDITOR environment variable. If not set,
the command will fail with an error.

Example:
  # Update page with ID 12345
  confluence-cli update page --id 12345

  # Update page with a version message
  confluence-cli update page --id 12345 --message "Fixed typos"`,
	RunE: runUpdatePage,
}

func init() {
	updateCmd.AddCommand(updatePageCmd)
	updatePageCmd.Flags().Int64VarP(&updatePageID, "id", "i", 0, "Page ID (required)")
	updatePageCmd.Flags().StringVarP(&updatePageMessage, "message", "m", "", "Version message for the update")
	updatePageCmd.MarkFlagRequired("id")
}

func runUpdatePage(_ *cobra.Command, _ []string) error {
	if updatePageID == 0 {
		return fmt.Errorf("page ID is required")
	}

	// Check for $EDITOR environment variable
	editor := os.Getenv("EDITOR")
	if editor == "" {
		return fmt.Errorf("$EDITOR environment variable is not set")
	}

	ctx := context.Background()
	client := newAPIClient()

	ctx = context.WithValue(ctx, swagger.ContextBasicAuth, swagger.BasicAuth{
		UserName: viper.GetString("email"),
		Password: viper.GetString("api_key"),
	})

	// Fetch the page
	page, _, err := client.PageAPI.GetPageById(ctx, updatePageID).
		BodyFormat(swagger.PRIMARY_BODY_REPRESENTATION_SINGLE_STORAGE).
		Execute()
	if err != nil {
		return fmt.Errorf("failed to get page: %w", err)
	}

	// Extract page metadata
	pageTitle := page.GetTitle()
	pageStatus := page.GetStatus()
	pageVersion := page.GetVersion()
	currentVersionNumber := pageVersion.GetNumber()

	// Extract storage content
	body, ok := page.GetBodyOk()
	if !ok || body == nil {
		return fmt.Errorf("page has no body content")
	}

	if !body.HasStorage() {
		return fmt.Errorf("page storage content not available")
	}

	storageBody := body.GetStorage()
	originalStorage := storageBody.GetValue()

	// Convert storage format to markdown
	originalMarkdown, err := convertStorageToMarkdown(originalStorage)
	if err != nil {
		return fmt.Errorf("failed to convert page to markdown: %w", err)
	}

	// Open editor and get modified content
	modifiedMarkdown, err := openEditor(editor, originalMarkdown)
	if err != nil {
		return fmt.Errorf("failed to edit content: %w", err)
	}

	// Check if content was modified
	if originalMarkdown == modifiedMarkdown {
		fmt.Println("No changes detected, page not updated")
		return nil
	}

	// Convert markdown back to storage format
	newStorage, err := convertMarkdownToStorage(modifiedMarkdown)
	if err != nil {
		return fmt.Errorf("failed to convert markdown to storage format: %w", err)
	}

	// Prepare update request
	newVersion := int32(currentVersionNumber + 1)
	version := swagger.NewUpdatePageRequestVersion()
	version.SetNumber(newVersion)
	if updatePageMessage != "" {
		version.SetMessage(updatePageMessage)
	}

	bodyWrite := swagger.NewPageBodyWrite()
	bodyWrite.SetRepresentation("storage")
	bodyWrite.SetValue(newStorage)

	updateRequest := swagger.NewUpdatePageRequest(
		strconv.FormatInt(updatePageID, 10),
		string(pageStatus),
		pageTitle,
		swagger.PageBodyWriteAsCreatePageRequestBody(bodyWrite),
		*version,
	)

	// Update the page
	_, _, err = client.PageAPI.UpdatePage(ctx, updatePageID).
		UpdatePageRequest(*updateRequest).
		Execute()
	if err != nil {
		return fmt.Errorf("failed to update page: %w", err)
	}

	fmt.Printf("Page '%s' (ID: %d) updated successfully (version %d)\n", pageTitle, updatePageID, newVersion)
	return nil
}

// openEditor opens the specified editor with the content and returns the modified content.
func openEditor(editor, content string) (string, error) {
	// Create a temporary file
	tmpFile, err := os.CreateTemp("", "confluence-page-*.md")
	if err != nil {
		return "", fmt.Errorf("failed to create temporary file: %w", err)
	}
	tmpPath := tmpFile.Name()
	defer os.Remove(tmpPath)

	// Write content to temporary file
	if _, err := tmpFile.WriteString(content); err != nil {
		tmpFile.Close()
		return "", fmt.Errorf("failed to write to temporary file: %w", err)
	}
	tmpFile.Close()

	// Open editor
	cmd := exec.Command(editor, tmpPath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("editor exited with error: %w", err)
	}

	// Read modified content
	modifiedContent, err := os.ReadFile(tmpPath)
	if err != nil {
		return "", fmt.Errorf("failed to read modified content: %w", err)
	}

	return string(modifiedContent), nil
}

// htmlCodeBlockPattern matches HTML code blocks produced by goldmark.
// Captures an optional language class and the code content.
// Examples:
//
//	<pre><code class="language-sql">SELECT 1;\n</code></pre>
//	<pre><code>func main() {}\n</code></pre>
var htmlCodeBlockPattern = regexp.MustCompile(`(?s)<pre><code(?:\s+class="language-([^"]+)")?>(.+?)</code></pre>`)

// htmlImagePattern matches XHTML self-closing img tags produced by goldmark.
// Captures the src and alt attribute values.
// Examples:
//
//	<img src="screenshot.png" alt="caption" />
//	<img src="diagram.png" alt="" />
var htmlImagePattern = regexp.MustCompile(`<img\s+src="([^"]+)"\s+alt="([^"]*?)"\s*/>`)

// convertMarkdownToStorage converts GitHub-flavored Markdown to Confluence storage format (XHTML).
func convertMarkdownToStorage(markdown string) (string, error) {
	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM, // GitHub Flavored Markdown (tables, strikethrough, autolinks, task lists)
		),
		goldmark.WithRendererOptions(
			html.WithXHTML(), // Output XHTML for Confluence compatibility
		),
	)

	var buf bytes.Buffer
	if err := md.Convert([]byte(markdown), &buf); err != nil {
		return "", fmt.Errorf("failed to convert markdown to HTML: %w", err)
	}

	// Post-process: convert <pre><code> blocks to Confluence code macros so that
	// code is rendered with syntax highlighting in Confluence instead of as plain
	// preformatted text.
	result := postprocessStorageContent(buf.String())

	return result, nil
}

// postprocessStorageContent converts standard HTML elements to their Confluence
// storage format equivalents:
//   - <pre><code> blocks → ac:structured-macro code blocks
//   - <img> tags → ac:image blocks with ri:attachment
func postprocessStorageContent(content string) string {
	// Convert <pre><code> blocks to Confluence code macros.
	content = htmlCodeBlockPattern.ReplaceAllStringFunc(content, func(match string) string {
		submatches := htmlCodeBlockPattern.FindStringSubmatch(match)
		if len(submatches) < 3 {
			return match
		}
		lang := submatches[1]
		code := submatches[2]

		if lang != "" {
			return fmt.Sprintf(
				`<ac:structured-macro ac:name="code"><ac:parameter ac:name="language">%s</ac:parameter><ac:plain-text-body><![CDATA[%s]]></ac:plain-text-body></ac:structured-macro>`,
				lang, code,
			)
		}
		return fmt.Sprintf(
			`<ac:structured-macro ac:name="code"><ac:plain-text-body><![CDATA[%s]]></ac:plain-text-body></ac:structured-macro>`,
			code,
		)
	})

	// Convert <img> tags to Confluence ac:image blocks with ri:attachment.
	// If alt is empty or equal to the filename (the default set during preprocessing),
	// omit ac:alt so Confluence treats it as having no caption.
	content = htmlImagePattern.ReplaceAllStringFunc(content, func(match string) string {
		submatches := htmlImagePattern.FindStringSubmatch(match)
		if len(submatches) < 3 {
			return match
		}
		src := submatches[1]
		alt := submatches[2]

		if alt != "" && alt != src {
			return fmt.Sprintf(
				`<ac:image ac:alt="%s"><ri:attachment ri:filename="%s"/></ac:image>`,
				alt, src,
			)
		}
		return fmt.Sprintf(
			`<ac:image><ri:attachment ri:filename="%s"/></ac:image>`,
			src,
		)
	})

	return content
}
