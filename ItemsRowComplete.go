package main

import (
	"time"

	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type ItemsRowComplete struct {
	theme     *material.Theme
	widget    *widget.Clickable
	completed time.Time
}

func NewItemsRowComplete(theme *material.Theme, completed time.Time) *ItemsRowComplete {
	row := ItemsRowComplete{
		theme:     theme,
		widget:    &widget.Clickable{},
		completed: completed,
	}

	return &row
}

func (row ItemsRowComplete) Layout(gtx layout.Context) layout.Dimensions {
	var dims layout.Dimensions

	if row.widget.Clicked() {
		state.SetStatus("Mark item complete clicked")
	}

	if row.completed.IsZero() {
		// Incomplete
		dims = material.Button(row.theme, row.widget, "Complete").Layout(gtx)
	} else {
		dims = material.Body1(row.theme, row.completed.Format(dt_format)).Layout(gtx)
	}

	if DebugLayout() {
		return DebugDimensions(gtx, dims, row.theme)
	}

	return dims
}
