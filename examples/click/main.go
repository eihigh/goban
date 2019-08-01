package main

import (
	"context"

	"github.com/eihigh/goban"
)

var (
	text = "click here"
)

func main() {
	goban.Main(app)
}

func app(_ context.Context, es goban.Events) error {
	v := &buttonView{}
	goban.PushView(v)

	for {
		goban.Show()
		if e := es.ReadMouse(); v.button.IsClicked(e) {
			if text == "on" {
				text = "off"
			} else {
				text = "on"
			}
		}
	}
}

type buttonView struct {
	button *goban.Box
}

func (v *buttonView) View() {
	v.button = goban.NewBox(0, 0, 15, 3).CenterOf(goban.Screen()).Enclose("button")

	b := goban.Buffer{
		HAlign: goban.AlignCenter,
	}
	b.Prints(text)
	b.Flush(v.button)
}
