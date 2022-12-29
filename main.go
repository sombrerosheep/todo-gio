package main

import (
	"log"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget/material"
)

var state *State

func main() {
	var err error

	state, err = NewStateFromFile("./todo.json")
	if err != nil {
		log.Println(err)
		return
	}

	for _, v := range state.GetListKeys() {
		log.Println(v)
	}

	go func() {
		w := app.NewWindow()
		loop(w)
	}()

	app.Main()
}

func loop(w *app.Window) {
	th := material.NewTheme(gofont.Collection())
	var ops op.Ops

	for {
		e := <-w.Events()

		switch e := e.(type) {
		case system.FrameEvent:
			{
				gtx := layout.NewContext(&ops, e)

				var list layout.List
				layout.W.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
						layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
							names := state.GetListKeys()
							return list.Layout(gtx, len(names), func(gtx layout.Context, i int) layout.Dimensions {
								return material.Body1(th, names[i]).Layout(gtx)
							})
						}),
					)
				})

				e.Frame(gtx.Ops)
			}

		}
	}
}
