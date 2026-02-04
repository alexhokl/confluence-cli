package cmd

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var deleteSpaceKey string

var deleteSpaceCmd = &cobra.Command{
	Use:   "space",
	Short: "Delete a space",
	Long: `Delete a Confluence space by its key.

This operation is irreversible. The space and all its contents will be permanently deleted.

**Permissions required**: 'Admin' permission for the space.`,
	RunE: runDeleteSpace,
}

func init() {
	deleteCmd.AddCommand(deleteSpaceCmd)
	deleteSpaceCmd.Flags().StringVarP(&deleteSpaceKey, "key", "k", "", "Space key to delete (required)")
	deleteSpaceCmd.MarkFlagRequired("key")
}

func runDeleteSpace(_ *cobra.Command, _ []string) error {
	if deleteSpaceKey == "" {
		return fmt.Errorf("space key is required")
	}

	// Use the legacy v1 API as v2 API doesn't support space deletion
	org := viper.GetString("organization")
	url := fmt.Sprintf("https://%s.atlassian.net/wiki/rest/api/space/%s", org, deleteSpaceKey)

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.SetBasicAuth(viper.GetString("email"), viper.GetString("api_key"))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to delete space: %w", err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusAccepted, http.StatusNoContent:
		fmt.Printf("Space '%s' has been deleted successfully\n", deleteSpaceKey)
		return nil
	case http.StatusNotFound:
		return fmt.Errorf("space '%s' not found", deleteSpaceKey)
	case http.StatusUnauthorized:
		return fmt.Errorf("unauthorized: check your credentials")
	case http.StatusForbidden:
		return fmt.Errorf("forbidden: you don't have permission to delete space '%s'", deleteSpaceKey)
	default:
		return fmt.Errorf("failed to delete space: HTTP %d", resp.StatusCode)
	}
}
