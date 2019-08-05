package main

import (
	"github.com/eihigh/goban"
	"github.com/gdamore/tcell"
)

func main() {
	goban.Run(app, view)
}

func app(w *goban.Window) error {
	m := &menu{}
	w.Start(m)
	<-w.Events()
	popup(w)

	confirm := NewConfirm()
	confirm.Start()
	<-confirm.Done()

	return nil
}

func view(w *goban.Window) {
	b := w.Box()
	b.Prints("hello")
}

func popup(w *goban.Window) {
	v := func() {
		b := w.Box()
		b.Prints("popup")
	}
	w.PushViewFunc(v)
	defer w.PopView()

	<-w.Events()
}

// implements goban.UI
type menu struct {
	cursor int
}

func (m *menu) Main(w *goban.Window) {
	for {
		w.Show()
		switch k := w.Events().ReadKey(); k.Key() {
		case tcell.KeyUp:
			m.cursor--
		case tcell.KeyDown:
			m.cursor++
		}
	}
}

func (m *menu) View(l *goban.Layer) {
	// b := w.Box()
	// b.Print(m.cursor)
}
