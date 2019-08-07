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
	cb       *tcell.CellBuffer
	done     chan struct{}
}

type Application interface {
	Main(*Window) error
	View(*Box)
}

func newWindow() *Window {
	w := &Window{}
	pushWindow(w)
	return w
}

func (w *Window) box() *Box {
	return nil
}

func (w *Window) layer() layer {
	if w.cb == nil {
		return screen
	}
	return w.cb
}

func (w *Window) render(dst layer) {
	w.Lock()
	// render myself
	copyLayer(dst, w.layer())
	// render children
	for _, child := range w.children {
		child.render(dst)
	}
	w.Unlock()
}

func (w *Window) Show() {
	b := w.box()
	w.Lock()
	for _, v := range w.views {
		v.View(b)
	}
	w.Unlock()
	render()
}

func (w *Window) Events() Events {
	return w.events
}

func (w *Window) Run(app Application) error {
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
