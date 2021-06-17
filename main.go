package main

import (
	_ "embed"
	"github.com/qaware/dev-tool-kit/backend/core"
	"github.com/qaware/dev-tool-kit/backend/ui"
	"github.com/wailsapp/wails"
	"io/ioutil"
	"os"
	"path"
	"runtime/debug"
	"time"
)

var version = "3.4.0"

//go:embed frontend/build/main.js
var js string

//go:embed frontend/build/main.css
var css string

func main() {
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

	defer panicRecover()

	app.Run()
}

func panicRecover() {
	recovered := recover()
	if recovered == nil {
		return
	}

	exe, err := os.Executable()
	if err != nil {
		return
	}
	dir := path.Dir(exe)

	appError, ok := recovered.(error)
	if ok {
		content := time.Now().String() + "\nVersion: " + version + "\n\n" + appError.Error() + "\n" + string(debug.Stack())
		_ = ioutil.WriteFile(path.Join(dir, "panic.txt"), []byte(content), 0777)
	}
}
