package goban

// View represents a drawing function.
type View interface {
	View(*Window)
}

type UI interface {
	View
	Main(*Layer)
}

type ViewFunc func(*Window)

func (f ViewFunc) View(w *Window) {
	f(w)
}
