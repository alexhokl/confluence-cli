package cmd

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/alexhokl/confluence-cli/swagger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listPagesLimit int32
var listPagesSpaceID int64
var listPagesSpaceKey string
var listPagesStatus string
var listPagesTitle string

var listPagesCmd = &cobra.Command{
	Use:   "pages",
	Short: "List all pages",
	Long: `List all Confluence pages that the user has permission to view.

If --space-key or --space-id is specified, only pages in that space will be returned.
Otherwise, pages from all accessible spaces will be returned.

Note: --space-key takes precedence over --space-id if both are specified.`,
	RunE: runListPages,
}

func init() {
	listCmd.AddCommand(listPagesCmd)
	listPagesCmd.Flags().Int32VarP(&listPagesLimit, "limit", "l", 25, "Maximum number of pages per page")
	listPagesCmd.Flags().Int64Var(&listPagesSpaceID, "space-id", 0, "Filter by space ID")
	listPagesCmd.Flags().StringVarP(&listPagesSpaceKey, "space-key", "s", "", "Filter by space key (e.g., MYSPACE)")
	listPagesCmd.Flags().StringVarP(&listPagesStatus, "status", "S", "", "Filter by page status (current, archived, trashed)")
	listPagesCmd.Flags().StringVarP(&listPagesTitle, "title", "t", "", "Filter by page title (partial match, case-insensitive)")
}

func runListPages(_ *cobra.Command, _ []string) error {
	ctx := context.Background()
	client := newAPIClient()

	ctx = context.WithValue(ctx, swagger.ContextBasicAuth, swagger.BasicAuth{
		UserName: viper.GetString("email"),
		Password: viper.GetString("api_key"),
	})

	// Resolve space ID from space key if provided
	spaceID := listPagesSpaceID
	if listPagesSpaceKey != "" {
		resolvedID, err := getSpaceIDByKey(ctx, client, listPagesSpaceKey)
		if err != nil {
			return err
		}
		spaceID = resolvedID
	}

	var allPages []swagger.PageBulk
	var cursor string

	for {
		var result *swagger.MultiEntityResultPage
		var err error

		if spaceID > 0 {
			// Use GetPagesInSpace when space ID is specified
			req := client.PageAPI.GetPagesInSpace(ctx, spaceID).Limit(listPagesLimit)
			if listPagesStatus != "" {
				req = req.Status([]string{listPagesStatus})
			}
			if cursor != "" {
				req = req.Cursor(cursor)
			}
			result, _, err = req.Execute()
		} else {
			// Use GetPages for all pages
			req := client.PageAPI.GetPages(ctx).Limit(listPagesLimit)
			if listPagesStatus != "" {
				req = req.Status([]string{listPagesStatus})
			}
			if cursor != "" {
				req = req.Cursor(cursor)
			}
			result, _, err = req.Execute()
		}

		if err != nil {
			return fmt.Errorf("failed to list pages: %w", err)
		}

		// Filter pages by title (partial match, case-insensitive)
		for _, page := range result.GetResults() {
			if listPagesTitle == "" || containsIgnoreCase(page.GetTitle(), listPagesTitle) {
				allPages = append(allPages, page)
			}
		}

		// Check if there are more pages
		links, hasLinks := result.GetLinksOk()
		if !hasLinks || !links.HasNext() {
			break
		}

		// Extract cursor from the next URL
		nextURL := links.GetNext()
		cursor, err = extractCursor(nextURL)
		if err != nil {
			return fmt.Errorf("failed to parse pagination cursor: %w", err)
		}
	}

	printPages(allPages)
	return nil
}

// getSpaceIDByKey looks up a space by its key and returns its numeric ID.
func getSpaceIDByKey(ctx context.Context, client *swagger.APIClient, spaceKey string) (int64, error) {
	result, _, err := client.SpaceAPI.GetSpaces(ctx).Keys([]string{spaceKey}).Limit(1).Execute()
	if err != nil {
		return 0, fmt.Errorf("failed to look up space by key: %w", err)
	}

	spaces := result.GetResults()
	if len(spaces) == 0 {
		return 0, fmt.Errorf("space with key '%s' not found", spaceKey)
	}

	// Space ID is returned as a string, convert to int64
	spaceIDStr := spaces[0].GetId()
	spaceID, err := strconv.ParseInt(spaceIDStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid space ID format: %w", err)
	}

	return spaceID, nil
}

// containsIgnoreCase checks if s contains substr (case-insensitive).
func containsIgnoreCase(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}

func printPages(pages []swagger.PageBulk) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tTITLE\tSTATUS\tSPACE ID\tPARENT ID")

	for _, page := range pages {
		status := ""
		if page.HasStatus() {
			status = string(page.GetStatus())
		}
		parentID := ""
		if page.HasParentId() {
			parentID = page.GetParentId()
		}
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n",
			page.GetId(),
			page.GetTitle(),
			status,
			page.GetSpaceId(),
			parentID,
		)
	}
	w.Flush()
}
