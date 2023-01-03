package main

import (
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget/material"
)

type TodoApp struct {
	window *app.Window

	theme *material.Theme
}

func NewTodoApp(w *app.Window) *TodoApp {
	theme := material.NewTheme(gofont.Collection())

	todo := TodoApp{
		window: w,
		theme:  theme,
	}

	return &todo
}

func (todo *TodoApp) Run() error {
	var ops op.Ops

	for {
		select {
		case e := <-todo.window.Events():

			switch e := e.(type) {
			case system.FrameEvent:
				gtx := layout.NewContext(&ops, e)
				todo.Layout(gtx)
				e.Frame(gtx.Ops)

			case system.DestroyEvent:
				return e.Err
			}
		}
	}
}

func (todo *TodoApp) Layout(gtx layout.Context) {
	layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.N.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return material.H5(todo.theme, "todo:>").Layout(gtx)
			})
		}),
		ListSelections(state, todo.theme, gtx),
	)
}
