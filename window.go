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

func newWindow() *Window {
	w := &Window{
		events: makeEvents(),
		cb:     &tcell.CellBuffer{},
		done:   make(chan struct{}),
	}
	w.cb.Resize(screen.Size())
	pushWindow(w)
	return w
}

func (w *Window) box() *Box {
	width, height := w.cb.Size()
	b := NewBox(0, 0, width, height)
	b.layer = w.cb
	return b
}

func (w *Window) render() {
	w.Lock()
	// render myself
	copyLayer(screen, w.cb)
	// render children
	for _, child := range w.children {
		select {
		case <-child.done:
			continue
		default:
		}
		child.render()
	}
	w.Unlock()
}

func (w *Window) Show() {
	w.Lock()
	w.cb.Fill(' ', tcell.StyleDefault)
	for _, v := range w.views {
		v.View(w.box())
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
