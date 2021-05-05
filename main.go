package main

import (
	"github.com/leaanthony/mewn"
	"github.com/qaware/dev-tool-kit/backend/core"
	"github.com/qaware/dev-tool-kit/backend/ui"
	"github.com/wailsapp/wails"
)

var version = "3.1.2"

func main() {
	js := mewn.String("./frontend/build/main.js")
	css := mewn.String("./frontend/build/main.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:     1300,
		Height:    800,
		Title:     "DevToolKit",
		JS:        js,
		CSS:       css,
		Colour:    "#D4DCE5",
		Resizable: true,
	})

	bus := &ui.Bus{}
	app.Bind(bus)

	go core.InitUpgrade(version)
	go core.CreateAsciiFont()

	app.Run()
}
