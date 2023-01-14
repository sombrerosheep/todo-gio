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

type ListItem struct {
	list    string
	newItem string
}

type TodoApp struct {
	window *app.Window

	theme *material.Theme

	listPanel  *ListPanel
	itemsPanel *ItemsPanel
	statusBar  *StatusBar

	addList       chan string
	remList       chan string
	selectNewList chan string
	addItem       chan ListItem
	remItem       chan ListItem
	completeItem  chan ListItem
}

func NewTodoApp(w *app.Window) *TodoApp {
	theme := material.NewTheme(gofont.Collection())

	selectNewListChan := make(chan string, 8)

	todo := TodoApp{
		window:     w,
		theme:      theme,
		listPanel:  NewListPanel(selectNewListChan, theme),
		itemsPanel: NewItemsPanel(theme),
		statusBar:  &StatusBar{theme: theme},

		selectNewList: selectNewListChan,
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

		case newList := <-todo.addList:
			log.Printf("chan received new list: %s\n", newList)
		case list := <-todo.remList:
			log.Printf("chan received remove list: %s\n", list)
		case l := <-todo.selectNewList:
			if l == state.GetSelected() {
				break
			}

			err := state.SetSelected(l)
			if err != nil {
				log.Println(err)
			}

		case item := <-todo.addItem:
			log.Printf("chan received new item [%s](%s)\n", item.list, item.newItem)
		case item := <-todo.remItem:
			log.Printf("chan received remove item [%s](%s)\n", item.list, item.newItem)
		case item := <-todo.completeItem:
			log.Printf("chan received complete item [%s](%s)\n", item.list, item.newItem)
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
