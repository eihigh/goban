package goban

import (
	"sync"

	"github.com/gdamore/tcell"
)

type Window struct {
	sync.Mutex
	events   Events
	recv     Events
	views    []View
	children []*Window
	done     chan struct{}

	cb tcell.CellBuffer
}

func newWindow() *Window {
	return &Window{}
}

func (w *Window) window() *Window {
	return w
}

func (w *Window) box() *Box {
	return nil
}

func (w *Window) Show() {
	b := w.box()
	w.Lock()
	for _, v := range w.views {
		v.Draw(b)
	}
	w.Unlock()
}

func (w *Window) Events() Events {
	return w.events
}

func (w *Window) Send(e tcell.Event) {
	w.events <- e
}

func (w *Window) Done() chan struct{} {
	return w.done
}

func (w *Window) Close() {
	close(w.done)
}

func (w *Window) Run(a Application) {
	w.PushView(a)
	child := a.window()
	child.Close()
}

func (w *Window) Link(a Application) Events {
	child := a.window()
	ch := make(Events, 1)
	go func() {
		for e := range w.events {
			ch <- e
			select {
			case <-child.Done():
				return
			default:
				child.Send(e)
			}
		}
	}()
	return ch
}

func (w *Window) dispatch() {
	for e := range w.events {
		w.recv <- e
		for _, c := range w.children {
			c.Send(e)
		}
	}
}

// func (w *Window) Run(ctx context.Context, a Application) error {
// 	w.Lock()
// 	child := newWindow()
// 	child.PushView(a)
// 	w.children = append(w.children, child)
// 	w.Unlock()
// 	return a.Main(ctx, child)
// }
//
func (w *Window) PushView(v View) {
	w.Lock()
	w.views = append(w.views, v)
	w.Unlock()
}

// func (w *Window) PushViewFunc(v ViewFunc) {
// 	w.Lock()
// 	w.views = append(w.views, v)
// 	w.Unlock()
// }
//
// func (w *Window) PopView() {
// 	w.Lock()
// 	w.views = w.views[:len(w.views)-1]
// 	w.Unlock()
// }
//
// func (w *Window) resize() {
//
// }
//
// type layer interface {
// 	GetContent(x, y int) (rune, []rune, tcell.Style, int)
// 	SetContent(int, int, rune, []rune, tcell.Style)
// }
//
// func copyLayer(src, dst layer) {
//
// }
//
// func (w *Window) render(dst layer) {
// 	w.Lock()
// 	copyLayer(&w.cb, dst)
// 	w.Unlock()
//
// 	for _, child := range w.children {
// 		child.render(dst)
// 	}
// }
