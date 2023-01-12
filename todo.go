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

	listPanel  *ListPanel
	itemsPanel *ItemsPanel
	statusBar  *StatusBar
}

func NewTodoApp(w *app.Window) *TodoApp {
	theme := material.NewTheme(gofont.Collection())

	todo := TodoApp{
		window:     w,
		theme:      theme,
		listPanel:  NewListPanel(theme),
		itemsPanel: NewItemsPanel(theme),
		statusBar:  &StatusBar{theme: theme},
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
	layout.Flex{
		Axis:    layout.Axis(layout.Vertical),
		Spacing: layout.SpaceStart,
	}.Layout(
		gtx,
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{
				Axis: layout.Axis(layout.Horizontal),
			}.Layout(
				gtx,
				layout.Flexed(0.35, todo.listPanel.Layout),
				layout.Flexed(1, todo.itemsPanel.Layout),
			)
		}),
		layout.Rigid(todo.statusBar.Layout),
	)
}
