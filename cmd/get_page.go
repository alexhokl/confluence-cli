package cmd

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/JohannesKaufmann/html-to-markdown/v2/converter"
	"github.com/JohannesKaufmann/html-to-markdown/v2/plugin/base"
	"github.com/JohannesKaufmann/html-to-markdown/v2/plugin/commonmark"
	"github.com/JohannesKaufmann/html-to-markdown/v2/plugin/strikethrough"
	"github.com/JohannesKaufmann/html-to-markdown/v2/plugin/table"

	"github.com/alexhokl/confluence-cli/swagger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type getPageOptions struct {
	id       int64
	format   string
	noImages bool
}

var getPageOpts = getPageOptions{}

var getPageCmd = &cobra.Command{
	Use:   "page",
	Short: "Get page content",
	Long: `Get the content of a Confluence page by its ID.

The content is returned in the specified format. Available formats:
  - storage: Confluence storage format (XHTML-based, default)
  - atlas_doc_format: Atlassian Document Format (JSON)
  - view: Rendered HTML view
  - md: GitHub-flavored Markdown (converted from storage format)

The storage format preserves the original content structure including any
markdown or formatting that was used when creating the page.

Examples:
  # Get page content in markdown format (images displayed by default in supported terminals)
  confluence-cli get page -i 12345 -f md

  # Get page content without inline images
  confluence-cli get page -i 12345 -f md --no-images

Note: Images are displayed inline by default in terminals that support
the Kitty graphics protocol (Ghostty, Kitty, WezTerm).`,
	RunE: runGetPage,
}

func init() {
	getCmd.AddCommand(getPageCmd)
	getPageCmd.Flags().Int64VarP(&getPageOpts.id, "id", "i", 0, "Page ID (required)")
	getPageCmd.Flags().StringVarP(&getPageOpts.format, "format", "f", "storage", "Content format (storage, atlas_doc_format, view, md)")
	getPageCmd.Flags().BoolVar(&getPageOpts.noImages, "no-images", false, "Do not display images inline")
	getPageCmd.MarkFlagRequired("id")
}

func runGetPage(_ *cobra.Command, _ []string) error {
	if getPageOpts.id == 0 {
		return fmt.Errorf("page ID is required")
	}

	ctx := context.Background()
	client := newAPIClient()

	ctx = context.WithValue(ctx, swagger.ContextBasicAuth, swagger.BasicAuth{
		UserName: viper.GetString("email"),
		Password: viper.GetString("api_key"),
	})

	// Map format string to API enum
	bodyFormat, err := parseBodyFormat(getPageOpts.format)
	if err != nil {
		return err
	}

	result, _, err := client.PageAPI.GetPageById(ctx, getPageOpts.id).
		BodyFormat(bodyFormat).
		Execute()
	if err != nil {
		return fmt.Errorf("failed to get page: %w", err)
	}

	// Check if we should show images (default: yes, unless --no-images or terminal doesn't support it)
	showImages := !getPageOpts.noImages && supportsKittyGraphics() && getPageOpts.format == "md"

	// Build attachment maps if showing images
	var attMaps *attachmentMaps
	if showImages {
		attMaps, err = buildAttachmentMaps(client, ctx, getPageOpts.id)
		if err != nil {
			// Non-fatal: just warn and continue without images
			fmt.Printf("Warning: could not fetch attachments: %v\n\n", err)
			showImages = false
		}
	}

	// Extract and print the body content
	content, err := extractPageContent(result, getPageOpts.format)
	if err != nil {
		return err
	}

	// If showing images, process content with inline image display
	if showImages && attMaps != nil {
		printContentWithImages(content, attMaps)
	} else {
		fmt.Println(content)
	}
	return nil
}

// parseBodyFormat converts a format string to the API enum type.
// For "md" format, we request storage format and convert it later.
func parseBodyFormat(format string) (swagger.PrimaryBodyRepresentationSingle, error) {
	switch format {
	case "storage", "md":
		return swagger.PRIMARY_BODY_REPRESENTATION_SINGLE_STORAGE, nil
	case "atlas_doc_format":
		return swagger.PRIMARY_BODY_REPRESENTATION_SINGLE_ATLAS_DOC_FORMAT, nil
	case "view":
		return swagger.PRIMARY_BODY_REPRESENTATION_SINGLE_VIEW, nil
	default:
		return "", fmt.Errorf("invalid format '%s': must be one of storage, atlas_doc_format, view, md", format)
	}
}

// extractPageContent extracts the body content from the API response based on the requested format.
func extractPageContent(page *swagger.CreatePage200Response, format string) (string, error) {
	if page == nil {
		return "", fmt.Errorf("page response is nil")
	}

	body, ok := page.GetBodyOk()
	if !ok || body == nil {
		return "", fmt.Errorf("page has no body content")
	}

	var bodyType *swagger.BodyType
	// For "md" format, we use storage format and convert it
	actualFormat := format
	if format == "md" {
		actualFormat = "storage"
	}

	switch actualFormat {
	case "storage":
		if body.HasStorage() {
			bt := body.GetStorage()
			bodyType = &bt
		}
	case "atlas_doc_format":
		if body.HasAtlasDocFormat() {
			bt := body.GetAtlasDocFormat()
			bodyType = &bt
		}
	case "view":
		if body.HasView() {
			bt := body.GetView()
			bodyType = &bt
		}
	}

	if bodyType == nil {
		return "", fmt.Errorf("content not available in '%s' format", actualFormat)
	}

	content := bodyType.GetValue()

	// Convert to markdown if requested
	if format == "md" {
		return convertStorageToMarkdown(content)
	}

	return content, nil
}

