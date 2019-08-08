package goban

import (
	"fmt"
	"sync"

	"github.com/gdamore/tcell"
)

var (
	// ErrAborted represents that the program was aborted
	// for some reason, such as Ctrl+C or OS signals.
	ErrAborted = fmt.Errorf("program aborted")
)

var (
	screen tcell.Screen
	root   *Window

	windows struct {
		sync.Mutex
		slice []*Window
	}
)

func RunFunc(main func(*Window) error, view func(*Box)) error {
	return nil
}

func Run(app Application) error {
	root.Lock()
	defer root.Unlock()
	var err error

	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	screen, err = tcell.NewScreen()
	if err != nil {
		return err
	}
	defer screen.Fini()

	if err = screen.Init(); err != nil {
		return err
	}

	screen.SetStyle(tcell.StyleDefault)
	screen.EnableMouse()
	screen.Clear()

	root = newWindow()
	root.PushView(app)

	once := &sync.Once{}
	done := make(chan struct{})

	go func() {
		e := app.Main(root)
		once.Do(func() {
			err = e
			close(done)
		})
	}()
	go func() {
		e := poll(done)
		once.Do(func() {
			err = e
			close(done)
		})
	}()

	<-done
	close(root.done)

	return err
}

func poll(done chan struct{}) error {
	for {
		select {
		case <-done:
			return nil
		default:
		}

		e := screen.PollEvent()
		switch e := e.(type) {
		case *tcell.EventKey:
			switch e.Key() {
			case tcell.KeyCtrlC:
				return ErrAborted
			}
		}
		dispatch(e)
	}
}

func dispatch(e tcell.Event) {
	windows.Lock()
	del := []int{}
	for i, w := range windows.slice {
		select {
		case <-w.done:
			del = append(del, i)
			continue
		default:
		}

		select {
		case w.events <- e:
		default:
		}
	}
	next := make([]*Window, 0, len(windows.slice))

	for i, w := range windows.slice {
		keep := true
		for _, j := range del {
			if i == j {
				keep = false
				break
			}
		}
		if keep {
			next = append(next, w)
		}
	}
	windows.slice = next
	windows.Unlock()
}

func pushWindow(w *Window) {
	windows.Lock()
	windows.slice = append(windows.slice, w)
	windows.Unlock()
}

func render() {
	screen.Clear()
	root.render()
	screen.Show()
}
