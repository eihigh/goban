package goban

import (
	"bufio"
	"strings"

	"github.com/mattn/go-runewidth"
)

type Align int

const (
	AlignStart Align = iota
	AlignCenter
	AlignEnd
)

type Buffer struct {
	VAlign, HAlign Align
	lines          []string
}

func (b *Buffer) Flush(box *Box) {
	for y, line := range b.lines {
		w := runewidth.StringWidth(line)
		x := 0
		switch b.HAlign {
		case AlignCenter:
			x += box.Size.X/2 - w/2
		case AlignEnd:
			x += box.Size.X - w
		}
		box.cursor.X = x
		box.cursor.Y = y
		box.Prints(line)
	}
}

func (b *Buffer) Prints(s string) {
	r := bufio.NewScanner(strings.NewReader(s))
	for r.Scan() {
		b.lines = append(b.lines, r.Text())
	}
}
