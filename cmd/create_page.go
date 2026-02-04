package cmd

import (
	"context"
	"fmt"
	"strconv"

	"github.com/alexhokl/confluence-cli/swagger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var createPageTitle string
var createPageSpaceID int64
var createPageSpaceKey string
var createPageParentID int64

var createPageCmd = &cobra.Command{
	Use:   "page",
	Short: "Create a new page",
	Long: `Create a new Confluence page with the specified title.

The page is created as an empty published page. You must specify the space
where the page will be created, either by space ID or space key.

Examples:
  # Create a page in a space by space ID
  confluence-cli create page --title "My New Page" --space-id 12345

  # Create a page in a space by space key
  confluence-cli create page --title "My New Page" --space-key MYSPACE

  # Create a page under a specific parent page
  confluence-cli create page --title "Child Page" --space-key MYSPACE --parent-id 67890`,
	RunE: runCreatePage,
}

func init() {
	createCmd.AddCommand(createPageCmd)
	createPageCmd.Flags().StringVarP(&createPageTitle, "title", "t", "", "Page title (required)")
	createPageCmd.Flags().Int64VarP(&createPageSpaceID, "space-id", "", 0, "Space ID")
	createPageCmd.Flags().StringVarP(&createPageSpaceKey, "space-key", "s", "", "Space key")
	createPageCmd.Flags().Int64VarP(&createPageParentID, "parent-id", "p", 0, "Parent page ID (optional)")
	createPageCmd.MarkFlagRequired("title")
}

func runCreatePage(_ *cobra.Command, _ []string) error {
	if createPageTitle == "" {
		return fmt.Errorf("page title is required")
	}

	if createPageSpaceID == 0 && createPageSpaceKey == "" {
		return fmt.Errorf("either --space-id or --space-key must be specified")
	}

	if createPageSpaceID != 0 && createPageSpaceKey != "" {
		return fmt.Errorf("only one of --space-id or --space-key should be specified, not both")
	}

	ctx := context.Background()
	client := newAPIClient()

	ctx = context.WithValue(ctx, swagger.ContextBasicAuth, swagger.BasicAuth{
		UserName: viper.GetString("email"),
		Password: viper.GetString("api_key"),
	})

	// Resolve space ID from space key if needed
	spaceID := createPageSpaceID
	if createPageSpaceKey != "" {
		resolvedID, err := getSpaceIDByKey(ctx, client, createPageSpaceKey)
		if err != nil {
			return err
		}
		spaceID = resolvedID
	}

	// Create the page request
	createRequest := swagger.NewCreatePageRequest(strconv.FormatInt(spaceID, 10))
	createRequest.SetTitle(createPageTitle)
	createRequest.SetStatus("current") // Published page

	// Set parent ID if specified
	if createPageParentID != 0 {
		createRequest.SetParentId(strconv.FormatInt(createPageParentID, 10))
	}

	// Set empty body in storage format
	bodyWrite := swagger.NewPageBodyWrite()
	bodyWrite.SetRepresentation("storage")
	bodyWrite.SetValue("")
	createRequest.SetBody(swagger.PageBodyWriteAsCreatePageRequestBody(bodyWrite))

	// Create the page
	result, _, err := client.PageAPI.CreatePage(ctx).
		CreatePageRequest(*createRequest).
		Execute()
	if err != nil {
		return fmt.Errorf("failed to create page: %w", err)
	}

	pageID := result.GetId()
	fmt.Printf("Page '%s' created successfully (ID: %s)\n", createPageTitle, pageID)
	return nil
}
