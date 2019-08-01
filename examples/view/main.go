package main

import (
	"context"

	"github.com/eihigh/goban"
	"github.com/gdamore/tcell"
)

func main() {
	goban.Main(app)
}

func app(_ context.Context, es goban.Events) error {
	v := &menuView{
		items: []string{
			"foo", "bar", "baz",
		},
	}
	goban.PushView(v)

	for {
		goban.Show()
		switch k := es.ReadKey(); k.Key() {
		case tcell.KeyUp:
			if v.cursor > 0 {
				v.cursor--
			}
		case tcell.KeyDown:
			if v.cursor < len(v.items)-1 {
				v.cursor++
			}
		}
	}
}

// view model implements goban.View.
type menuView struct {
	cursor int
	items  []string
}

func (v *menuView) View() {
	b := goban.NewBox(0, 0, 50, 20).Enclose("menu")
	b.Puts("↑, ↓: move cursor")
	for i, item := range v.items {
		if i == v.cursor {
			b.Print("> ")
		}
		b.Puts(item)
	}
}
