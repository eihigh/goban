package main

import (
	"github.com/eihigh/goban"
)

func main() {
	goban.Main(app, view)
}

func app(es goban.Events) error {
	goban.Show()
	es.ReadKey()
	return nil
}

func view() {
	goban.Screen().Enclose("hello").Prints("Hello World!\nPress any key to exit.")
}
