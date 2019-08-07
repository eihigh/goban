package goban

import (
	"sync"

	"github.com/gdamore/tcell"
)

type Window struct {
	sync.Mutex
	events   Events
	views    []View
	children []*Window
	done     chan struct{}

	cb tcell.CellBuffer
}

type window interface {
	View
	Main() error
	window() *Window
}

func (w *Window) window() *Window     { return w }
func (w *Window) Done() chan struct{} { return w.done }
func (w *Window) Close()              { close(w.done) }

func (w *Window) box() *Box {
	return nil
}

func (w *Window) Show() {
	b := w.box()
	w.Lock()
	for _, v := range w.views {
		v.View(b)
	}
	w.Unlock()
}

func (w *Window) Events() Events {
	return w.events
}

func (w *Window) Run(child window) error {
	return nil
}

func (w *Window) RunFunc(main func(*Window) error, view *Box) error {
	return nil
}

func (w *Window) PushView(v View) {
	w.Lock()
	w.views = append(w.views, v)
	w.Unlock()
}

func (w *Window) PopView() {
	w.Lock()
	w.views = w.views[:len(w.views)-1]
	w.Unlock()
}
