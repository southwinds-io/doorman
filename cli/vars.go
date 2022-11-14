package main

import "github.com/rivo/tview"

var (
	app         *tview.Application // The tview application.
	pages       *tview.Pages       // The application pages.
	finderFocus tview.Primitive    // The primitive in the Finder that last had focus.
)

func getString(key string) (string, error) {
	return "", nil
}
