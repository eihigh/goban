package goban

import (
	"context"
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
	m      sync.Mutex
	root   *Window
	screen tcell.Screen

	receivers []Events // TODO Lock
)

type Application interface {
	Main(context.Context, *Window) error
	Draw(*Box)
}

func addReceiver(ch Events) {
	receivers = append(receivers, ch)
}

func removeReceiver(ch Events) {
	i := -1
	for j, r := range receivers {
		if r == ch {
			i = j
		}
	}
	if i != -1 {
		receivers = append(receivers[:i], receivers[i+1:]...)
	}
}

func render() {
	screen.Clear()
	root.render(screen)
	screen.Show()
}

type appFunc struct {
	main func(context.Context, *Window) error
	draw func(*Box)
}

func (a appFunc) Main(ctx context.Context, w *Window) error {
	return a.main(ctx, w)
}

func (a appFunc) Draw(b *Box) {
	a.draw(b)
}

func RunFunc(main func(ctx context.Context, w *Window) error, draw func(*Box)) error {
	a := appFunc{main, draw}
	return Run(a)
}

func Run(a Application) error {

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

	m.Lock()
	root = newWindow()
	root.PushView(a)
	m.Unlock()
	ctx := context.TODO()

	addReceiver(root.events)
	err = a.Main(ctx, root)
	removeReceiver(root.events)
	return err
}

func poll(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
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

		// broadcast
		for _, r := range receivers {
			r := r
			go func() {
				r <- e
			}()
		}
	}
}

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
