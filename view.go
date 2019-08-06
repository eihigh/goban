package goban

// View represents a drawing function.
type View interface {
	Draw(*Box)
}

type ViewFunc func(*Box)

func (f ViewFunc) Draw(b *Box) {
	f(b)
}
