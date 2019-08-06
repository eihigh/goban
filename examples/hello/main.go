package main

import (
	"context"

	"github.com/eihigh/goban"
)

func main() {
	goban.RunFunc(app, draw)
}

func app(ctx context.Context, w *goban.Window) error {
	popup(w)
	w.Show()
	return nil
}

func draw(b *goban.Box) {
	b.Enclose("hello")
}

func popup(w *goban.Window) {
	v := func(b *goban.Box) {
		b.Prints("press any key to close")
	}
	w.PushViewFunc(v)
	defer w.PopView()

	w.Show()
	w.Events().ReadKey()
}
