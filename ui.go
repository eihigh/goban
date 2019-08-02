package goban

import (
	"sync"
	"time"

	"github.com/gdamore/tcell"
)

var (
	boxes  []*Box
	uiLock sync.Mutex
)

type EventHover struct {
	when  time.Time
	box   *Box
	mouse *tcell.EventMouse
}

func NewEventHover(b *Box) *EventHover {
	return &EventHover{
		when: time.Now(),
		box:  b,
	}
}

func (e *EventHover) When() time.Time          { return e.when }
func (e *EventHover) Box() *Box                { return e.box }
func (e *EventHover) Button() tcell.ButtonMask { return e.mouse.Buttons() }

func (e *EventHover) ButtonIs(btn tcell.ButtonMask) bool {
	return e.Button()&btn != 0
}

func addBox(b *Box) {
	uiLock.Lock()
	boxes = append(boxes, b)
	uiLock.Unlock()
}

func findBox(e *tcell.EventMouse) *Box {
	uiLock.Lock()
	defer uiLock.Unlock()

	x, y := e.Position()
	for _, b := range boxes {
		ax, ay := b.Pos.X, b.Pos.Y
		bx, by := ax+b.Size.X, ay+b.Size.Y
		if ax <= x && ay <= y && x < bx && y < by {
			return b
		}
	}
	return nil
}
