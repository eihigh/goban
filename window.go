package goban

import "github.com/gdamore/tcell"

type Window struct {
	events Events
	views  []View
	layer  tcell.CellBuffer
}

type Layer struct{}

func newWindow() *Window {
	return &Window{}
}

func (w *Window) Box() *Box {
	return nil
}

func (w *Window) Events() Events {
	return w.events
}

func (w *Window) Show() {
	for _, v := range w.views {
		v.View(w)
	}
}

func (w *Window) PushView(v View) {
	w.views = append(w.views, v)
}

func (w *Window) PushViewFunc(f func()) {
	v := func(*Window) {
		f()
	}
	w.PushView(ViewFunc(v))
}

func (w *Window) PopView() {
	w.views = w.views[:len(w.views)-1]
}

func (w *Window) Start(ui UI) {
	child := newWindow()
	child.PushView(ui)
	go ui.Main(child)
}
