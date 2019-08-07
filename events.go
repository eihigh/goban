package goban

import "github.com/gdamore/tcell"

// Events is an alias of chan tcell.Event and
// provides some reading functions.
type Events chan tcell.Event

// makeEvents makes Events with a fixed capacity.
func makeEvents() Events {
	return make(Events, 16)
}

// ReadKey waits for a key event to the channel and returns it.
// Other events are ignored.
func (es Events) ReadKey() *tcell.EventKey {
	for {
		e := <-es
		if e, ok := e.(*tcell.EventKey); ok {
			return e
		}
	}
}

// ReadMouse waits for a mouse event to the channel and returns it.
// Other events are ignored.
func (es Events) ReadMouse() *tcell.EventMouse {
	for {
		e := <-es
		if e, ok := e.(*tcell.EventMouse); ok {
			return e
		}
	}
}
