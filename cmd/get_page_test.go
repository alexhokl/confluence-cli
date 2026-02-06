package cmd

import (
	"testing"

	"github.com/alexhokl/confluence-cli/swagger"
)

func TestParseBodyFormat(t *testing.T) {
	tests := []struct {
		name        string
		format      string
		expected    swagger.PrimaryBodyRepresentationSingle
		expectError bool
	}{
		{
			name:        "storage format",
			format:      "storage",
			expected:    swagger.PRIMARY_BODY_REPRESENTATION_SINGLE_STORAGE,
			expectError: false,
		},
		{
			name:        "atlas_doc_format",
			format:      "atlas_doc_format",
			expected:    swagger.PRIMARY_BODY_REPRESENTATION_SINGLE_ATLAS_DOC_FORMAT,
			expectError: false,
		},
		{
			name:        "view format",
			format:      "view",
			expected:    swagger.PRIMARY_BODY_REPRESENTATION_SINGLE_VIEW,
			expectError: false,
		},
		{
			name:        "md format maps to storage",
			format:      "md",
			expected:    swagger.PRIMARY_BODY_REPRESENTATION_SINGLE_STORAGE,
			expectError: false,
		},
		{
			name:        "invalid format",
			format:      "invalid",
			expected:    "",
			expectError: true,
		},
		{
			name:        "empty format",
			format:      "",
			expected:    "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parseBodyFormat(tt.format)
			if tt.expectError {
				if err == nil {
					t.Errorf("parseBodyFormat(%q) expected error, got nil", tt.format)
				}
			} else {
				if err != nil {
					t.Errorf("parseBodyFormat(%q) unexpected error: %v", tt.format, err)
				}
				if result != tt.expected {
					t.Errorf("parseBodyFormat(%q) = %v, want %v", tt.format, result, tt.expected)
				}
			}
		})
	}
}

func TestExtractPageContent(t *testing.T) {
	// Test with nil page
	_, err := extractPageContent(nil, "storage")
	if err == nil {
		t.Error("extractPageContent(nil) expected error, got nil")
	}

	// Test with page that has no body
	pageNoBody := swagger.NewCreatePage200Response()
	_, err = extractPageContent(pageNoBody, "storage")
	if err == nil {
		t.Error("extractPageContent with no body expected error, got nil")
	}

	// Test with page that has storage content
	storageValue := "<p>Hello World</p>"
	storageBody := swagger.NewBodyType()
	storageBody.SetValue(storageValue)
	storageBody.SetRepresentation("storage")

	bodySingle := swagger.NewBodySingle()
	bodySingle.SetStorage(*storageBody)

	pageWithStorage := swagger.NewCreatePage200Response()
	pageWithStorage.SetBody(*bodySingle)

	content, err := extractPageContent(pageWithStorage, "storage")
	if err != nil {
		t.Errorf("extractPageContent with storage body unexpected error: %v", err)
	}
	if content != storageValue {
		t.Errorf("extractPageContent = %q, want %q", content, storageValue)
	}

	// Test requesting format that doesn't exist in response
	_, err = extractPageContent(pageWithStorage, "view")
	if err == nil {
		t.Error("extractPageContent requesting unavailable format expected error, got nil")
	}
}

func TestGetPageFlags(t *testing.T) {
	// Verify default flag values by checking the options struct
	// Note: default values are set by Cobra flags, not by the struct initialization
	// After init(), the flags will have default values configured

	// Check that the id flag has the expected default
	idFlag := getPageCmd.Flags().Lookup("id")
	if idFlag == nil {
		t.Error("expected 'id' flag to exist")
	}
	if idFlag.DefValue != "0" {
		t.Errorf("expected default page ID of 0, got %s", idFlag.DefValue)
	}

	// Check that the format flag has the expected default
	formatFlag := getPageCmd.Flags().Lookup("format")
	if formatFlag == nil {
		t.Error("expected 'format' flag to exist")
	}
	if formatFlag.DefValue != "storage" {
		t.Errorf("expected default format of 'storage', got %s", formatFlag.DefValue)
	}

	// Check that the no-images flag exists and has the expected default
	noImagesFlag := getPageCmd.Flags().Lookup("no-images")
	if noImagesFlag == nil {
		t.Error("expected 'no-images' flag to exist")
	}
	if noImagesFlag.DefValue != "false" {
		t.Errorf("expected default no-images of 'false', got %s", noImagesFlag.DefValue)
	}
}

