package main

import (
	"context"

	"github.com/eihigh/goban"
)

var (
	grid = goban.NewGrid(
		"    1fr    1fr    1fr ",
		"1fr header header header",
		"3fr side   content ",
		"1fr footer footer footer",
	)
)

func main() {
	goban.Main(app, view)
}

func app(_ context.Context, es goban.Events) error {
	goban.Show()
	es.ReadKey()
	return nil
}

func view() {
	b := goban.Screen().Enclose("")
	header := b.GridItem(grid, "header").DrawSides("", 0, 0, 0, 1)
	header.Prints("Header")
	footer := b.GridItem(grid, "footer").DrawSides("", 0, 1, 0, 0)
	footer.Prints("Footer")
	side := b.GridItem(grid, "side").DrawSides("", 0, 0, 1, 0)
	side.Prints("Side")
	content := b.GridItem(grid, "content").DrawSides("", 0, 0, 1, 0)
	content.Prints("Main Content")
}
