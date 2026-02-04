package cmd

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/alexhokl/confluence-cli/swagger"
)

func TestExtractCursor(t *testing.T) {
	tests := []struct {
		name           string
		url            string
		expectedCursor string
		expectError    bool
	}{
		{
			name:           "valid URL with cursor",
			url:            "https://example.atlassian.net/wiki/api/v2/spaces?cursor=abc123",
			expectedCursor: "abc123",
			expectError:    false,
		},
		{
			name:           "valid URL with empty cursor",
			url:            "https://example.atlassian.net/wiki/api/v2/spaces?cursor=",
			expectedCursor: "",
			expectError:    false,
		},
		{
			name:           "valid URL without cursor parameter",
			url:            "https://example.atlassian.net/wiki/api/v2/spaces?limit=25",
			expectedCursor: "",
			expectError:    false,
		},
		{
			name:           "valid URL with multiple parameters",
			url:            "https://example.atlassian.net/wiki/api/v2/spaces?limit=25&cursor=xyz789&status=current",
			expectedCursor: "xyz789",
			expectError:    false,
		},
		{
			name:           "URL with encoded cursor",
			url:            "https://example.atlassian.net/wiki/api/v2/spaces?cursor=abc%2B123%3D",
			expectedCursor: "abc+123=",
			expectError:    false,
		},
		{
			name:           "empty URL",
			url:            "",
			expectedCursor: "",
			expectError:    false,
		},
		{
			name:           "relative URL with cursor",
			url:            "/wiki/api/v2/spaces?cursor=relative123",
			expectedCursor: "relative123",
			expectError:    false,
		},
		{
			name:           "invalid URL",
			url:            "://invalid",
			expectedCursor: "",
			expectError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cursor, err := extractCursor(tt.url)

			if tt.expectError {
				if err == nil {
					t.Errorf("extractCursor(%q) expected error, got nil", tt.url)
				}
				return
			}

			if err != nil {
				t.Errorf("extractCursor(%q) unexpected error: %v", tt.url, err)
				return
			}

			if cursor != tt.expectedCursor {
				t.Errorf("extractCursor(%q) = %q, want %q", tt.url, cursor, tt.expectedCursor)
			}
		})
	}
}

func TestPrintSpaces(t *testing.T) {
	tests := []struct {
		name           string
		spaces         []swagger.SpaceBulk
		expectedOutput []string
	}{
		{
			name:           "empty spaces",
			spaces:         []swagger.SpaceBulk{},
			expectedOutput: []string{"KEY", "NAME", "TYPE", "STATUS"},
		},
		{
			name: "single space with all fields",
			spaces: func() []swagger.SpaceBulk {
				space := swagger.NewSpaceBulk()
				space.SetId("1")
				space.SetKey("TEST")
				space.SetName("Test Space")
				space.SetType(swagger.SPACE_TYPE_GLOBAL)
				space.SetStatus(swagger.CURRENT)
				return []swagger.SpaceBulk{*space}
			}(),
			expectedOutput: []string{"TEST", "Test Space", "global", "current"},
		},
		{
			name: "space without optional fields",
			spaces: func() []swagger.SpaceBulk {
				space := swagger.NewSpaceBulk()
				space.SetId("2")
				space.SetKey("MINIMAL")
				space.SetName("Minimal Space")
				return []swagger.SpaceBulk{*space}
			}(),
			expectedOutput: []string{"MINIMAL", "Minimal Space"},
		},
		{
			name: "multiple spaces",
			spaces: func() []swagger.SpaceBulk {
				space1 := swagger.NewSpaceBulk()
				space1.SetId("1")
				space1.SetKey("SPACE1")
				space1.SetName("First Space")
				space1.SetType(swagger.SPACE_TYPE_GLOBAL)
				space1.SetStatus(swagger.CURRENT)

				space2 := swagger.NewSpaceBulk()
				space2.SetId("2")
				space2.SetKey("SPACE2")
				space2.SetName("Second Space")
				space2.SetType(swagger.SPACE_TYPE_PERSONAL)
				space2.SetStatus(swagger.ARCHIVED)

				return []swagger.SpaceBulk{*space1, *space2}
			}(),
			expectedOutput: []string{"SPACE1", "First Space", "global", "current", "SPACE2", "Second Space", "personal", "archived"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture stdout
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			printSpaces(tt.spaces)

			w.Close()
			os.Stdout = oldStdout

			var buf bytes.Buffer
			io.Copy(&buf, r)
			output := buf.String()

			for _, expected := range tt.expectedOutput {
				if !strings.Contains(output, expected) {
					t.Errorf("expected output to contain '%s', got:\n%s", expected, output)
				}
			}
		})
	}
}

func TestListSpacesFlags(t *testing.T) {
	cmd := listSpacesCmd

	// Check limit flag exists
	limitFlag := cmd.Flags().Lookup("limit")
	if limitFlag == nil {
		t.Error("expected 'limit' flag to exist")
	}
	if limitFlag.Shorthand != "l" {
		t.Errorf("expected 'limit' flag shorthand to be 'l', got '%s'", limitFlag.Shorthand)
	}
	if limitFlag.DefValue != "25" {
		t.Errorf("expected 'limit' flag default to be '25', got '%s'", limitFlag.DefValue)
	}

	// Check type flag exists
	typeFlag := cmd.Flags().Lookup("type")
	if typeFlag == nil {
		t.Error("expected 'type' flag to exist")
	}
	if typeFlag.Shorthand != "t" {
		t.Errorf("expected 'type' flag shorthand to be 't', got '%s'", typeFlag.Shorthand)
	}

	// Check status flag exists
	statusFlag := cmd.Flags().Lookup("status")
	if statusFlag == nil {
		t.Error("expected 'status' flag to exist")
	}
	if statusFlag.Shorthand != "s" {
		t.Errorf("expected 'status' flag shorthand to be 's', got '%s'", statusFlag.Shorthand)
	}
}

func TestListSpacesCommandMetadata(t *testing.T) {
	cmd := listSpacesCmd

	if cmd.Use != "spaces" {
		t.Errorf("expected Use to be 'spaces', got '%s'", cmd.Use)
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
}
