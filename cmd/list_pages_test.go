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
