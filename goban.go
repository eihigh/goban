package goban

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gdamore/tcell"
)

var (
	screen tcell.Screen
	views  []View
	m      sync.Mutex
)

var (
	// ErrAborted represents that the program was aborted
	// for some reason, such as Ctrl+C or OS signals.
	ErrAborted = fmt.Errorf("program aborted")
)

// Events is an alias of chan tcell.Event and
// provides some reading functions.
type Events chan tcell.Event

// ReadKey waits for a key event to the channel and returns it.
// Other events are ignored.
func (es Events) ReadKey() *tcell.EventKey {
	for {
		e := <-es
		if e, ok := e.(*tcell.EventKey); ok {
			return e
		}
	}
}

// Show calls each View to refresh the screen.
func Show() {
	screen.Clear()
	for _, v := range views {
		v.View()
	}
	screen.Show()
}

// Sync is similar to Show, but calls tcell.Screen's Sync instead.
func Sync() {
	screen.Clear()
	for _, v := range views {
		v.View()
	}
	screen.Sync()
}

// Main runs the main process.
// When the app exits, Main also exits.
// If cancelled for any other reason, Main exits and the
// cancellation is propagated to the context.
func Main(app func(context.Context, Events) error, viewfns ...func()) error {
	m.Lock()
	defer m.Unlock()

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
	screen.Clear()

	for _, f := range viewfns {
		PushViewFunc(f)
	}

	events := make(Events)
	once := &sync.Once{}
	done := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	sigc := make(chan os.Signal)
	signal.Notify(sigc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		e := app(ctx, events)
		once.Do(func() {
			err = e
			close(done)
		})
	}()
	go func() {
		e := poll(ctx, events)
		once.Do(func() {
			err = e
			close(done)
		})
	}()
	go func() {
		<-sigc
		once.Do(func() {
			err = ErrAborted
			close(done)
		})
	}()

	<-done
	cancel()

	return err
}

func poll(ctx context.Context, es Events) error {
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
		go func() {
			es <- e
		}()
	}
}
