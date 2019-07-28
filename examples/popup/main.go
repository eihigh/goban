package main

import (
	"context"

	"github.com/eihigh/goban"
)

func main() {
	goban.Main(app)
}

func app(_ context.Context, es goban.Events) error {
	v := func() {
		b := goban.Screen()
		b.Puts("Press any key to open popup")
		b.Puts("Ctrl+C to exit")
	}
	goban.PushViewFunc(v)

	for {
		goban.Show()
		es.ReadKey()
		popup(es)
	}
}

func popup(es goban.Events) {
	v := func() {
		b := goban.NewBox(0, 0, 40, 5).Enclose("popup window")
		b.Prints("Press any key to close popup")
	}

	// This is the recommended way to use `PushView` and `defer PopView`
	// when using modal views such as popup.
	goban.PushViewFunc(v)
	defer goban.PopView()

	goban.Show()
	es.ReadKey()
}
