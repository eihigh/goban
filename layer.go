package goban

import "github.com/gdamore/tcell"

type layer interface {
	GetContent(x, y int) (rune, []rune, tcell.Style, int)
}

func copyLayer(dst, src layer) {}
