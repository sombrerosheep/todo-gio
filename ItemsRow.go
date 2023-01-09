package main

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
)

type ItemsRow struct {
}

func (row ItemsRow) Layout(gtx layout.Context, item Item, theme *material.Theme) layout.Dimensions {
	dims := material.Body1(theme, item.String()).Layout(gtx)

	if DebugLayout() {
		return DebugDimensions(gtx, dims, theme)
	}

	return dims
}
