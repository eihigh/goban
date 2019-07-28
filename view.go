package goban

// View represents a drawing function.
type View interface {
	View()
}

// PushView pushes the view on top.
func PushView(v View) {
	views = append(views, v)
}

// PopView pops the view on top.
func PopView() {
	views = views[:len(views)-1]
}

// RemoveView removes the specified view.
func RemoveView(v View) {
	is := []int{}
	for i, view := range views {
		if v == view {
			is = append(is, i)
			break
		}
	}
	if len(is) == 0 {
		return
	}
	for _, i := range is {
		views = append(views[:i], views[i+1:]...)
	}
}

type viewFunc func()

func (f viewFunc) View() {
	f()
}

// PushViewFunc pushes the function as view on top.
func PushViewFunc(f func()) {
	PushView(viewFunc(f))
}
