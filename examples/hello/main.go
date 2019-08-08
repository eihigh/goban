package main

import (
	"github.com/eihigh/goban"
)

func main() {
	goban.RunFunc(app, view)
}

func app(w *goban.Window) error {
	// confirm := &Confirm{}
	// w.Run(confirm)
	// w.Start(confirm)
	//
	// for {
	// 	select {
	// 	// case <-confirm.Done():
	// 	case <-w.Events():
	// 	default:
	// 	}
	// }
	// return nil
	w.Show()
	w.Events().ReadKey()

	menu := &menuView{}
	w.PushView(menu)
	defer w.PopView()
	w.Show()
	w.Events().ReadKey()

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
	cursor int
}

func (c *Confirm) Main(w *goban.Window) error {
	w.Show()
	<-w.Events()
	return nil
}

func (c *Confirm) View(b *goban.Box) {
	b.Print(c.cursor)
}
