package main

import (
	"context"

	"github.com/eihigh/goban"
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
	goban.Screen().Enclose("hello").Prints("Hello World!\nPress any key to exit.")
}
