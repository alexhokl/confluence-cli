package cmd

import (
	"os"
	"testing"
)

func TestSupportsKittyGraphics(t *testing.T) {
	// Save original environment variables
	origTermProgram := os.Getenv("TERM_PROGRAM")
	origTerm := os.Getenv("TERM")
	origGhosttyResources := os.Getenv("GHOSTTY_RESOURCES_DIR")
	origKittyWindowID := os.Getenv("KITTY_WINDOW_ID")

	// Restore environment after test
	defer func() {
		os.Setenv("TERM_PROGRAM", origTermProgram)
		os.Setenv("TERM", origTerm)
		os.Setenv("GHOSTTY_RESOURCES_DIR", origGhosttyResources)
		os.Setenv("KITTY_WINDOW_ID", origKittyWindowID)
	}()

	tests := []struct {
		name              string
		termProgram       string
		term              string
		ghosttyResources  string
		kittyWindowID     string
		expectedSupported bool
	}{
		{
			name:              "ghostty terminal",
			termProgram:       "ghostty",
			expectedSupported: true,
		},
		{
			name:              "kitty terminal via TERM_PROGRAM",
			termProgram:       "kitty",
			expectedSupported: true,
		},
		{
			name:              "wezterm terminal",
			termProgram:       "wezterm",
			expectedSupported: true,
		},
		{
			name:              "kitty via TERM variable",
			term:              "xterm-kitty",
			expectedSupported: true,
		},
		{
			name:              "ghostty via resources dir",
			ghosttyResources:  "/path/to/ghostty",
			expectedSupported: true,
		},
		{
			name:              "kitty via window ID",
			kittyWindowID:     "12345",
			expectedSupported: true,
		},
		{
			name:              "unsupported terminal",
			termProgram:       "Apple_Terminal",
			term:              "xterm-256color",
			expectedSupported: false,
		},
		{
			name:              "no terminal env vars",
			expectedSupported: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Clear all relevant env vars
			os.Unsetenv("TERM_PROGRAM")
			os.Unsetenv("TERM")
			os.Unsetenv("GHOSTTY_RESOURCES_DIR")
			os.Unsetenv("KITTY_WINDOW_ID")

			// Set test-specific env vars
			if tt.termProgram != "" {
				os.Setenv("TERM_PROGRAM", tt.termProgram)
			}
			if tt.term != "" {
				os.Setenv("TERM", tt.term)
			}
			if tt.ghosttyResources != "" {
				os.Setenv("GHOSTTY_RESOURCES_DIR", tt.ghosttyResources)
			}
			if tt.kittyWindowID != "" {
				os.Setenv("KITTY_WINDOW_ID", tt.kittyWindowID)
			}

			result := supportsKittyGraphics()
			if result != tt.expectedSupported {
				t.Errorf("supportsKittyGraphics() = %v, want %v", result, tt.expectedSupported)
			}
		})
	}
}

func TestIsImageMimeType(t *testing.T) {
	tests := []struct {
		mimeType string
		expected bool
	}{
		{"image/png", true},
		{"image/jpeg", true},
		{"image/jpg", true},
		{"image/gif", true},
		{"IMAGE/PNG", true},
		{"Image/Jpeg", true},
		{"image/webp", false},
		{"image/svg+xml", false},
		{"text/plain", false},
		{"application/pdf", false},
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.mimeType, func(t *testing.T) {
			result := isImageMimeType(tt.mimeType)
			if result != tt.expected {
				t.Errorf("isImageMimeType(%q) = %v, want %v", tt.mimeType, result, tt.expected)
			}
		})
	}
}

func TestAttachmentMapsInitialization(t *testing.T) {
	// Test that attachment maps can be created and used correctly
	maps := &attachmentMaps{
		byID:     make(map[string]attachmentInfo),
		byTitle:  make(map[string]attachmentInfo),
		byFileID: make(map[string]attachmentInfo),
	}

	testAttachment := attachmentInfo{
		id:           "123",
		mediaType:    "image/png",
		downloadLink: "/download/attachments/123/test.png",
		title:        "test.png",
		fileID:       "abc-123",
	}

	// Add to maps
	maps.byID[testAttachment.id] = testAttachment
	maps.byTitle[testAttachment.title] = testAttachment
	maps.byFileID[testAttachment.fileID] = testAttachment

	// Verify lookups work
	if info, found := maps.byID["123"]; !found || info.title != "test.png" {
		t.Error("expected to find attachment by ID")
	}
	if info, found := maps.byTitle["test.png"]; !found || info.id != "123" {
		t.Error("expected to find attachment by title")
	}
	if info, found := maps.byFileID["abc-123"]; !found || info.downloadLink != "/download/attachments/123/test.png" {
		t.Error("expected to find attachment by fileID")
	}

	// Verify non-existent lookups
	if _, found := maps.byID["nonexistent"]; found {
		t.Error("expected not to find nonexistent ID")
	}
}

func TestDisplayImageKittyWithInvalidData(t *testing.T) {
	// Test with invalid image data
	invalidData := []byte("not an image")
	err := displayImageKitty(invalidData)
	if err == nil {
		t.Error("displayImageKitty with invalid data expected error, got nil")
	}
}
