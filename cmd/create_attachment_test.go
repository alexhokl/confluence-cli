package cmd

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCreateAttachmentFlags(t *testing.T) {
	cmd := createAttachmentCmd

	flags := []struct {
		name      string
		shorthand string
	}{
		{name: "page-id", shorthand: "i"},
		{name: "file", shorthand: "f"},
		{name: "comment", shorthand: "m"},
	}

	for _, tt := range flags {
		t.Run(tt.name, func(t *testing.T) {
			flag := cmd.Flags().Lookup(tt.name)
			if flag == nil {
				t.Fatalf("expected '%s' flag to exist", tt.name)
			}
			if flag.Shorthand != tt.shorthand {
				t.Errorf("expected '%s' flag shorthand to be '%s', got '%s'", tt.name, tt.shorthand, flag.Shorthand)
			}
		})
	}
}

func TestRunCreateAttachmentValidation(t *testing.T) {
	originalPageID := createAttachmentPageID
	originalFile := createAttachmentFile

	defer func() {
		createAttachmentPageID = originalPageID
		createAttachmentFile = originalFile
	}()

	dir := t.TempDir()

	tests := []struct {
		name     string
		pageID   int64
		file     string
		errorMsg string
	}{
		{
			name:     "missing page ID",
			pageID:   0,
			file:     "diagram.png",
			errorMsg: "page ID is required",
		},
		{
			name:     "missing file",
			pageID:   12345,
			file:     "",
			errorMsg: "file is required",
		},
		{
			name:     "file does not exist",
			pageID:   12345,
			file:     filepath.Join(dir, "missing.png"),
			errorMsg: "failed to read file",
		},
		{
			name:     "file is a directory",
			pageID:   12345,
			file:     dir,
			errorMsg: "is a directory",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createAttachmentPageID = tt.pageID
			createAttachmentFile = tt.file

			err := runCreateAttachment(nil, nil)

			if err == nil {
				t.Fatalf("expected error containing '%s', got nil", tt.errorMsg)
			}
			if !strings.Contains(err.Error(), tt.errorMsg) {
				t.Errorf("expected error containing '%s', got '%s'", tt.errorMsg, err.Error())
			}
		})
	}
}

func TestUploadAttachment(t *testing.T) {
	filePath := filepath.Join(t.TempDir(), "diagram.png")
	if err := os.WriteFile(filePath, []byte("png-data"), 0o600); err != nil {
		t.Fatalf("failed to write test file: %v", err)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected method PUT, got %s", r.Method)
		}
		if r.URL.Path != "/rest/api/content/123/child/attachment" {
			t.Errorf("unexpected path %s", r.URL.Path)
		}
		if r.Header.Get("X-Atlassian-Token") != "no-check" {
			t.Errorf("expected X-Atlassian-Token header to be 'no-check', got '%s'", r.Header.Get("X-Atlassian-Token"))
		}
		user, pass, ok := r.BasicAuth()
		if !ok || user != "user@example.com" || pass != "secret" {
			t.Errorf("unexpected basic auth %q %q %v", user, pass, ok)
		}

		file, header, err := r.FormFile("file")
		if err != nil {
			t.Fatalf("failed to read form file: %v", err)
		}
		defer file.Close()
		if header.Filename != "diagram.png" {
			t.Errorf("expected file name 'diagram.png', got '%s'", header.Filename)
		}
		content, _ := io.ReadAll(file)
		if string(content) != "png-data" {
			t.Errorf("unexpected file content '%s'", string(content))
		}
		if r.FormValue("comment") != "Updated diagram" {
			t.Errorf("expected comment 'Updated diagram', got '%s'", r.FormValue("comment"))
		}
		if r.FormValue("minorEdit") != "true" {
			t.Errorf("expected minorEdit 'true', got '%s'", r.FormValue("minorEdit"))
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"results":[{"id":"att456","title":"diagram.png"}],"size":1}`))
	}))
	defer server.Close()

	result, err := uploadAttachment(context.Background(), server.Client(), server.URL, "user@example.com", "secret", 123, filePath, "Updated diagram")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.ID != "att456" {
		t.Errorf("expected ID 'att456', got '%s'", result.ID)
	}
	if result.Title != "diagram.png" {
		t.Errorf("expected title 'diagram.png', got '%s'", result.Title)
	}
}

func TestUploadAttachmentErrorStatus(t *testing.T) {
	filePath := filepath.Join(t.TempDir(), "diagram.png")
	if err := os.WriteFile(filePath, []byte("png-data"), 0o600); err != nil {
		t.Fatalf("failed to write test file: %v", err)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(`{"message":"not permitted"}`))
	}))
	defer server.Close()

	_, err := uploadAttachment(context.Background(), server.Client(), server.URL, "user@example.com", "secret", 123, filePath, "")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !strings.Contains(err.Error(), "status 403") {
		t.Errorf("expected error containing 'status 403', got '%s'", err.Error())
	}
}

func TestCreateAttachmentCommandMetadata(t *testing.T) {
	cmd := createAttachmentCmd

	if cmd.Use != "attachment" {
		t.Errorf("expected Use to be 'attachment', got '%s'", cmd.Use)
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
