package main

import (
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget/material"
)

func main() {
	go func(){
		w := app.NewWindow()
	if err := loop(w); err != nil {
		log.Printf("err: %v", err)
	}
	os.Exit(0)
	}()

	app.Main()
}

func loop(w *app.Window) error {
	var ops op.Ops
	var th = material.NewTheme(gofont.Collection())

	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)
			h1 := material.H1(th, "hello gio")
			h1.Color = color.NRGBA{G: 0xf0, A:0xff}
			layout.Center.Layout(gtx, h1.Layout)
			e.Frame(gtx.Ops)
		}
	}
}
