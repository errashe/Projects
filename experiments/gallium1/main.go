package main

import (
	"log"
	"os"
	"runtime"

	"github.com/alexflint/gallium"
)

func main() {
	runtime.LockOSThread()         // must be the first statement in main - see below
	gallium.Loop(os.Args, onReady) // must be called from main function
}

func onReady(app *gallium.App) {
	app.OpenWindow("http://example.com/", gallium.FramedWindow)
	app.SetMenu([]gallium.Menu{
		gallium.Menu{
			Title: "demo",
			Entries: []gallium.MenuEntry{
				gallium.MenuItem{
					Title:   "About",
					OnClick: handleMenuAbout,
				},
				gallium.Separator,
				gallium.MenuItem{
					Title:    "Quit",
					Shortcut: gallium.KeyCombination{"q", gallium.ModifierCmd},
					OnClick:  handleMenuQuit,
				},
			},
		},
	})
}

func handleMenuAbout() {
	log.Println("about clicked")
	os.Exit(0)
}

func handleMenuQuit() {
	log.Println("quit clicked")
	os.Exit(0)
}