func TestConvertStorageToMarkdown(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "simple paragraph",
			input:    "<p>Hello World</p>",
			expected: "Hello World",
		},
		{
			name:     "heading",
			input:    "<h1>Title</h1>",
			expected: "# Title",
		},
		{
			name:     "multiple headings",
			input:    "<h1>Title</h1><h2>Subtitle</h2>",
			expected: "# Title\n\n## Subtitle",
		},
		{
			name:     "bold text",
			input:    "<p><strong>bold text</strong></p>",
			expected: "**bold text**",
		},
		{
			name:     "italic text",
			input:    "<p><em>italic text</em></p>",
			expected: "*italic text*",
		},
		{
			name:     "strikethrough text",
			input:    "<p><del>deleted text</del></p>",
			expected: "~~deleted text~~",
		},
		{
			name:     "unordered list",
			input:    "<ul><li>item 1</li><li>item 2</li></ul>",
			expected: "- item 1\n- item 2",
		},
		{
			name:     "ordered list",
			input:    "<ol><li>first</li><li>second</li></ol>",
			expected: "1. first\n2. second",
		},
		{
			name:     "link",
			input:    `<p><a href="https://example.com">Example</a></p>`,
			expected: "[Example](https://example.com)",
		},
		{
			name:     "code inline",
			input:    "<p><code>code</code></p>",
			expected: "`code`",
		},
		{
			name:     "code block",
			input:    "<pre><code>func main() {}</code></pre>",
			expected: "```\nfunc main() {}\n```",
		},
		{
			name:     "simple table",
			input:    "<table><thead><tr><th>Header</th></tr></thead><tbody><tr><td>Cell</td></tr></tbody></table>",
			expected: "| Header |\n|--------|\n| Cell   |",
		},
		{
			name:     "blockquote",
			input:    "<blockquote><p>quoted text</p></blockquote>",
			expected: "> quoted text",
		},
		{
			name:     "empty content",
			input:    "",
			expected: "",
		},
		{
			name:     "whitespace only",
			input:    "   \n\t  ",
			expected: "",
		},
		{
			name:     "complex confluence content",
			input:    "<h1>Page Title</h1><p>This is a <strong>paragraph</strong> with <em>formatting</em>.</p><ul><li>Item one</li><li>Item two</li></ul>",
			expected: "# Page Title\n\nThis is a **paragraph** with *formatting*.\n\n- Item one\n- Item two",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := convertStorageToMarkdown(tt.input)
			if err != nil {
				t.Errorf("convertStorageToMarkdown(%q) unexpected error: %v", tt.input, err)
				return
			}
			if result != tt.expected {
				t.Errorf("convertStorageToMarkdown(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestExtractPageContentWithMarkdownFormat(t *testing.T) {
	// Test with page that has storage content, requested as markdown
	storageValue := "<h1>Title</h1><p>Hello World</p>"
	expectedMarkdown := "# Title\n\nHello World"

	storageBody := swagger.NewBodyType()
	storageBody.SetValue(storageValue)
	storageBody.SetRepresentation("storage")

	bodySingle := swagger.NewBodySingle()
	bodySingle.SetStorage(*storageBody)

	pageWithStorage := swagger.NewCreatePage200Response()
	pageWithStorage.SetBody(*bodySingle)

	content, err := extractPageContent(pageWithStorage, "md")
	if err != nil {
		t.Errorf("extractPageContent with md format unexpected error: %v", err)
	}
	if content != expectedMarkdown {
		t.Errorf("extractPageContent with md format = %q, want %q", content, expectedMarkdown)
	}
}

func TestExtractPageContentWithAtlasDocFormat(t *testing.T) {
	// Test with page that has atlas_doc_format content
	atlasDocValue := `{"type":"doc","content":[{"type":"paragraph","content":[{"type":"text","text":"Hello"}]}]}`

	atlasBody := swagger.NewBodyType()
	atlasBody.SetValue(atlasDocValue)
	atlasBody.SetRepresentation("atlas_doc_format")

	bodySingle := swagger.NewBodySingle()
	bodySingle.SetAtlasDocFormat(*atlasBody)

	pageWithAtlasDoc := swagger.NewCreatePage200Response()
	pageWithAtlasDoc.SetBody(*bodySingle)

	content, err := extractPageContent(pageWithAtlasDoc, "atlas_doc_format")
	if err != nil {
		t.Errorf("extractPageContent with atlas_doc_format unexpected error: %v", err)
	}
	if content != atlasDocValue {
		t.Errorf("extractPageContent with atlas_doc_format = %q, want %q", content, atlasDocValue)
	}
}

func TestExtractPageContentWithViewFormat(t *testing.T) {
	// Test with page that has view format content
	viewValue := "<div class='content-wrapper'><p>Hello World</p></div>"

	viewBody := swagger.NewBodyType()
	viewBody.SetValue(viewValue)
	viewBody.SetRepresentation("view")

	bodySingle := swagger.NewBodySingle()
	bodySingle.SetView(*viewBody)

	pageWithView := swagger.NewCreatePage200Response()
	pageWithView.SetBody(*bodySingle)

	content, err := extractPageContent(pageWithView, "view")
	if err != nil {
		t.Errorf("extractPageContent with view format unexpected error: %v", err)
	}
	if content != viewValue {
		t.Errorf("extractPageContent with view format = %q, want %q", content, viewValue)
	}
}

func TestGetPageCommandMetadata(t *testing.T) {
	cmd := getPageCmd

	if cmd.Use != "page" {
		t.Errorf("expected Use to be 'page', got '%s'", cmd.Use)
	}

	if cmd.Short == "" {
		t.Error("expected Short description to be set")
	}

	if cmd.Long == "" {
		t.Error("expected Long description to be set")
	}

	if cmd.RunE == nil {
		t.Error("expected RunE to be set")
	}

	// Check id flag exists and is required
	idFlag := cmd.Flags().Lookup("id")
	if idFlag == nil {
		t.Error("expected 'id' flag to exist")
	}
	if idFlag.Shorthand != "i" {
		t.Errorf("expected 'id' flag shorthand to be 'i', got '%s'", idFlag.Shorthand)
	}

	// Check format flag exists
	formatFlag := cmd.Flags().Lookup("format")
	if formatFlag == nil {
		t.Error("expected 'format' flag to exist")
	}
	if formatFlag.Shorthand != "f" {
		t.Errorf("expected 'format' flag shorthand to be 'f', got '%s'", formatFlag.Shorthand)
	}
	if formatFlag.DefValue != "storage" {
		t.Errorf("expected 'format' flag default to be 'storage', got '%s'", formatFlag.DefValue)
	}

	// Check no-images flag exists
	noImagesFlag := cmd.Flags().Lookup("no-images")
	if noImagesFlag == nil {
		t.Error("expected 'no-images' flag to exist")
	}
	if noImagesFlag.DefValue != "false" {
		t.Errorf("expected 'no-images' flag default to be 'false', got '%s'", noImagesFlag.DefValue)
	}
}

func TestExtractFilename(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "simple filename",
			input:    "image.png",
			expected: "image.png",
		},
		{
			name:     "path with slashes",
			input:    "/download/attachments/123/image.png",
			expected: "image.png",
		},
		{
			name:     "path with query parameters",
			input:    "image.png?version=1&modificationDate=123",
			expected: "image.png",
		},
		{
			name:     "full URL-like path with query",
			input:    "/wiki/download/attachments/12345/screenshot.jpg?api=v2",
			expected: "screenshot.jpg",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "just a slash",
			input:    "/",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractFilename(tt.input)
			if result != tt.expected {
				t.Errorf("extractFilename(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestImagePattern(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		shouldMatch bool
		altText     string
		imageRef    string
	}{
		{
			name:        "simple image",
			input:       "![alt text](image.png)",
			shouldMatch: true,
			altText:     "alt text",
			imageRef:    "image.png",
		},
		{
			name:        "image with path",
			input:       "![screenshot](/attachments/screenshot.jpg)",
			shouldMatch: true,
			altText:     "screenshot",
			imageRef:    "/attachments/screenshot.jpg",
		},
		{
			name:        "image with empty alt",
			input:       "![](image.png)",
			shouldMatch: true,
			altText:     "",
			imageRef:    "image.png",
		},
		{
			name:        "not an image",
			input:       "[link text](url)",
			shouldMatch: false,
		},
		{
			name:        "plain text",
			input:       "Hello World",
			shouldMatch: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matches := imagePattern.FindStringSubmatch(tt.input)
			if tt.shouldMatch {
				if matches == nil {
					t.Errorf("imagePattern.FindStringSubmatch(%q) expected match, got nil", tt.input)
					return
				}
				if len(matches) < 3 {
					t.Errorf("imagePattern.FindStringSubmatch(%q) expected 3 groups, got %d", tt.input, len(matches))
					return
				}
				if matches[1] != tt.altText {
					t.Errorf("alt text = %q, want %q", matches[1], tt.altText)
				}
				if matches[2] != tt.imageRef {
					t.Errorf("image ref = %q, want %q", matches[2], tt.imageRef)
				}
			} else {
				if matches != nil {
					t.Errorf("imagePattern.FindStringSubmatch(%q) expected no match, got %v", tt.input, matches)
				}
			}
		})
	}
}
