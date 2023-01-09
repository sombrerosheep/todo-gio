package main

import (
	"log"

	"gioui.org/layout"
	"gioui.org/widget/material"
)

type ItemsList struct {
	theme    *material.Theme
	listName string
}

func NewItemsList(theme *material.Theme, listName string) *ItemsList {
	list := ItemsList{
		theme:    theme,
		listName: listName,
	}

	return &list
}

func (list ItemsList) Layout(gtx layout.Context) layout.Dimensions {
	items, err := state.GetList(state.GetSelected())
	if err != nil {
		log.Println(err)
		return layout.Dimensions{}
	}

	itemList := layout.List{Axis: layout.Vertical}

	dims := itemList.Layout(gtx, len(items), func(gtx layout.Context, i int) layout.Dimensions {
		return ItemsRow{}.Layout(gtx, items[i], list.theme)
	})

	if DebugLayout() {
		return DebugDimensions(gtx, dims, list.theme)
	}

	return dims

}
