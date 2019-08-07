package goban

// View represents a drawing function.
type View interface {
	View(*Box)
}

type ViewFunc func(*Box)

func (f ViewFunc) View(b *Box) { f(b) }
