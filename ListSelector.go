package main

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
)

type ListSelector struct{}

func (ls ListSelector) Layout(gtx layout.Context, listNames []string, theme *material.Theme) layout.Dimensions {
	list := layout.List{Axis: layout.Vertical}

	dims := list.Layout(gtx, len(listNames), func(gtx layout.Context, i int) layout.Dimensions {
		return ListRow{}.Layout(gtx, listNames[i], theme)
	})

	if DebugLayout() {
		return DebugDimensions(gtx, dims, theme)
	}

	return dims
}
