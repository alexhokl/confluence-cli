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
	// Verify default flag values
	if getPageID != 0 {
		t.Errorf("expected default page ID of 0, got %d", getPageID)
	}
	if getPageFormat != "storage" {
		t.Errorf("expected default format of 'storage', got %s", getPageFormat)
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
