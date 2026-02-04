package cmd

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/alexhokl/confluence-cli/swagger"
)

func TestGetTreeFlags(t *testing.T) {
	cmd := getTreeCmd

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

	// Check title flag exists
	titleFlag := cmd.Flags().Lookup("title")
	if titleFlag == nil {
		t.Error("expected 'title' flag to exist")
	}
	if titleFlag.Shorthand != "t" {
		t.Errorf("expected 'title' flag shorthand to be 't', got '%s'", titleFlag.Shorthand)
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

func TestRunGetTreeValidation(t *testing.T) {
	// Save original values
	originalSpaceID := getTreeSpaceID
	originalSpaceKey := getTreeSpaceKey

	defer func() {
		getTreeSpaceID = originalSpaceID
		getTreeSpaceKey = originalSpaceKey
	}()

	tests := []struct {
		name        string
		spaceID     int64
		spaceKey    string
		expectError bool
		errorMsg    string
	}{
		{
			name:        "missing space identifier",
			spaceID:     0,
			spaceKey:    "",
			expectError: true,
			errorMsg:    "either --space-id or --space-key must be specified",
		},
		{
			name:        "both space identifiers provided",
			spaceID:     12345,
			spaceKey:    "MYSPACE",
			expectError: true,
			errorMsg:    "only one of --space-id or --space-key should be specified, not both",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getTreeSpaceID = tt.spaceID
			getTreeSpaceKey = tt.spaceKey

			err := runGetTree(nil, nil)

			if tt.expectError {
				if err == nil {
					t.Errorf("expected error containing '%s', got nil", tt.errorMsg)
					return
				}
				if err.Error() != tt.errorMsg {
					t.Errorf("expected error '%s', got '%s'", tt.errorMsg, err.Error())
				}
			}
		})
	}
}

func TestGetTreeCommandMetadata(t *testing.T) {
	cmd := getTreeCmd

	if cmd.Use != "tree" {
		t.Errorf("expected Use to be 'tree', got '%s'", cmd.Use)
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

func TestSortNodes(t *testing.T) {
	// Create unsorted nodes
	nodes := []*pageNode{
		{id: "3", title: "Zebra", children: []*pageNode{}},
		{id: "1", title: "Apple", children: []*pageNode{}},
		{id: "2", title: "Banana", children: []*pageNode{
			{id: "2c", title: "Banana Child C", children: []*pageNode{}},
			{id: "2a", title: "Banana Child A", children: []*pageNode{}},
			{id: "2b", title: "Banana Child B", children: []*pageNode{}},
		}},
	}

	sortNodes(nodes)

	// Check root level sorting
	if nodes[0].title != "Apple" {
		t.Errorf("expected first node to be 'Apple', got '%s'", nodes[0].title)
	}
	if nodes[1].title != "Banana" {
		t.Errorf("expected second node to be 'Banana', got '%s'", nodes[1].title)
	}
	if nodes[2].title != "Zebra" {
		t.Errorf("expected third node to be 'Zebra', got '%s'", nodes[2].title)
	}

	// Check children sorting
	bananaChildren := nodes[1].children
	if bananaChildren[0].title != "Banana Child A" {
		t.Errorf("expected first child to be 'Banana Child A', got '%s'", bananaChildren[0].title)
	}
	if bananaChildren[1].title != "Banana Child B" {
		t.Errorf("expected second child to be 'Banana Child B', got '%s'", bananaChildren[1].title)
	}
	if bananaChildren[2].title != "Banana Child C" {
		t.Errorf("expected third child to be 'Banana Child C', got '%s'", bananaChildren[2].title)
	}
}

func TestPrintPageTree(t *testing.T) {
	// Helper to create a PageBulk with optional parent
	createPage := func(id, title, parentID string) swagger.PageBulk {
		page := swagger.NewPageBulk()
		page.SetId(id)
		page.SetTitle(title)
		if parentID != "" {
			page.SetParentId(parentID)
		}
		return *page
	}

	tests := []struct {
		name           string
		pages          []swagger.PageBulk
		expectedOutput []string
	}{
		{
			name:           "empty pages",
			pages:          []swagger.PageBulk{},
			expectedOutput: []string{"No pages found"},
		},
		{
			name: "single root page",
			pages: []swagger.PageBulk{
				createPage("1", "Root Page", ""),
			},
			expectedOutput: []string{"└── Root Page (ID: 1)"},
		},
		{
			name: "multiple root pages",
			pages: []swagger.PageBulk{
				createPage("1", "Page A", ""),
				createPage("2", "Page B", ""),
			},
			expectedOutput: []string{"├── Page A (ID: 1)", "└── Page B (ID: 2)"},
		},
		{
			name: "parent with children",
			pages: []swagger.PageBulk{
				createPage("1", "Parent", ""),
				createPage("2", "Child 1", "1"),
				createPage("3", "Child 2", "1"),
			},
			expectedOutput: []string{
				"└── Parent (ID: 1)",
				"    ├── Child 1 (ID: 2)",
				"    └── Child 2 (ID: 3)",
			},
		},
		{
			name: "nested hierarchy",
			pages: []swagger.PageBulk{
				createPage("1", "Root", ""),
				createPage("2", "Child", "1"),
				createPage("3", "Grandchild", "2"),
			},
			expectedOutput: []string{
				"└── Root (ID: 1)",
				"    └── Child (ID: 2)",
				"        └── Grandchild (ID: 3)",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture stdout
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			printPageTree(tt.pages, "", "")

			w.Close()
			os.Stdout = oldStdout

			var buf bytes.Buffer
			io.Copy(&buf, r)
			output := buf.String()

			for _, expected := range tt.expectedOutput {
				if !strings.Contains(output, expected) {
					t.Errorf("expected output to contain '%s', got:\n%s", expected, output)
				}
			}
		})
	}
}

func TestPrintNode(t *testing.T) {
	// Create a simple tree
	root := &pageNode{
		id:    "1",
		title: "Root",
		children: []*pageNode{
			{
				id:       "2",
				title:    "Child",
				children: []*pageNode{},
			},
		},
	}

	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	printNode(root, "", true)

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	// Check output contains expected elements
	if !strings.Contains(output, "└── Root (ID: 1)") {
		t.Errorf("expected output to contain root node, got:\n%s", output)
	}
	if !strings.Contains(output, "└── Child (ID: 2)") {
		t.Errorf("expected output to contain child node, got:\n%s", output)
	}
}

func TestPrintPageTreeWithTitleFilter(t *testing.T) {
	// Helper to create a PageBulk with optional parent
	createPage := func(id, title, parentID string) swagger.PageBulk {
		page := swagger.NewPageBulk()
		page.SetId(id)
		page.SetTitle(title)
		if parentID != "" {
			page.SetParentId(parentID)
		}
		return *page
	}

	tests := []struct {
		name             string
		pages            []swagger.PageBulk
		titleFilter      string
		expectedOutput   []string
		unexpectedOutput []string
	}{
		{
			name: "filter matches leaf node - shows path to root",
			pages: []swagger.PageBulk{
				createPage("1", "Root", ""),
				createPage("2", "Parent", "1"),
				createPage("3", "Target Guide", "2"),
			},
			titleFilter:      "guide",
			expectedOutput:   []string{"Root", "Parent", "Target Guide"},
			unexpectedOutput: []string{},
		},
		{
			name: "filter matches middle node - shows ancestors and descendants",
			pages: []swagger.PageBulk{
				createPage("1", "Root", ""),
				createPage("2", "User Guide", "1"),
				createPage("3", "Chapter 1", "2"),
				createPage("4", "Chapter 2", "2"),
			},
			titleFilter:      "guide",
			expectedOutput:   []string{"Root", "User Guide", "Chapter 1", "Chapter 2"},
			unexpectedOutput: []string{},
		},
		{
			name: "filter matches root node - shows all descendants",
			pages: []swagger.PageBulk{
				createPage("1", "Guide Root", ""),
				createPage("2", "Child 1", "1"),
				createPage("3", "Child 2", "1"),
			},
			titleFilter:      "guide",
			expectedOutput:   []string{"Guide Root", "Child 1", "Child 2"},
			unexpectedOutput: []string{},
		},
		{
			name: "filter excludes non-matching branches",
			pages: []swagger.PageBulk{
				createPage("1", "Documentation", ""),
				createPage("2", "User Guide", "1"),
				createPage("3", "API Reference", "1"),
				createPage("4", "Other Section", ""),
				createPage("5", "Unrelated Page", "4"),
			},
			titleFilter:      "guide",
			expectedOutput:   []string{"Documentation", "User Guide"},
			unexpectedOutput: []string{"Other Section", "Unrelated Page", "API Reference"},
		},
		{
			name: "filter with no matches",
			pages: []swagger.PageBulk{
				createPage("1", "Root", ""),
				createPage("2", "Child", "1"),
			},
			titleFilter:      "nonexistent",
			expectedOutput:   []string{"No pages found matching"},
			unexpectedOutput: []string{"Root", "Child"},
		},
		{
			name: "filter is case insensitive",
			pages: []swagger.PageBulk{
				createPage("1", "Root", ""),
				createPage("2", "USER GUIDE", "1"),
			},
			titleFilter:      "guide",
			expectedOutput:   []string{"Root", "USER GUIDE"},
			unexpectedOutput: []string{},
		},
		{
			name: "empty filter shows all pages",
			pages: []swagger.PageBulk{
				createPage("1", "Root", ""),
				createPage("2", "Child", "1"),
			},
			titleFilter:      "",
			expectedOutput:   []string{"Root", "Child"},
			unexpectedOutput: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture stdout
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			printPageTree(tt.pages, tt.titleFilter, "")

			w.Close()
			os.Stdout = oldStdout

			var buf bytes.Buffer
			io.Copy(&buf, r)
			output := buf.String()

			for _, expected := range tt.expectedOutput {
				if !strings.Contains(output, expected) {
					t.Errorf("expected output to contain '%s', got:\n%s", expected, output)
				}
			}

			for _, unexpected := range tt.unexpectedOutput {
				if strings.Contains(output, unexpected) {
					t.Errorf("expected output NOT to contain '%s', got:\n%s", unexpected, output)
				}
			}
		})
	}
}

func TestMarkMatchingNodes(t *testing.T) {
	// Create a tree structure
	grandchild := &pageNode{id: "3", title: "Target Guide", children: []*pageNode{}}
	child := &pageNode{id: "2", title: "Parent", children: []*pageNode{grandchild}}
	root := &pageNode{id: "1", title: "Root", children: []*pageNode{child}}

	// Mark matching nodes
	result := markMatchingNodes(root, "guide")

	// Root should be marked (has matching descendant)
	if !root.matched {
		t.Error("expected root to be marked as matched")
	}

	// Child should be marked (has matching descendant)
	if !child.matched {
		t.Error("expected child to be marked as matched")
	}

	// Grandchild should be marked (matches directly)
	if !grandchild.matched {
		t.Error("expected grandchild to be marked as matched")
	}

	if !result {
		t.Error("expected markMatchingNodes to return true")
	}
}

func TestFilterMatchedBranches(t *testing.T) {
	// Create nodes with mixed matched status
	nodes := []*pageNode{
		{id: "1", title: "Matched", matched: true, children: []*pageNode{
			{id: "1a", title: "Matched Child", matched: true, children: []*pageNode{}},
			{id: "1b", title: "Unmatched Child", matched: false, children: []*pageNode{}},
		}},
		{id: "2", title: "Unmatched", matched: false, children: []*pageNode{}},
	}

	result := filterMatchedBranches(nodes)

	// Should only have one root node
	if len(result) != 1 {
		t.Errorf("expected 1 filtered node, got %d", len(result))
	}

	// Should only have one child
	if len(result[0].children) != 1 {
		t.Errorf("expected 1 filtered child, got %d", len(result[0].children))
	}

	// The child should be the matched one
	if result[0].children[0].title != "Matched Child" {
		t.Errorf("expected matched child, got '%s'", result[0].children[0].title)
	}
}

func TestMarkMatchingNodesNoMatch(t *testing.T) {
	// Create a tree structure with no matches
	grandchild := &pageNode{id: "3", title: "Grandchild Page", children: []*pageNode{}}
	child := &pageNode{id: "2", title: "Parent Page", children: []*pageNode{grandchild}}
	root := &pageNode{id: "1", title: "Root Page", children: []*pageNode{child}}

	// Mark matching nodes with a filter that won't match anything
	result := markMatchingNodes(root, "nonexistent")

	// Nothing should be matched
	if root.matched {
		t.Error("expected root to NOT be marked as matched")
	}
	if child.matched {
		t.Error("expected child to NOT be marked as matched")
	}
	if grandchild.matched {
		t.Error("expected grandchild to NOT be marked as matched")
	}
	if result {
		t.Error("expected markMatchingNodes to return false")
	}
}

func TestMarkMatchingNodesMatchesRoot(t *testing.T) {
	// Create a tree structure where root matches
	grandchild := &pageNode{id: "3", title: "Grandchild", children: []*pageNode{}}
	child := &pageNode{id: "2", title: "Child", children: []*pageNode{grandchild}}
	root := &pageNode{id: "1", title: "Target Guide", children: []*pageNode{child}}

	// Mark matching nodes
	result := markMatchingNodes(root, "guide")

	// All nodes should be marked (root matches, so all descendants should be included)
	if !root.matched {
		t.Error("expected root to be marked as matched")
	}
	if !child.matched {
		t.Error("expected child to be marked as matched (descendant of match)")
	}
	if !grandchild.matched {
		t.Error("expected grandchild to be marked as matched (descendant of match)")
	}
	if !result {
		t.Error("expected markMatchingNodes to return true")
	}
}

func TestMarkAllDescendants(t *testing.T) {
	// Create a tree structure
	grandchild1 := &pageNode{id: "3", title: "Grandchild 1", children: []*pageNode{}, matched: false}
	grandchild2 := &pageNode{id: "4", title: "Grandchild 2", children: []*pageNode{}, matched: false}
	child := &pageNode{id: "2", title: "Child", children: []*pageNode{grandchild1, grandchild2}, matched: false}
	root := &pageNode{id: "1", title: "Root", children: []*pageNode{child}, matched: false}

	// Mark all descendants
	markAllDescendants(root)

	// All nodes should be marked
	if !root.matched {
		t.Error("expected root to be marked")
	}
	if !child.matched {
		t.Error("expected child to be marked")
	}
	if !grandchild1.matched {
		t.Error("expected grandchild1 to be marked")
	}
	if !grandchild2.matched {
		t.Error("expected grandchild2 to be marked")
	}
}

func TestSortNodesEmpty(t *testing.T) {
	// Test sorting empty slice doesn't panic
	nodes := []*pageNode{}
	sortNodes(nodes)
	if len(nodes) != 0 {
		t.Error("expected empty slice after sorting empty slice")
	}
}

func TestSortNodesSingleNode(t *testing.T) {
	// Test sorting single node
	node := &pageNode{id: "1", title: "Only Node", children: []*pageNode{}}
	nodes := []*pageNode{node}
	sortNodes(nodes)
	if len(nodes) != 1 {
		t.Error("expected single node after sorting")
	}
	if nodes[0].title != "Only Node" {
		t.Errorf("expected title 'Only Node', got '%s'", nodes[0].title)
	}
}

func TestFilterMatchedBranchesEmpty(t *testing.T) {
	// Test filtering empty slice
	nodes := []*pageNode{}
	result := filterMatchedBranches(nodes)
	if len(result) != 0 {
		t.Error("expected empty slice when filtering empty input")
	}
}

func TestFilterMatchedBranchesAllUnmatched(t *testing.T) {
	// Test filtering when all nodes are unmatched
	nodes := []*pageNode{
		{id: "1", title: "Unmatched 1", matched: false, children: []*pageNode{}},
		{id: "2", title: "Unmatched 2", matched: false, children: []*pageNode{}},
	}
	result := filterMatchedBranches(nodes)
	if len(result) != 0 {
		t.Errorf("expected empty slice when all unmatched, got %d", len(result))
	}
}

func TestFilterMatchedBranchesAllMatched(t *testing.T) {
	// Test filtering when all nodes are matched
	nodes := []*pageNode{
		{id: "1", title: "Matched 1", matched: true, children: []*pageNode{}},
		{id: "2", title: "Matched 2", matched: true, children: []*pageNode{}},
	}
	result := filterMatchedBranches(nodes)
	if len(result) != 2 {
		t.Errorf("expected 2 nodes when all matched, got %d", len(result))
	}
}

func TestPrintPageTreeOrphanedPages(t *testing.T) {
	// Test with pages that have parent IDs pointing to non-existent parents
	createPage := func(id, title, parentID string) swagger.PageBulk {
		page := swagger.NewPageBulk()
		page.SetId(id)
		page.SetTitle(title)
		if parentID != "" {
			page.SetParentId(parentID)
		}
		return *page
	}

	// Page with parent ID that doesn't exist in the set
	pages := []swagger.PageBulk{
		createPage("1", "Orphaned Page", "999"), // Parent 999 doesn't exist
		createPage("2", "Another Orphaned", "888"),
	}

	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	printPageTree(pages, "", "")

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	// Both pages should be treated as root nodes
	if !strings.Contains(output, "Orphaned Page") {
		t.Errorf("expected output to contain 'Orphaned Page', got:\n%s", output)
	}
	if !strings.Contains(output, "Another Orphaned") {
		t.Errorf("expected output to contain 'Another Orphaned', got:\n%s", output)
	}
}

func TestPrintPageTreeWithParentID(t *testing.T) {
	// Helper to create a PageBulk with optional parent
	createPage := func(id, title, parentID string) swagger.PageBulk {
		page := swagger.NewPageBulk()
		page.SetId(id)
		page.SetTitle(title)
		if parentID != "" {
			page.SetParentId(parentID)
		}
		return *page
	}

	tests := []struct {
		name             string
		pages            []swagger.PageBulk
		parentID         string
		expectedOutput   []string
		unexpectedOutput []string
	}{
		{
			name: "show subtree from middle node",
			pages: []swagger.PageBulk{
				createPage("1", "Root", ""),
				createPage("2", "Parent Node", "1"),
				createPage("3", "Child 1", "2"),
				createPage("4", "Child 2", "2"),
				createPage("5", "Other Branch", "1"),
			},
			parentID:         "2",
			expectedOutput:   []string{"Parent Node", "Child 1", "Child 2"},
			unexpectedOutput: []string{"Root", "Other Branch"},
		},
		{
			name: "show subtree from leaf node",
			pages: []swagger.PageBulk{
				createPage("1", "Root", ""),
				createPage("2", "Parent", "1"),
				createPage("3", "Leaf", "2"),
			},
			parentID:         "3",
			expectedOutput:   []string{"Leaf"},
			unexpectedOutput: []string{"Root", "Parent"},
		},
		{
			name: "show full tree from root",
			pages: []swagger.PageBulk{
				createPage("1", "Root", ""),
				createPage("2", "Child", "1"),
			},
			parentID:         "1",
			expectedOutput:   []string{"Root", "Child"},
			unexpectedOutput: []string{},
		},
		{
			name: "parent-id not found",
			pages: []swagger.PageBulk{
				createPage("1", "Root", ""),
				createPage("2", "Child", "1"),
			},
			parentID:         "999",
			expectedOutput:   []string{"Page with ID '999' not found"},
			unexpectedOutput: []string{"Root", "Child"},
		},
		{
			name: "empty parent-id shows all",
			pages: []swagger.PageBulk{
				createPage("1", "Root", ""),
				createPage("2", "Child", "1"),
			},
			parentID:         "",
			expectedOutput:   []string{"Root", "Child"},
			unexpectedOutput: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture stdout
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			printPageTree(tt.pages, "", tt.parentID)

			w.Close()
			os.Stdout = oldStdout

			var buf bytes.Buffer
			io.Copy(&buf, r)
			output := buf.String()

			for _, expected := range tt.expectedOutput {
				if !strings.Contains(output, expected) {
					t.Errorf("expected output to contain '%s', got:\n%s", expected, output)
				}
			}

			for _, unexpected := range tt.unexpectedOutput {
				if strings.Contains(output, unexpected) {
					t.Errorf("expected output NOT to contain '%s', got:\n%s", unexpected, output)
				}
			}
		})
	}
}

func TestPrintPageTreeWithParentIDAndTitleFilter(t *testing.T) {
	// Helper to create a PageBulk with optional parent
	createPage := func(id, title, parentID string) swagger.PageBulk {
		page := swagger.NewPageBulk()
		page.SetId(id)
		page.SetTitle(title)
		if parentID != "" {
			page.SetParentId(parentID)
		}
		return *page
	}

	// Test combining parent-id and title filter
	pages := []swagger.PageBulk{
		createPage("1", "Root", ""),
		createPage("2", "Documentation", "1"),
		createPage("3", "User Guide", "2"),
		createPage("4", "API Reference", "2"),
		createPage("5", "Other Section", "1"),
	}

	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Show subtree from "Documentation" and filter by "guide"
	printPageTree(pages, "guide", "2")

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	// Should show Documentation (parent) and User Guide (matches)
	if !strings.Contains(output, "Documentation") {
		t.Errorf("expected output to contain 'Documentation', got:\n%s", output)
	}
	if !strings.Contains(output, "User Guide") {
		t.Errorf("expected output to contain 'User Guide', got:\n%s", output)
	}

	// Should NOT show Root, Other Section, or API Reference
	if strings.Contains(output, "Root") {
		t.Errorf("expected output NOT to contain 'Root', got:\n%s", output)
	}
	if strings.Contains(output, "Other Section") {
		t.Errorf("expected output NOT to contain 'Other Section', got:\n%s", output)
	}
	if strings.Contains(output, "API Reference") {
		t.Errorf("expected output NOT to contain 'API Reference', got:\n%s", output)
	}
}
