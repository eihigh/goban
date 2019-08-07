package main

import (
	"github.com/eihigh/goban"
)

func main() {
	goban.RunFunc(Main, view)
}

func Main(w *goban.Window) error {
	confirm := &Confirm{}
	w.Run(confirm)

	menu := &menuView{}
	w.PushView(menu)
	defer w.PopView()

	for {
		select {
		case <-confirm.Done():
		case <-w.Events():
		}
	}
	return nil
}

func view(b *goban.Box) {
	b.Prints("hoge")
}

type menuView struct {
	cursor int
}

func (v *menuView) View(b *goban.Box) {
	b.Print(v.cursor)
}

type Confirm struct {
	goban.Window
	cursor int
}

func (c *Confirm) Main() error {
	<-c.Events()
	return nil
}

func (c *Confirm) View(b *goban.Box) {
	b.Print(c.cursor)
}
