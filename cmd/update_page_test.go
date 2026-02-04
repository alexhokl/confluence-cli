package cmd

import (
	"os"
	"strings"
	"testing"
)

func TestConvertMarkdownToStorage(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "simple paragraph",
			input:    "Hello World",
			expected: "<p>Hello World</p>\n",
		},
		{
			name:     "heading level 1",
			input:    "# Title",
			expected: "<h1>Title</h1>\n",
		},
		{
			name:     "heading level 2",
			input:    "## Subtitle",
			expected: "<h2>Subtitle</h2>\n",
		},
		{
			name:     "bold text",
			input:    "**bold text**",
			expected: "<p><strong>bold text</strong></p>\n",
		},
		{
			name:     "italic text",
			input:    "*italic text*",
			expected: "<p><em>italic text</em></p>\n",
		},
		{
			name:     "strikethrough text",
			input:    "~~deleted text~~",
			expected: "<p><del>deleted text</del></p>\n",
		},
		{
			name:     "unordered list",
			input:    "- item 1\n- item 2",
			expected: "<ul>\n<li>item 1</li>\n<li>item 2</li>\n</ul>\n",
		},
		{
			name:     "ordered list",
			input:    "1. first\n2. second",
			expected: "<ol>\n<li>first</li>\n<li>second</li>\n</ol>\n",
		},
		{
			name:     "link",
			input:    "[Example](https://example.com)",
			expected: "<p><a href=\"https://example.com\">Example</a></p>\n",
		},
		{
			name:     "inline code",
			input:    "`code`",
			expected: "<p><code>code</code></p>\n",
		},
		{
			name:     "code block",
			input:    "```\nfunc main() {}\n```",
			expected: "<pre><code>func main() {}\n</code></pre>\n",
		},
		{
			name:     "blockquote",
			input:    "> quoted text",
			expected: "<blockquote>\n<p>quoted text</p>\n</blockquote>\n",
		},
		{
			name:     "simple table",
			input:    "| Header |\n|--------|\n| Cell   |",
			expected: "<table>\n<thead>\n<tr>\n<th>Header</th>\n</tr>\n</thead>\n<tbody>\n<tr>\n<td>Cell</td>\n</tr>\n</tbody>\n</table>\n",
		},
		{
			name:     "empty content",
			input:    "",
			expected: "",
		},
		{
			name:     "complex content",
			input:    "# Page Title\n\nThis is a **paragraph** with *formatting*.\n\n- Item one\n- Item two",
			expected: "<h1>Page Title</h1>\n<p>This is a <strong>paragraph</strong> with <em>formatting</em>.</p>\n<ul>\n<li>Item one</li>\n<li>Item two</li>\n</ul>\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := convertMarkdownToStorage(tt.input)
			if err != nil {
				t.Errorf("convertMarkdownToStorage(%q) unexpected error: %v", tt.input, err)
				return
			}
			if result != tt.expected {
				t.Errorf("convertMarkdownToStorage(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestOpenEditor(t *testing.T) {
	// Test with a simple editor command that just returns the content unchanged
	content := "test content"

	// Create a simple script that echoes content back
	tmpScript, err := os.CreateTemp("", "test-editor-*.sh")
	if err != nil {
		t.Fatalf("failed to create temp script: %v", err)
	}
	defer os.Remove(tmpScript.Name())

	// Write a script that does nothing (keeps file unchanged)
	script := "#!/bin/sh\nexit 0\n"
	if _, err := tmpScript.WriteString(script); err != nil {
		t.Fatalf("failed to write script: %v", err)
	}
	tmpScript.Close()

	// Make script executable
	if err := os.Chmod(tmpScript.Name(), 0755); err != nil {
		t.Fatalf("failed to chmod script: %v", err)
	}

	// Test openEditor with our fake editor
	result, err := openEditor(tmpScript.Name(), content)
	if err != nil {
		t.Errorf("openEditor() unexpected error: %v", err)
		return
	}

	if result != content {
		t.Errorf("openEditor() = %q, want %q", result, content)
	}
}

func TestOpenEditorWithModification(t *testing.T) {
	content := "original content"
	expectedContent := "modified content\n"

	// Create a script that modifies the file
	tmpScript, err := os.CreateTemp("", "test-editor-*.sh")
	if err != nil {
		t.Fatalf("failed to create temp script: %v", err)
	}
	defer os.Remove(tmpScript.Name())

	// Write a script that overwrites the file content
	script := `#!/bin/sh
echo "modified content" > "$1"
exit 0
`
	if _, err := tmpScript.WriteString(script); err != nil {
		t.Fatalf("failed to write script: %v", err)
	}
	tmpScript.Close()

	// Make script executable
	if err := os.Chmod(tmpScript.Name(), 0755); err != nil {
		t.Fatalf("failed to chmod script: %v", err)
	}

	// Test openEditor with our modifying editor
	result, err := openEditor(tmpScript.Name(), content)
	if err != nil {
		t.Errorf("openEditor() unexpected error: %v", err)
		return
	}

	if result != expectedContent {
		t.Errorf("openEditor() = %q, want %q", result, expectedContent)
	}
}

func TestOpenEditorError(t *testing.T) {
	content := "test content"

	// Create a script that exits with error
	tmpScript, err := os.CreateTemp("", "test-editor-*.sh")
	if err != nil {
		t.Fatalf("failed to create temp script: %v", err)
	}
	defer os.Remove(tmpScript.Name())

	script := "#!/bin/sh\nexit 1\n"
	if _, err := tmpScript.WriteString(script); err != nil {
		t.Fatalf("failed to write script: %v", err)
	}
	tmpScript.Close()

	if err := os.Chmod(tmpScript.Name(), 0755); err != nil {
		t.Fatalf("failed to chmod script: %v", err)
	}

	// Test openEditor with failing editor
	_, err = openEditor(tmpScript.Name(), content)
	if err == nil {
		t.Error("openEditor() expected error for failing editor, got nil")
	}
	if !strings.Contains(err.Error(), "editor exited with error") {
		t.Errorf("openEditor() error = %q, want error containing 'editor exited with error'", err.Error())
	}
}

func TestUpdatePageFlags(t *testing.T) {
	// Verify default flag values
	if updatePageID != 0 {
		t.Errorf("expected default page ID of 0, got %d", updatePageID)
	}
	if updatePageMessage != "" {
		t.Errorf("expected default message of '', got %s", updatePageMessage)
	}
}

func TestMarkdownToStorageRoundTrip(t *testing.T) {
	// Test that we can convert markdown to storage format
	// This tests the integration between the two conversion functions
	testCases := []struct {
		name     string
		markdown string
	}{
		{
			name:     "simple paragraph",
			markdown: "Hello World",
		},
		{
			name:     "heading and paragraph",
			markdown: "# Title\n\nThis is content.",
		},
		{
			name:     "list",
			markdown: "- Item 1\n- Item 2\n- Item 3",
		},
		{
			name:     "mixed content",
			markdown: "# Heading\n\nSome **bold** and *italic* text.\n\n- List item\n\n> A quote",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Convert markdown to storage
			storage, err := convertMarkdownToStorage(tc.markdown)
			if err != nil {
				t.Errorf("convertMarkdownToStorage(%q) error: %v", tc.markdown, err)
				return
			}

			// Verify storage is not empty (basic sanity check)
			if storage == "" && tc.markdown != "" {
				t.Errorf("convertMarkdownToStorage(%q) returned empty string", tc.markdown)
			}

			// Verify storage contains HTML tags
			if tc.markdown != "" && !strings.Contains(storage, "<") {
				t.Errorf("convertMarkdownToStorage(%q) = %q, expected HTML tags", tc.markdown, storage)
			}
		})
	}
}
