package goban

type acs int

const (
	_ acs = 1 << iota
	acsL
	acsT
	acsR
	acsB
)

const (
	acsH   = '─'
	acsV   = '│'
	acsLT  = '┘'
	acsLB  = '┐'
	acsTR  = '└'
	acsRB  = '┌'
	acsLTB = '┤'
	acsTLR = '┴'
	acsRTB = '├'
	acsBLR = '┬'
	acsAll = '┼'
)

var (
	acsList = []struct {
		flag acs
		r    rune
	}{
		{acsL | acsR, acsH},
		{acsT | acsB, acsV},
		{acsL | acsT, acsLT},
		{acsL | acsB, acsLB},
		{acsT | acsR, acsTR},
		{acsR | acsB, acsRB},
		{acsL | acsT | acsB, acsLTB},
		{acsL | acsT | acsR, acsTLR},
		{acsR | acsT | acsB, acsRTB},
		{acsB | acsL | acsR, acsBLR},
		{acsL | acsT | acsR | acsB, acsAll},
	}
)

func (a acs) Rune() rune {
	for _, pair := range acsList {
		if a&(acsL|acsT|acsR|acsB) == pair.flag {
			return pair.r
		}
	}
	return 0
}

func rune2acs(r rune) acs {
	for _, pair := range acsList {
		if pair.r == r {
			return pair.flag
		}
	}
	return 0
}
