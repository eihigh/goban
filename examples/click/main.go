package main

import (
	"context"

	"github.com/eihigh/goban"
)

var (
	text = "click here"

	grid = goban.NewGrid(
		"    1fr    1fr",
		"1fr msg    msg",
		"3em cancel ok",
	)
)

func main() {
	goban.Main(app)
}

func app(_ context.Context, es goban.Events) error {
	v := &buttonView{}
	goban.PushView(v)

	for {
		goban.Show()
		switch e := es.ReadMouse(); {
		case v.cancel.IsClicked(e):
			text = "canceled"
		case v.ok.IsClicked(e):
			text = "ok"
		}
	}
}

type buttonView struct {
	cancel, ok *goban.Box
}

func (v *buttonView) View() {
	b := goban.NewBox(0, 0, 30, 7).CenterOf(goban.Screen()).Enclose("")
	msgBox := b.GridItem(grid, "msg")
	buf := goban.Buffer{
		HAlign: goban.AlignCenter,
	}
	buf.Prints(text)
	buf.Flush(msgBox)

	okBox := goban.NewBox(0, 0, 4, 1).CenterOf(b.GridItem(grid, "ok"))
	okBox.Prints("[ok]")
	v.ok = okBox

	cancelBox := goban.NewBox(0, 0, 8, 1).CenterOf(b.GridItem(grid, "cancel"))
	cancelBox.Prints("[cancel]")
	v.cancel = cancelBox
}
