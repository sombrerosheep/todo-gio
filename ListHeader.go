package main

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type ListHeader struct {
	theme  *material.Theme
	widget *widget.Clickable
}

func NewListHeader(theme *material.Theme) *ListHeader {
	header := ListHeader{
		theme:  theme,
		widget: &widget.Clickable{},
	}

	return &header
}

func (header ListHeader) Layout(gtx layout.Context) layout.Dimensions {
	dims := layout.Flex{
		Axis: layout.Horizontal,
	}.Layout(
		gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return material.Body1(header.theme, "List Header").Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return material.Button(header.theme, header.widget, "click me").Layout(gtx)
		}),
	)

	if DebugLayout() {
		return DebugDimensions(gtx, dims, header.theme)
	}

	return dims
}
