package goban

import (
	"strconv"
	"strings"
)

type Length struct {
	Unit  LengthUnit
	Value int
}

type LengthUnit int

const (
	Em LengthUnit = iota
	Fr
)

type item struct {
	left, top, right, bottom int
}

type Grid struct {
	Cols, Rows []Length
	items      map[string]*item
}

func NewGrid(layout ...string) *Grid {
	g := &Grid{
		items: map[string]*item{},
	}

	if len(layout) == 0 {
		return g
	}

	head := layout[0]

	// col tracks
	for _, col := range strings.Fields(head) {
		l := parseLength(col)
		g.Cols = append(g.Cols, l)
	}

	for y, row := range layout[1:] {
		cells := strings.Fields(row)
		head := cells[0]
		l := parseLength(head)
		g.Rows = append(g.Rows, l)

		for x, name := range cells[1:] {
			if a, ok := g.items[name]; ok {
				if a.right <= x {
					a.right = x + 1
				}
				if a.bottom <= y {
					a.bottom = y + 1
				}
			} else {
				g.items[name] = &item{
					left:   x,
					top:    y,
					right:  x + 1,
					bottom: y + 1,
				}
			}
		}
	}

	return g
}

func parseLength(s string) Length {
	for u := Em; u <= Fr; u++ {
		suffix := u.String()
		if strings.HasSuffix(s, suffix) {
			x, err := strconv.Atoi(strings.TrimSuffix(s, suffix))
			if err != nil {
				panic(err)
			}
			return Length{u, x}
		}
	}
	panic("undefined length name")
}

func (u LengthUnit) String() string {
	switch u {
	case Em:
		return "em"
	case Fr:
		return "fr"
	default:
		return ""
	}
}

func (b *Box) GridItem(g *Grid, name string) *Box {
	a := g.items[name]
	return b.GridCell(g, a.left, a.top, a.right, a.bottom)
}

func (b *Box) GridCell(g *Grid, left, top, right, bottom int) *Box {
	cols := absLengths(b.Size.X, g.Cols)
	rows := absLengths(b.Size.Y, g.Rows)

	left = cols[left]
	top = rows[top]
	right = cols[right]
	bottom = rows[bottom]

	return NewBox(b.Pos.X+left, b.Pos.Y+top, right-left, bottom-top)
}

func NFr(n int) Length {
	return Length{Fr, n}
}

func NEm(n int) Length {
	return Length{Em, n}
}

func absLengths(total int, lens []Length) []int {
	x := total
	frs := 0
	abs := []int{0}
	for _, l := range lens {
		if l.Unit == Fr {
			frs += l.Value
		} else {
			x -= l.Value // Em
		}
	}

	last := 0
	for _, l := range lens {
		if l.Unit == Fr {
			last += x / frs * l.Value
		} else {
			last += l.Value
		}
		abs = append(abs, last)
	}

	if lens[len(lens)-1].Unit == Fr {
		abs[len(abs)-1] = total
	}
	return abs
}
