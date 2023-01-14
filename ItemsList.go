package main

import (
	"log"

	"gioui.org/layout"
	"gioui.org/widget/material"
)

type ItemsList struct {
	theme    *material.Theme
	listName string
	rows     []ItemsRow
}

func NewItemsList(theme *material.Theme, listName string) *ItemsList {
	list := ItemsList{
		theme:    theme,
		listName: listName,
	}

	list.SetSelectedItems(listName)

	return &list
}

func (list *ItemsList) SetSelectedItems(listName string) {
	items, err := state.GetList(listName)
	if err != nil {
		log.Println(err)
		return
	}

	list.rows = make([]ItemsRow, len(items))

	for i, v := range items {
		list.rows[i] = NewItemsRow(list.theme, v)
	}

	return
}

func (list *ItemsList) Layout(gtx layout.Context) layout.Dimensions {
	selected := state.GetSelected()
	if list.listName != selected {
		list.listName = selected
		list.SetSelectedItems(selected)
	}

	itemList := layout.List{Axis: layout.Vertical}

	dims := itemList.Layout(gtx, len(list.rows), func(gtx layout.Context, i int) layout.Dimensions {
		return list.rows[i].Layout(gtx)
	})

	if DebugLayout() {
		return DebugDimensions(gtx, dims, list.theme)
	}

	return dims

}
