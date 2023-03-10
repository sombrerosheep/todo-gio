package main

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type ItemsHeader struct {
	theme *material.Theme

	widget *widget.Clickable
}

func NewItemsHeader(theme *material.Theme) *ItemsHeader {
	header := ItemsHeader{
		theme:  theme,
		widget: &widget.Clickable{},
	}

	return &header
}

func (header ItemsHeader) Layout(gtx layout.Context) layout.Dimensions {
	if header.widget.Clicked() {
		state.SetStatus("New Item Clicked")
	}

	dims := layout.Flex{Axis: layout.Horizontal}.Layout(
		gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return material.Body1(header.theme, "ItemsHeader").Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.E.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return material.Button(header.theme, header.widget, "New Item").Layout(gtx)
			})
		}),
	)

	if debugLayout {
		return DebugDimensions(gtx, dims, header.theme)
	}

	return dims
}
