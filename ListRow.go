package main

import (
	"fmt"

	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type ListRow struct {
	theme    *material.Theme
	listName string
	widget   *widget.Clickable

	rowSelected chan<- string
}

func NewListRow(listName string, rowSelected chan<- string, theme *material.Theme) *ListRow {
	row := ListRow{
		theme:       theme,
		listName:    listName,
		widget:      &widget.Clickable{},
		rowSelected: rowSelected,
	}

	return &row
}

func (ls ListRow) Layout(gtx layout.Context) layout.Dimensions {
	if ls.widget.Clicked() {
		state.SetStatus(fmt.Sprintf("Selecting list \"%s\"", ls.listName))
		ls.rowSelected <- ls.listName
	}

	dims := ls.widget.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return material.Body1(ls.theme, ls.listName).Layout(gtx)
	})

	if DebugLayout() {
		return DebugDimensions(gtx, dims, ls.theme)
	}

	return dims
}
