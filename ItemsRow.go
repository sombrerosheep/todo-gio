package main

import (
	"time"

	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

const dt_format = time.RFC822

type ItemsRow struct {
	theme  *material.Theme
	widget *widget.Clickable
	item   Item

	complete *ItemsRowComplete
}

func NewItemsRow(theme *material.Theme, item Item) ItemsRow {
	row := ItemsRow{
		theme:    theme,
		widget:   &widget.Clickable{},
		item:     item,
		complete: NewItemsRowComplete(theme, item.Completed),
	}

	return row
}

func (row ItemsRow) Layout(gtx layout.Context) layout.Dimensions {
	dims := layout.Flex{
		Axis: layout.Axis(layout.Horizontal),
	}.Layout(
		gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return material.Body1(row.theme, row.item.Name).Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return material.Body1(row.theme, row.item.Created.Format(dt_format)).Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return row.complete.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return material.Button(row.theme, row.widget, "Remove").Layout(gtx)
		}),
	)

	if DebugLayout() {
		return DebugDimensions(gtx, dims, row.theme)
	}

	return dims
}
