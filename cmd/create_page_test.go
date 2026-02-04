package cmd

import (
	"testing"
)

func TestCreatePageFlags(t *testing.T) {
	// Reset flags to default values for testing
	// Note: In a real test, we'd want to reset these after each test
	originalTitle := createPageTitle
	originalSpaceID := createPageSpaceID
	originalSpaceKey := createPageSpaceKey
	originalParentID := createPageParentID

	defer func() {
		createPageTitle = originalTitle
		createPageSpaceID = originalSpaceID
		createPageSpaceKey = originalSpaceKey
		createPageParentID = originalParentID
	}()

	// Test that command has expected flags
	cmd := createPageCmd

	// Check title flag exists
	titleFlag := cmd.Flags().Lookup("title")
	if titleFlag == nil {
		t.Error("expected 'title' flag to exist")
	}
	if titleFlag.Shorthand != "t" {
		t.Errorf("expected 'title' flag shorthand to be 't', got '%s'", titleFlag.Shorthand)
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

	// Check parent-id flag exists
	parentIDFlag := cmd.Flags().Lookup("parent-id")
	if parentIDFlag == nil {
		t.Error("expected 'parent-id' flag to exist")
	}
	if parentIDFlag.Shorthand != "p" {
		t.Errorf("expected 'parent-id' flag shorthand to be 'p', got '%s'", parentIDFlag.Shorthand)
	}
}

func TestRunCreatePageValidation(t *testing.T) {
	// Save original values
	originalTitle := createPageTitle
	originalSpaceID := createPageSpaceID
	originalSpaceKey := createPageSpaceKey

	defer func() {
		createPageTitle = originalTitle
		createPageSpaceID = originalSpaceID
		createPageSpaceKey = originalSpaceKey
	}()

	tests := []struct {
		name        string
		title       string
		spaceID     int64
		spaceKey    string
		expectError bool
		errorMsg    string
	}{
		{
			name:        "missing title",
			title:       "",
			spaceID:     12345,
			spaceKey:    "",
			expectError: true,
			errorMsg:    "page title is required",
		},
		{
			name:        "missing space identifier",
			title:       "Test Page",
			spaceID:     0,
			spaceKey:    "",
			expectError: true,
			errorMsg:    "either --space-id or --space-key must be specified",
		},
		{
			name:        "both space identifiers provided",
			title:       "Test Page",
			spaceID:     12345,
			spaceKey:    "MYSPACE",
			expectError: true,
			errorMsg:    "only one of --space-id or --space-key should be specified, not both",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createPageTitle = tt.title
			createPageSpaceID = tt.spaceID
			createPageSpaceKey = tt.spaceKey

			err := runCreatePage(nil, nil)

			if tt.expectError {
				if err == nil {
					t.Errorf("expected error containing '%s', got nil", tt.errorMsg)
					return
				}
				if err.Error() != tt.errorMsg {
					t.Errorf("expected error '%s', got '%s'", tt.errorMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
			}
		})
	}
}

func TestCreatePageCommandMetadata(t *testing.T) {
	cmd := createPageCmd

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
}
