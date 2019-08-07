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
	m       sync.Mutex
	root    *Window
	screen  tcell.Screen
	windows []*Window
)

func RunFunc(main func(*Window) error, view func(*Box)) error {
	return nil
}

func Run(w window) error {
	return nil
}

func poll() {
	for {
		e := screen.PollEvent()
		m.Lock()
		for _, w := range windows {
			select {
			case <-w.Done():
			case w.Events() <- e:
			default:
			}
		}
		m.Unlock()
	}
}

// func app(ui *goban.UI) error {
// 	popup(ui)
// 	confirm := widgets.NewConfirm()
// 	ui.Start(confirm)
// 	for {
// 		select {
// 		case <-ui.Done():
// 		case e := <-ui.Events():
// 		default:
// 		}
// 	}
// 	return nil
// }
//
// type someWidget struct {
// 	goban.Window
// 	cursor int
// 	done   chan struct{}
// }
//
// func (w *someWidget) Main(ui *goban.UI) {
// 	<-ui.Events()
// }
//
// func (w *someWidget) View(b *goban.Box) {
// 	b.Prints("confirming")
// }
//
// func (w *someWidget) Done() chan struct{} { return w.done }
// func (w *someWidget) Close()              { close(w.done) }
//
// func popup(ui *goban.UI) {
// 	ui.PushViewFunc(func(b *goban.Box) {
// 		b.Prints(msg)
// 	})
// 	defer ui.PopView()
//
// 	ui.Show()
// 	defer ui.Show()
//
// 	for {
// 		if k := ui.Events().ReadKey(); k.Rune() == ' ' {
// 			return false
// 		}
// 		return true
// 	}
// }

// func poll(ctx context.Context) error {
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			return nil
// 		default:
// 		}
//
// 		e := screen.PollEvent()
// 		switch e := e.(type) {
// 		case *tcell.EventKey:
// 			switch e.Key() {
// 			case tcell.KeyCtrlC:
// 				return ErrAborted
// 			}
// 		}
//
// 		// broadcast
// 		for _, r := range receivers {
// 			r := r
// 			go func() {
// 				r <- e
// 			}()
// 		}
// 	}
// }

// Main runs the main process.
// When the app exits, Main also exits.
// If cancelled for any other reason, Main exits and the
// cancellation is propagated to the context.
// func Main(app func(Events) error, viewfns ...func()) error {
// 	m.Lock()
// 	defer m.Unlock()
//
// 	var err error
//
// 	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
// 	screen, err = tcell.NewScreen()
// 	if err != nil {
// 		return err
// 	}
// 	defer screen.Fini()
//
// 	if err = screen.Init(); err != nil {
// 		return err
// 	}
//
// 	screen.SetStyle(tcell.StyleDefault)
// 	screen.EnableMouse()
// 	screen.Clear()
//
// 	for _, f := range viewfns {
// 		PushViewFunc(f)
// 	}
//
// 	events := make(Events)
// 	once := &sync.Once{}
// 	done := make(chan struct{})
// 	ctx, cancel := context.WithCancel(context.Background())
// 	sigc := make(chan os.Signal)
// 	signal.Notify(sigc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
//
// 	go func() {
// 		e := app(events)
// 		once.Do(func() {
// 			err = e
// 			close(done)
// 		})
// 	}()
// 	go func() {
// 		e := poll(ctx, events)
// 		once.Do(func() {
// 			err = e
// 			close(done)
// 		})
// 	}()
// 	go func() {
// 		<-sigc
// 		once.Do(func() {
// 			err = ErrAborted
// 			close(done)
// 		})
// 	}()
//
// 	<-done
// 	cancel()
//
// 	return err
// }
//
// func poll(ctx context.Context, es Events) error {
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			return nil
// 		default:
// 		}
// 		e := screen.PollEvent()
// 		switch e := e.(type) {
// 		case *tcell.EventKey:
// 			switch e.Key() {
// 			case tcell.KeyCtrlC:
// 				return ErrAborted
// 			}
// 		}
// 		go func() {
// 			es <- e
// 		}()
// 	}
// }
