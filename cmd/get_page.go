package cmd

import (
	"context"
	"fmt"
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

var getPageID int64
var getPageFormat string

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
markdown or formatting that was used when creating the page.`,
	RunE: runGetPage,
}

func init() {
	getCmd.AddCommand(getPageCmd)
	getPageCmd.Flags().Int64VarP(&getPageID, "id", "i", 0, "Page ID (required)")
	getPageCmd.Flags().StringVarP(&getPageFormat, "format", "f", "storage", "Content format (storage, atlas_doc_format, view, md)")
	getPageCmd.MarkFlagRequired("id")
}

func runGetPage(_ *cobra.Command, _ []string) error {
	if getPageID == 0 {
		return fmt.Errorf("page ID is required")
	}

	ctx := context.Background()
	client := newAPIClient()

	ctx = context.WithValue(ctx, swagger.ContextBasicAuth, swagger.BasicAuth{
		UserName: viper.GetString("email"),
		Password: viper.GetString("api_key"),
	})

	// Map format string to API enum
	bodyFormat, err := parseBodyFormat(getPageFormat)
	if err != nil {
		return err
	}

	result, _, err := client.PageAPI.GetPageById(ctx, getPageID).
		BodyFormat(bodyFormat).
		Execute()
	if err != nil {
		return fmt.Errorf("failed to get page: %w", err)
	}

	// Extract and print the body content
	content, err := extractPageContent(result, getPageFormat)
	if err != nil {
		return err
	}

	fmt.Println(content)
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
