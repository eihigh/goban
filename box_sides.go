package goban

func (b *Box) Enclose(title string) *Box {
	newb := b.DrawSides(title, 1, 1, 1, 1)
	newb.Clear()
	return newb
}

func (b *Box) DrawSides(title string, left, top, right, bottom int) *Box {
	newb := b.InsideSides(left, top, right, bottom)

	if left != 0 {
		x := newb.Pos.X - 1
		for y := 0; y < newb.Size.Y; y++ {
			screen.SetContent(x, y+newb.Pos.Y, '│', nil, b.Style)
		}
	}

	if right != 0 {
		x := newb.Pos.X + newb.Size.X
		for y := 0; y < newb.Size.Y; y++ {
			screen.SetContent(x, y+newb.Pos.Y, '│', nil, b.Style)
		}
	}

	if top != 0 {
		y := newb.Pos.Y - 1
		for x := 0; x < newb.Size.X; x++ {
			screen.SetContent(x+newb.Pos.X, y, '─', nil, b.Style)
		}
	}

	if bottom != 0 {
		y := newb.Pos.Y + newb.Size.Y
		for x := 0; x < newb.Size.X; x++ {
			screen.SetContent(x+newb.Pos.X, y, '─', nil, b.Style)
		}
	}

	ax := b.Pos.X
	ay := b.Pos.Y
	bx := ax
	by := ay
	if b.Size.X > 0 {
		bx += b.Size.X - 1
	}
	if b.Size.Y > 0 {
		by += b.Size.Y - 1
	}

	if left != 0 {
		if top == 0 {
			b.joinACS(ax, ay-1)
		} else {
			b.setACS(ax, ay, '┌')
		}
		if bottom == 0 {
			b.joinACS(ax, by+1)
		} else {
			b.setACS(ax, by, '└')
		}
	}
	if right != 0 {
		if top == 0 {
			b.joinACS(bx, ay-1)
		} else {
			b.setACS(bx, ay, '┐')
		}
		if bottom == 0 {
			b.joinACS(bx, by+1)
		} else {
			b.setACS(bx, by, '┘')
		}
	}
	if top != 0 {
		if left == 0 {
			b.joinACS(ax-1, ay)
		}
		if right == 0 {
			b.joinACS(bx+1, ay)
		}
	}
	if bottom != 0 {
		if left == 0 {
			b.joinACS(ax-1, by)
		}
		if right == 0 {
			b.joinACS(bx+1, by)
		}
	}

	// draw title
	tb := NewBox(b.Pos.X+1, b.Pos.Y, b.Size.X-1, 1)
	if left != 0 {
		tb.Pos.X++
		tb.Size.X--
	}
	tb.Prints(title)

	return newb
}

func (b *Box) InsideSides(left, top, right, bottom int) *Box {
	newb := NewBox(b.Pos.X, b.Pos.Y, b.Size.X, b.Size.Y)

	if left != 0 {
		newb.Pos.X++
		newb.Size.X--
	}

	if right != 0 {
		newb.Size.X--
	}

	if top != 0 {
		newb.Pos.Y++
		newb.Size.Y--
	}

	if bottom != 0 {
		newb.Size.Y--
	}
	return newb
}

func (b *Box) setACS(x, y int, r rune) {
	a0 := acsAt(x, y)
	a1 := rune2acs(r)
	screen.SetContent(x, y, (a0 | a1).Rune(), nil, b.Style)
}

func (b *Box) joinACS(x, y int) {
	me := acsAt(x, y)

	if acsAt(x-1, y)&acsR != 0 {
		me = me | acsL
	} else {
		me = me &^ acsL
	}

	if acsAt(x, y-1)&acsB != 0 {
		me = me | acsT
	} else {
		me = me &^ acsT
	}

	if acsAt(x+1, y)&acsL != 0 {
		me = me | acsR
	} else {
		me = me &^ acsR
	}

	if acsAt(x, y+1)&acsT != 0 {
		me = me | acsB
	} else {
		me = me &^ acsB
	}

	screen.SetContent(x, y, me.Rune(), nil, b.Style)
}

func acsAt(x, y int) acs {
	r, _, _, _ := screen.GetContent(x, y)
	return rune2acs(r)
}
