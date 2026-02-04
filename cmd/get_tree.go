package cmd

import (
	"context"
	"fmt"
	"sort"

	"github.com/alexhokl/confluence-cli/swagger"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var getTreeSpaceID int64
var getTreeSpaceKey string
var getTreeTitle string
var getTreeParentID string

var getTreeCmd = &cobra.Command{
	Use:   "tree",
	Short: "Show page tree of a space",
	Long: `Display a hierarchical tree view of all pages in a Confluence space.

The tree shows the parent-child relationships between pages, with indentation
indicating the nesting level.

If --parent-id is specified, only the subtree starting from that page will be shown.

If --title is specified, only branches containing pages with matching titles
(partial match, case-insensitive) will be shown. This includes the matching
pages, their ancestors (path to root), and their descendants.

Examples:
  # Show page tree by space ID
  confluence-cli get tree --space-id 12345

  # Show page tree by space key
  confluence-cli get tree --space-key MYSPACE

  # Show subtree starting from a specific page
  confluence-cli get tree --space-key MYSPACE --parent-id 67890

  # Show only branches with pages matching "guide" in the title
  confluence-cli get tree --space-key MYSPACE --title guide`,
	RunE: runGetTree,
}

func init() {
	getCmd.AddCommand(getTreeCmd)
	getTreeCmd.Flags().Int64Var(&getTreeSpaceID, "space-id", 0, "Space ID")
	getTreeCmd.Flags().StringVarP(&getTreeSpaceKey, "space-key", "s", "", "Space key (e.g., MYSPACE)")
	getTreeCmd.Flags().StringVarP(&getTreeTitle, "title", "t", "", "Filter by page title (partial match, case-insensitive)")
	getTreeCmd.Flags().StringVarP(&getTreeParentID, "parent-id", "p", "", "Show subtree starting from this page ID")
}

func runGetTree(_ *cobra.Command, _ []string) error {
	if getTreeSpaceID == 0 && getTreeSpaceKey == "" {
		return fmt.Errorf("either --space-id or --space-key must be specified")
	}

	if getTreeSpaceID != 0 && getTreeSpaceKey != "" {
		return fmt.Errorf("only one of --space-id or --space-key should be specified, not both")
	}

	ctx := context.Background()
	client := newAPIClient()

	ctx = context.WithValue(ctx, swagger.ContextBasicAuth, swagger.BasicAuth{
		UserName: viper.GetString("email"),
		Password: viper.GetString("api_key"),
	})

	// Resolve space ID from space key if needed
	spaceID := getTreeSpaceID
	if getTreeSpaceKey != "" {
		resolvedID, err := getSpaceIDByKey(ctx, client, getTreeSpaceKey)
		if err != nil {
			return err
		}
		spaceID = resolvedID
	}

	// Fetch all pages in the space
	pages, err := fetchAllPagesInSpace(ctx, client, spaceID)
	if err != nil {
		return err
	}

	// Build and print the tree
	printPageTree(pages, getTreeTitle, getTreeParentID)
	return nil
}

// fetchAllPagesInSpace retrieves all pages in a space with pagination.
func fetchAllPagesInSpace(ctx context.Context, client *swagger.APIClient, spaceID int64) ([]swagger.PageBulk, error) {
	var allPages []swagger.PageBulk
	var cursor string
	limit := int32(100) // Use larger limit for efficiency

	for {
		req := client.PageAPI.GetPagesInSpace(ctx, spaceID).
			Limit(limit).
			Status([]string{"current"}) // Only get current (published) pages

		if cursor != "" {
			req = req.Cursor(cursor)
		}

		result, _, err := req.Execute()
		if err != nil {
			return nil, fmt.Errorf("failed to fetch pages: %w", err)
		}

		allPages = append(allPages, result.GetResults()...)

		// Check if there are more pages
		links, hasLinks := result.GetLinksOk()
		if !hasLinks || !links.HasNext() {
			break
		}

		// Extract cursor from the next URL
		nextURL := links.GetNext()
		cursor, err = extractCursor(nextURL)
		if err != nil {
			return nil, fmt.Errorf("failed to parse pagination cursor: %w", err)
		}
	}

	return allPages, nil
}

// pageNode represents a node in the page tree.
type pageNode struct {
	id       string
	title    string
	children []*pageNode
	matched  bool // indicates if this node or any descendant matches the filter
}

// printPageTree builds a tree structure from pages and prints it.
// If parentID is non-empty, only the subtree starting from that page is shown.
// If titleFilter is non-empty, only branches containing matching pages are shown.
func printPageTree(pages []swagger.PageBulk, titleFilter string, parentID string) {
	color.NoColor = noColor

	if len(pages) == 0 {
		fmt.Println("No pages found in this space")
		return
	}

	// Create a map of page ID to page for quick lookup
	pageMap := make(map[string]*swagger.PageBulk)
	for i := range pages {
		pageMap[pages[i].GetId()] = &pages[i]
	}

	// Create nodes for each page
	nodeMap := make(map[string]*pageNode)
	for _, page := range pages {
		nodeMap[page.GetId()] = &pageNode{
			id:       page.GetId(),
			title:    page.GetTitle(),
			children: make([]*pageNode, 0),
			matched:  false,
		}
	}

	// Build tree structure by linking children to parents
	var rootNodes []*pageNode
	for _, page := range pages {
		node := nodeMap[page.GetId()]
		pageParentID := ""
		if page.HasParentId() {
			pageParentID = page.GetParentId()
		}

		if pageParentID == "" {
			// No parent - this is a root node
			rootNodes = append(rootNodes, node)
		} else if parentNode, exists := nodeMap[pageParentID]; exists {
			// Parent exists in our page set
			parentNode.children = append(parentNode.children, node)
		} else {
			// Parent doesn't exist in our set (might be in different space or deleted)
			// Treat as root node
			rootNodes = append(rootNodes, node)
		}
	}

	// Sort root nodes and all children alphabetically by title
	sortNodes(rootNodes)

	// If parent ID is specified, find that node and use it as the root
	if parentID != "" {
		parentNode, exists := nodeMap[parentID]
		if !exists {
			fmt.Printf("Page with ID '%s' not found in this space\n", parentID)
			return
		}
		// Use the parent node as the single root
		rootNodes = []*pageNode{parentNode}
	}

	// If title filter is specified, mark matching nodes and filter the tree
	if titleFilter != "" {
		// Mark nodes that match the filter (and their descendants)
		for _, root := range rootNodes {
			markMatchingNodes(root, titleFilter)
		}

		// Filter to only include branches with matches
		rootNodes = filterMatchedBranches(rootNodes)

		if len(rootNodes) == 0 {
			fmt.Printf("No pages found matching '%s'\n", titleFilter)
			return
		}
	}

	// Create color functions for tree output
	yellow := color.New(color.FgYellow).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()

	// Print the tree
	for i, root := range rootNodes {
		isLast := i == len(rootNodes)-1
		printNode(root, "", isLast, yellow, cyan)
	}
}

// markMatchingNodes marks nodes that match the title filter or have matching descendants.
// When a node matches, all its descendants are also marked.
// Returns true if this node or any descendant matches.
func markMatchingNodes(node *pageNode, titleFilter string) bool {
	// Check if this node's title matches
	nodeMatches := containsIgnoreCase(node.title, titleFilter)

	// If this node matches, mark all descendants
	if nodeMatches {
		markAllDescendants(node)
		return true
	}

	// Check if any children match (recursively)
	childMatches := false
	for _, child := range node.children {
		if markMatchingNodes(child, titleFilter) {
			childMatches = true
		}
	}

	// Mark this node if it has matching descendants
	node.matched = childMatches
	return node.matched
}

// markAllDescendants marks a node and all its descendants as matched.
func markAllDescendants(node *pageNode) {
	node.matched = true
	for _, child := range node.children {
		markAllDescendants(child)
	}
}

// filterMatchedBranches returns only nodes that are marked as matched.
// It also recursively filters children.
func filterMatchedBranches(nodes []*pageNode) []*pageNode {
	var result []*pageNode
	for _, node := range nodes {
		if node.matched {
			// Create a copy with filtered children
			filteredNode := &pageNode{
				id:       node.id,
				title:    node.title,
				children: filterMatchedBranches(node.children),
				matched:  node.matched,
			}
			result = append(result, filteredNode)
		}
	}
	return result
}

// sortNodes sorts a slice of nodes alphabetically by title, and recursively sorts children.
func sortNodes(nodes []*pageNode) {
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].title < nodes[j].title
	})
	for _, node := range nodes {
		sortNodes(node.children)
	}
}

// printNode prints a single node and its children with proper tree formatting.
func printNode(node *pageNode, prefix string, isLast bool, yellow func(a ...interface{}) string, cyan func(a ...interface{}) string) {
	// Choose the appropriate connector
	connector := "├── "
	if isLast {
		connector = "└── "
	}

	// Print this node with colored ID
	fmt.Printf("%s%s%s %s\n", cyan(prefix), cyan(connector), node.title, yellow("(ID: "+node.id+")"))

	// Determine the prefix for children
	childPrefix := prefix
	if isLast {
		childPrefix += "    "
	} else {
		childPrefix += "│   "
	}

	// Print children
	for i, child := range node.children {
		childIsLast := i == len(node.children)-1
		printNode(child, childPrefix, childIsLast, yellow, cyan)
	}
}
