# goban - minimal and concurrent CUI

package goban (碁盤, meainig of Go game board in Japanese) provides CUI with simple API.

## Hello World

```
package main

import (
	"context"

	"github.com/eihigh/goban"
)

func main() {
	goban.Main(app, view)
}

func app(_ context.Context, es goban.Events) error {
	goban.Show()
	es.ReadKey()
	return nil
}

func view() {
	goban.Screen().Enclose("hello").Prints("Hello World!\nPress any key to exit.")
}
```

## Features

* Minimal API
* Isolated Views and Controllers
* Receive events from channel instead of event handlers
* Color with escape sequences
* Box drawings
* Grid layouts

## Status

goban is under active development. The API is subject to change.

## TODO

* [] Flexbox layouts
* [] More widgets
* [] Mouse support

## Documentation

See https://godoc.org/github.com/eihigh/goban .

## Dependencies

This package is based on github.com/gdamore/tcell .
