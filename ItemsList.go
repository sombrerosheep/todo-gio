package main

import (
	"log"

	"gioui.org/layout"
	"gioui.org/widget/material"
)

type ItemsList struct {
	theme    *material.Theme
	listName string

	rows []*ItemsRow
}

func NewItemsList(theme *material.Theme, listName string) *ItemsList {
	items, err := state.GetList(state.GetSelected())
	if err != nil {
		log.Println(err)
		return nil
	}

	list := ItemsList{
		theme:    theme,
		listName: listName,

		rows: make([]*ItemsRow, len(items)),
	}

	for i, v := range items {
		list.rows[i] = NewItemsRow(theme, v)
	}

	return &list
}

func (list ItemsList) Layout(gtx layout.Context) layout.Dimensions {
	itemList := layout.List{Axis: layout.Vertical}

	dims := itemList.Layout(gtx, len(list.rows), func(gtx layout.Context, i int) layout.Dimensions {
		return list.rows[i].Layout(gtx)
	})

	if DebugLayout() {
		return DebugDimensions(gtx, dims, list.theme)
	}

	return dims

}
