package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var createAttachmentPageID int64
var createAttachmentFile string
var createAttachmentComment string

var createAttachmentCmd = &cobra.Command{
	Use:   "attachment",
	Short: "Upload a file as an attachment of a page",
	Long: `Upload a file as an attachment of a Confluence page.

The file name becomes the attachment name. If the page already has an
attachment with the same name, a new version of it is created.

An uploaded image can be shown on the page by referencing its file name in
markdown, e.g. ![diagram](diagram.png), via "update page".

Examples:
  # Upload an image to page with ID 12345
  confluence-cli create attachment --page-id 12345 --file diagram.png

  # Upload with a version comment
  confluence-cli create attachment --page-id 12345 --file diagram.png --comment "Updated diagram"`,
	RunE: runCreateAttachment,
}

func init() {
	createCmd.AddCommand(createAttachmentCmd)
	createAttachmentCmd.Flags().Int64VarP(&createAttachmentPageID, "page-id", "i", 0, "Page ID (required)")
	createAttachmentCmd.Flags().StringVarP(&createAttachmentFile, "file", "f", "", "Path of the file to upload (required)")
	createAttachmentCmd.Flags().StringVarP(&createAttachmentComment, "comment", "m", "", "Comment of the attachment version")
	createAttachmentCmd.MarkFlagRequired("page-id")
	createAttachmentCmd.MarkFlagRequired("file")
}

// attachmentUploadResult holds the attachment returned by the upload API.
type attachmentUploadResult struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func runCreateAttachment(_ *cobra.Command, _ []string) error {
	if createAttachmentPageID == 0 {
		return fmt.Errorf("page ID is required")
	}

	if createAttachmentFile == "" {
		return fmt.Errorf("file is required")
	}

	info, err := os.Stat(createAttachmentFile)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}
	if info.IsDir() {
		return fmt.Errorf("file %s is a directory", createAttachmentFile)
	}

	baseURL := fmt.Sprintf("https://%s.atlassian.net/wiki", viper.GetString("organization"))

	result, err := uploadAttachment(
		context.Background(),
		&http.Client{},
		baseURL,
		viper.GetString("email"),
		viper.GetString("api_key"),
		createAttachmentPageID,
		createAttachmentFile,
		createAttachmentComment,
	)
	if err != nil {
		return err
	}

	fmt.Printf("Attachment '%s' uploaded to page %d (ID: %s)\n", result.Title, createAttachmentPageID, result.ID)
	return nil
}

// uploadAttachment uploads a file to a page via REST API v1, as API v2 has no
// endpoint for it. PUT creates the attachment or a new version of an
// existing attachment with the same name.
func uploadAttachment(
	ctx context.Context,
	httpClient *http.Client,
	baseURL, email, apiKey string,
	pageID int64,
	filePath, comment string,
) (attachmentUploadResult, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return attachmentUploadResult{}, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return attachmentUploadResult{}, fmt.Errorf("failed to create form file: %w", err)
	}
	if _, err := io.Copy(part, file); err != nil {
		return attachmentUploadResult{}, fmt.Errorf("failed to read file: %w", err)
	}
	if comment != "" {
		if err := writer.WriteField("comment", comment); err != nil {
			return attachmentUploadResult{}, fmt.Errorf("failed to write comment: %w", err)
		}
	}
	if err := writer.WriteField("minorEdit", "true"); err != nil {
		return attachmentUploadResult{}, fmt.Errorf("failed to write minorEdit: %w", err)
	}
	if err := writer.Close(); err != nil {
		return attachmentUploadResult{}, fmt.Errorf("failed to close multipart writer: %w", err)
	}

	url := fmt.Sprintf("%s/rest/api/content/%d/child/attachment", baseURL, pageID)
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, &body)
	if err != nil {
		return attachmentUploadResult{}, fmt.Errorf("failed to create request: %w", err)
	}
	req.SetBasicAuth(email, apiKey)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("X-Atlassian-Token", "no-check")
	req.Header.Set("Accept", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return attachmentUploadResult{}, fmt.Errorf("failed to upload attachment: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return attachmentUploadResult{}, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return attachmentUploadResult{}, fmt.Errorf("failed to upload attachment: status %d: %s", resp.StatusCode, string(respBody))
	}

	var parsed struct {
		Results []attachmentUploadResult `json:"results"`
	}
	if err := json.Unmarshal(respBody, &parsed); err != nil {
		return attachmentUploadResult{}, fmt.Errorf("failed to parse response: %w", err)
	}
	if len(parsed.Results) == 0 {
		return attachmentUploadResult{}, fmt.Errorf("no attachment returned in response")
	}

	return parsed.Results[0], nil
}
