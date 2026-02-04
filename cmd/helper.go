package cmd

import (
	"fmt"
	"io"
	"regexp"

	"github.com/mattn/go-runewidth"
)

// ansiRegex matches ANSI escape sequences
var ansiRegex = regexp.MustCompile(`\x1b\[[0-9;]*m`)

// stripAnsi removes ANSI escape sequences from a string
func stripAnsi(s string) string {
	return ansiRegex.ReplaceAllString(s, "")
}

// tableWriter provides ANSI-aware tabular output.
// It calculates column widths ignoring ANSI escape codes.
type tableWriter struct {
	out      io.Writer
	rows     [][]string
	minWidth int
	padding  int
	padChar  byte
}

// newTableWriter creates a new tableWriter.
// minWidth is the minimum column width, padding is the space between columns.
func newTableWriter(out io.Writer, minWidth, padding int) *tableWriter {
	return &tableWriter{
		out:      out,
		rows:     make([][]string, 0),
		minWidth: minWidth,
		padding:  padding,
		padChar:  ' ',
	}
}

// row adds a row of cells to the table.
func (t *tableWriter) row(cells ...string) {
	t.rows = append(t.rows, cells)
}

// flush calculates column widths and writes the formatted table.
func (t *tableWriter) flush() {
	if len(t.rows) == 0 {
		return
	}

	// Calculate max visible width for each column
	maxCols := 0
	for _, row := range t.rows {
		if len(row) > maxCols {
			maxCols = len(row)
		}
	}

	colWidths := make([]int, maxCols)
	for _, row := range t.rows {
		for i, cell := range row {
			// Calculate visible width (without ANSI codes, accounting for wide characters like emojis)
			visibleWidth := runewidth.StringWidth(stripAnsi(cell))
			if visibleWidth > colWidths[i] {
				colWidths[i] = visibleWidth
			}
		}
	}

	// Apply minimum width
	for i := range colWidths {
		if colWidths[i] < t.minWidth {
			colWidths[i] = t.minWidth
		}
	}

	// Write rows with proper padding
	for _, row := range t.rows {
		for i, cell := range row {
			fmt.Fprint(t.out, cell)

			// Add padding (except for last column)
			if i < len(row)-1 {
				visibleWidth := runewidth.StringWidth(stripAnsi(cell))
				padding := colWidths[i] - visibleWidth + t.padding
				for j := 0; j < padding; j++ {
					fmt.Fprint(t.out, string(t.padChar))
				}
			}
		}
		fmt.Fprintln(t.out)
	}
}
