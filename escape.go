package goban

import (
	"io"
	"strconv"

	"github.com/gdamore/tcell"
)

type scanner struct {
	rd io.RuneReader
	ch rune
	b  *Box
}

func (s *scanner) next() {
	s.ch, _, _ = s.rd.ReadRune()
}

func (s *scanner) scan() {
	s.next() // skip '\e'
	if s.ch != '[' {
		return
	}
	s.next() // skip '['

	params := []int{}
	for {
		switch {
		case '0' <= s.ch && s.ch <= '9':
			params = append(params, s.scanNumber())
		case s.ch == 'm':
			s.apply(params)
			s.next()
		case s.ch == ';':
			s.next()
		default:
			return
		}
	}
}

func (s *scanner) scanNumber() int {
	str := string(s.ch)

loop:
	for {
		s.next()
		switch {
		case '0' <= s.ch && s.ch <= '9':
			str += string(s.ch)
		default:
			break loop
		}
	}

	n, _ := strconv.Atoi(str)
	return n
}

func (s *scanner) apply(params []int) {
	for i := 0; i < len(params); i++ {
		param := params[i]

		switch {
		case param == 0:
			s.b.Style = tcell.StyleDefault
		case param == 1:
			s.b.Style = s.b.Style.Bold(true)
		case param == 4:
			s.b.Style = s.b.Style.Underline(true)

		case 30 <= param && param <= 37:
			s.b.Style = s.b.Style.Foreground(tcell.Color(param - 30))
		case 40 <= param && param <= 47:
			s.b.Style = s.b.Style.Background(tcell.Color(param - 40))
		case 90 <= param && param <= 97:
			s.b.Style = s.b.Style.Foreground(tcell.Color(param - 90))
		case 100 <= param && param <= 107:
			s.b.Style = s.b.Style.Background(tcell.Color(param - 100))

		case param == 38:
			n, c := extColor(params[i+1:])
			s.b.Style = s.b.Style.Foreground(c)
			i += n
		case param == 48:
			n, c := extColor(params[i+1:])
			s.b.Style = s.b.Style.Background(c)
			i += n
		}
	}
}

func extColor(p []int) (int, tcell.Color) {
	switch x, xs := shift(p); x {
	case 5: // 256 colors
		code, _ := shift(xs)
		if code == -1 {
			return 0, tcell.ColorDefault
		}
		return 2, tcell.Color(code)

	case 2:
		r, xs := shift(xs)
		g, xs := shift(xs)
		b, _ := shift(xs)
		if r != -1 && g != -1 && b != -1 {
			return 4, tcell.NewRGBColor(r, g, b)
		}
	}

	return 0, tcell.ColorDefault
}

func shift(s []int) (int32, []int) {
	if len(s) == 0 {
		return -1, s[:0]
	}
	return int32(s[0]), s[1:]
}
