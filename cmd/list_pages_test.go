package cmd

import (
	"testing"

	"github.com/alexhokl/confluence-cli/swagger"
)

func TestPrintPages(t *testing.T) {
	// Test that printPages doesn't panic with empty slice
	pages := []swagger.PageBulk{}
	printPages(pages)

	// Test with a page that has values
	id := "123"
	title := "Test Page"
	spaceId := "456"
	status := swagger.CONTENT_STATUS_CURRENT

	page := swagger.PageBulk{}
	page.SetId(id)
	page.SetTitle(title)
	page.SetSpaceId(spaceId)
	page.SetStatus(status)

	pages = append(pages, page)
	printPages(pages)
}

func TestListPagesFlags(t *testing.T) {
	// Verify default flag values
	if listPagesLimit != 25 {
		t.Errorf("expected default limit of 25, got %d", listPagesLimit)
	}
	if listPagesSpaceID != 0 {
		t.Errorf("expected default space ID of 0, got %d", listPagesSpaceID)
	}
	if listPagesSpaceKey != "" {
		t.Errorf("expected default space key of empty string, got %s", listPagesSpaceKey)
	}
	if listPagesStatus != "" {
		t.Errorf("expected default status of empty string, got %s", listPagesStatus)
	}
	if listPagesTitle != "" {
		t.Errorf("expected default title of empty string, got %s", listPagesTitle)
	}
}

func TestPrintPagesWithParentID(t *testing.T) {
	// Test with a page that has a parent ID
	page := swagger.PageBulk{}
	page.SetId("123")
	page.SetTitle("Child Page")
	page.SetSpaceId("456")
	page.SetStatus(swagger.CONTENT_STATUS_CURRENT)
	page.SetParentId("789")

	pages := []swagger.PageBulk{page}

	// Capture stdout (we just verify it doesn't panic)
	printPages(pages)
}

func TestPrintPagesMultipleWithDifferentFields(t *testing.T) {
	// Test with multiple pages with different fields set
	page1 := swagger.PageBulk{}
	page1.SetId("1")
	page1.SetTitle("Page Without Parent")
	page1.SetSpaceId("100")
	page1.SetStatus(swagger.CONTENT_STATUS_CURRENT)

	page2 := swagger.PageBulk{}
	page2.SetId("2")
	page2.SetTitle("Page With Parent")
	page2.SetSpaceId("100")
	page2.SetStatus(swagger.CONTENT_STATUS_ARCHIVED)
	page2.SetParentId("1")

	page3 := swagger.PageBulk{}
	page3.SetId("3")
	page3.SetTitle("Trashed Page")
	page3.SetSpaceId("100")
	page3.SetStatus(swagger.CONTENT_STATUS_TRASHED)
	page3.SetParentId("2")

	pages := []swagger.PageBulk{page1, page2, page3}
	printPages(pages)
}

func TestListPagesCommandMetadata(t *testing.T) {
	cmd := listPagesCmd

	if cmd.Use != "pages" {
		t.Errorf("expected Use to be 'pages', got '%s'", cmd.Use)
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

	// Check space-id flag exists
	spaceIDFlag := cmd.Flags().Lookup("space-id")
	if spaceIDFlag == nil {
		t.Error("expected 'space-id' flag to exist")
	}

	// Check space-key flag exists
	spaceKeyFlag := cmd.Flags().Lookup("space-key")
	if spaceKeyFlag == nil {
		t.Error("expected 'space-key' flag to exist")
	}
	if spaceKeyFlag.Shorthand != "s" {
		t.Errorf("expected 'space-key' flag shorthand to be 's', got '%s'", spaceKeyFlag.Shorthand)
	}

	// Check status flag exists
	statusFlag := cmd.Flags().Lookup("status")
	if statusFlag == nil {
		t.Error("expected 'status' flag to exist")
	}
	if statusFlag.Shorthand != "S" {
		t.Errorf("expected 'status' flag shorthand to be 'S', got '%s'", statusFlag.Shorthand)
	}

	// Check title flag exists
	titleFlag := cmd.Flags().Lookup("title")
	if titleFlag == nil {
		t.Error("expected 'title' flag to exist")
	}
	if titleFlag.Shorthand != "t" {
		t.Errorf("expected 'title' flag shorthand to be 't', got '%s'", titleFlag.Shorthand)
	}
}

func TestContainsIgnoreCase(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		substr   string
		expected bool
	}{
		{
			name:     "exact match",
			s:        "Hello World",
			substr:   "Hello World",
			expected: true,
		},
		{
			name:     "partial match at start",
			s:        "Hello World",
			substr:   "Hello",
			expected: true,
		},
		{
			name:     "partial match at end",
			s:        "Hello World",
			substr:   "World",
			expected: true,
		},
		{
			name:     "partial match in middle",
			s:        "Hello World",
			substr:   "lo Wo",
			expected: true,
		},
		{
			name:     "case insensitive match - lowercase search",
			s:        "Hello World",
			substr:   "hello",
			expected: true,
		},
		{
			name:     "case insensitive match - uppercase search",
			s:        "Hello World",
			substr:   "WORLD",
			expected: true,
		},
		{
			name:     "case insensitive match - mixed case",
			s:        "Hello World",
			substr:   "hElLo WoRlD",
			expected: true,
		},
		{
			name:     "no match",
			s:        "Hello World",
			substr:   "Goodbye",
			expected: false,
		},
		{
			name:     "empty substring matches everything",
			s:        "Hello World",
			substr:   "",
			expected: true,
		},
		{
			name:     "empty string with empty substring",
			s:        "",
			substr:   "",
			expected: true,
		},
		{
			name:     "empty string with non-empty substring",
			s:        "",
			substr:   "Hello",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := containsIgnoreCase(tt.s, tt.substr)
			if result != tt.expected {
				t.Errorf("containsIgnoreCase(%q, %q) = %v, want %v", tt.s, tt.substr, result, tt.expected)
			}
		})
	}
}