// convertStorageToMarkdown converts Confluence storage format (XHTML) to GitHub-flavored Markdown.
func convertStorageToMarkdown(content string) (string, error) {
	// Preprocess Confluence-specific tags to standard HTML
	content = preprocessConfluenceContent(content)

	conv := converter.NewConverter(
		converter.WithPlugins(
			base.NewBasePlugin(),
			commonmark.NewCommonmarkPlugin(),
			strikethrough.NewStrikethroughPlugin(),
			table.NewTablePlugin(),
		),
	)

	markdown, err := conv.ConvertString(content)
	if err != nil {
		return "", fmt.Errorf("failed to convert to markdown: %w", err)
	}

	// Trim leading/trailing whitespace
	return strings.TrimSpace(markdown), nil
}

// acImagePattern matches Confluence ac:image tags with ri:attachment
// Example: <ac:image ...><ri:attachment ri:filename="image.png" .../></ac:image>
var acImagePattern = regexp.MustCompile(`<ac:image[^>]*>.*?<ri:attachment[^>]*ri:filename="([^"]+)"[^>]*/?>.*?</ac:image>`)

// acImagePatternWithAlt matches ac:image with ac:alt attribute
var acImagePatternWithAlt = regexp.MustCompile(`<ac:image[^>]*ac:alt="([^"]*)"[^>]*>.*?<ri:attachment[^>]*ri:filename="([^"]+)"[^>]*/?>.*?</ac:image>`)

// preprocessConfluenceContent converts Confluence-specific markup to standard HTML.
func preprocessConfluenceContent(content string) string {
	// First try to match images with alt text
	content = acImagePatternWithAlt.ReplaceAllStringFunc(content, func(match string) string {
		submatches := acImagePatternWithAlt.FindStringSubmatch(match)
		if len(submatches) >= 3 {
			alt := submatches[1]
			filename := submatches[2]
			return fmt.Sprintf(`<img src="%s" alt="%s" />`, filename, alt)
		}
		return match
	})

	// Then match images without alt text (that weren't already replaced)
	content = acImagePattern.ReplaceAllStringFunc(content, func(match string) string {
		// Skip if already converted to img tag
		if strings.Contains(match, "<img") {
			return match
		}
		submatches := acImagePattern.FindStringSubmatch(match)
		if len(submatches) >= 2 {
			filename := submatches[1]
			return fmt.Sprintf(`<img src="%s" alt="%s" />`, filename, filename)
		}
		return match
	})

	return content
}

// imagePattern matches markdown image syntax: ![alt text](url)
var imagePattern = regexp.MustCompile(`!\[([^\]]*)\]\(([^)]+)\)`)

// printContentWithImages prints markdown content while displaying images inline.
// It detects markdown image references and attempts to display them using Kitty graphics.
func printContentWithImages(content string, attMaps *attachmentMaps) {
	lines := strings.Split(content, "\n")

	for _, line := range lines {
		// Find all image references in this line
		matches := imagePattern.FindAllStringSubmatchIndex(line, -1)

		if len(matches) == 0 {
			// No images in this line, print as-is
			fmt.Println(line)
			continue
		}

		// Process line with images
		lastEnd := 0
		for _, match := range matches {
			// Print text before the image
			if match[0] > lastEnd {
				fmt.Print(line[lastEnd:match[0]])
			}

			// Extract alt text and URL from the match
			altText := line[match[2]:match[3]]
			imageRef := line[match[4]:match[5]]

			// Try to find the attachment and display the image
			displayed := false

			// The image reference might be a filename or a path
			// Try to match by title (filename)
			filename := extractFilename(imageRef)

			if info, found := attMaps.byTitle[filename]; found {
				if isImageMimeType(info.mediaType) && info.downloadLink != "" {
					// Print any pending text first
					fmt.Println()

					// Try to display the image
					if err := fetchAndDisplayImage(info.downloadLink); err == nil {
						displayed = true
						// Print alt text as caption
						if altText != "" {
							fmt.Printf("[%s]\n", altText)
						}
					}
				}
			}

			if !displayed {
				// Couldn't display image, print the original markdown
				fmt.Print(line[match[0]:match[1]])
			}

			lastEnd = match[1]
		}

		// Print remaining text after the last image
		if lastEnd < len(line) {
			fmt.Print(line[lastEnd:])
		}
		fmt.Println()
	}
}

// extractFilename extracts the filename from a path or URL.
func extractFilename(path string) string {
	// Handle Confluence attachment URLs which might be in various formats
	// e.g., "/download/attachments/123/image.png" or just "image.png"

	// Split by / and get the last part
	parts := strings.Split(path, "/")
	filename := parts[len(parts)-1]

	// Remove any query parameters
	if idx := strings.Index(filename, "?"); idx != -1 {
		filename = filename[:idx]
	}

	return filename
}
