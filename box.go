package goban

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/gdamore/tcell"
	"github.com/mattn/go-runewidth"
)

var (
	errOutOfBox = fmt.Errorf("position outside of the box")
)

// Point represents a point on the screen.
type Point struct {
	X, Y int
}

// Box represents an area on the screen and provides
// various drawing and layout functions.
type Box struct {
	Pos    Point
	Size   Point
	Scroll Point
	Style  tcell.Style

	cursor Point
}

type runeReader interface {
	io.RuneReader
	UnreadRune() error
}

// NewBox returns a box at the specified position and size.
func NewBox(x, y, w, h int) *Box {
	return &Box{
		Pos:  Point{x, y},
		Size: Point{w, h},
	}
}

// Screen returns a box of screen size.
func Screen() *Box {
	w, h := screen.Size()
	return &Box{
		Pos:  Point{0, 0},
		Size: Point{w, h},
	}
}

func (b *Box) Fit(dst *Box, dstx, dsty, srcx, srcy float64) *Box {
	spx, spy := b.rel(srcx, srcy)
	dpx, dpy := dst.rel(dstx, dsty)
	dx, dy := dpx-spx, dpy-spy
	return NewBox(b.Pos.X+dx, b.Pos.Y+dy, b.Size.X, b.Size.Y)
}

func (b *Box) CenterOf(dst *Box) *Box {
	return b.Fit(dst, 0, 0, 0, 0)
}

func (b *Box) rel(x, y float64) (int, int) {
	x = x/2 + 0.5 // [-1, 1] => [0, 1]
	y = y/2 + 0.5
	px := b.Pos.X + int(float64(b.Size.X)*x)
	py := b.Pos.Y + int(float64(b.Size.Y)*y)
	return px, py
}

func (b *Box) IsClicked(e *tcell.EventMouse) bool {
	if b == nil {
		return false
	}
	if e.Buttons()&tcell.Button1 == 0 {
		return false
	}
	x, y := e.Position()
	ax, ay := b.Pos.X, b.Pos.Y
	bx, by := ax+b.Size.X, ay+b.Size.Y
	if ax <= x && x < bx && ay <= y && y < by {
		return true
	}
	return false
}

func (b *Box) Clear() {
	for x := b.Pos.X; x < b.Pos.X+b.Size.X; x++ {
		for y := b.Pos.Y; y < b.Pos.Y+b.Size.Y; y++ {
			screen.SetContent(x, y, ' ', nil, b.Style)
		}
	}
}

func (b *Box) Write(p []byte) (int, error) {
	buf := bytes.NewBuffer(p)
	err := b.print(buf)
	return len(p), err
}

func (b *Box) Print(a ...interface{}) {
	b.Prints(fmt.Sprint(a...))
}

func (b *Box) Printf(format string, a ...interface{}) {
	b.Prints(fmt.Sprintf(format, a...))
}

func (b *Box) Println(a ...interface{}) {
	b.Prints(fmt.Sprintln(a...))
}

func (b *Box) Puts(s string) {
	b.Prints(s + "\n")
}

func (b *Box) Prints(s string) {
	b.print(strings.NewReader(s))
}

func (b *Box) print(reader runeReader) error {
	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			break
		}

		switch r {
		case '\r':
			b.newLine()
			b.cursor.Y--
		case '\n':
			b.newLine()
		case 0x1b: // escape sequence
			b.escape(reader)

		default:
			x, y, err := b.actualPoint()
			if err == nil {
				screen.SetContent(x, y, r, nil, b.Style)
				w := runewidth.RuneWidth(r)
				b.cursor.X += w
			}
		}
	}

	return nil
}

func (b *Box) actualPoint() (x, y int, err error) {
	// virtual point
	vx, vy := b.cursor.X-b.Scroll.X, b.cursor.Y-b.Scroll.Y
	if vx < 0 || vy < 0 || vx >= b.Size.X || vy >= b.Size.Y {
		return 0, 0, errOutOfBox
	}

	// actual point
	x, y = b.Pos.X+vx, b.Pos.Y+vy
	return
}

func (b *Box) newLine() {
	for {
		x, y, err := b.actualPoint()
		if err != nil {
			break
		}
		screen.SetContent(x, y, ' ', nil, b.Style)
		b.cursor.X++
	}
	b.cursor.X = 0
	b.cursor.Y++
}

func (b *Box) escape(rd runeReader) {
	s := scanner{
		rd: rd,
		b:  b,
	}
	s.scan()
	rd.UnreadRune()
}
