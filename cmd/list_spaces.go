package cmd

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"github.com/alexhokl/confluence-cli/swagger"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listSpacesLimit int32
var listSpacesType string
var listSpacesStatus string

var listSpacesCmd = &cobra.Command{
	Use:   "spaces",
	Short: "List all spaces",
	Long:  "List all Confluence spaces that the user has permission to view.",
	RunE:  runListSpaces,
}

func init() {
	listCmd.AddCommand(listSpacesCmd)
	listSpacesCmd.Flags().Int32VarP(&listSpacesLimit, "limit", "l", 25, "Maximum number of spaces per page")
	listSpacesCmd.Flags().StringVarP(&listSpacesType, "type", "t", "", "Filter by space type (global, personal, etc.)")
	listSpacesCmd.Flags().StringVarP(&listSpacesStatus, "status", "s", "", "Filter by space status (current, archived)")
}

func runListSpaces(_ *cobra.Command, _ []string) error {
	ctx := context.Background()
	client := newAPIClient()

	ctx = context.WithValue(ctx, swagger.ContextBasicAuth, swagger.BasicAuth{
		UserName: viper.GetString("email"),
		Password: viper.GetString("api_key"),
	})

	var allSpaces []swagger.SpaceBulk
	var cursor string

	for {
		req := client.SpaceAPI.GetSpaces(ctx).Limit(listSpacesLimit)
		if listSpacesType != "" {
			req = req.Type_(listSpacesType)
		}
		if listSpacesStatus != "" {
			req = req.Status(listSpacesStatus)
		}
		if cursor != "" {
			req = req.Cursor(cursor)
		}

		result, _, err := req.Execute()
		if err != nil {
			return fmt.Errorf("failed to list spaces: %w", err)
		}

		allSpaces = append(allSpaces, result.GetResults()...)

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

	printSpaces(allSpaces)
	return nil
}

func newAPIClient() *swagger.APIClient {
	cfg := swagger.NewConfiguration()
	org := viper.GetString("organization")
	cfg.Servers = swagger.ServerConfigurations{
		swagger.ServerConfiguration{
			URL: fmt.Sprintf("https://%s.atlassian.net/wiki/api/v2", org),
		},
	}
	return swagger.NewAPIClient(cfg)
}

func extractCursor(nextURL string) (string, error) {
	parsed, err := url.Parse(nextURL)
	if err != nil {
		return "", err
	}
	return parsed.Query().Get("cursor"), nil
}

func printSpaces(spaces []swagger.SpaceBulk) {
	color.NoColor = noColor

	yellow := color.New(color.FgYellow).SprintFunc()
	magenta := color.New(color.FgMagenta).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	w := newTableWriter(os.Stdout, 0, 2)
	w.row("KEY", "NAME", "TYPE", "STATUS")

	for _, space := range spaces {
		spaceType := ""
		if space.HasType() {
			spaceType = string(space.GetType())
		}
		status := ""
		if space.HasStatus() {
			status = string(space.GetStatus())
		}

		// Color the status based on value
		coloredStatus := status
		switch status {
		case "current":
			coloredStatus = green(status)
		case "archived":
			coloredStatus = red(status)
		}

		w.row(
			yellow(space.GetKey()),
			space.GetName(),
			magenta(spaceType),
			coloredStatus,
		)
	}
	w.flush()
}
