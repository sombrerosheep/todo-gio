package main

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
)

type ListRow struct{}

func (ls ListRow) Layout(gtx layout.Context, name string, theme *material.Theme) layout.Dimensions {
	dims := material.Body1(theme, name).Layout(gtx)

	if debug.Enabled {
		return DebugDimensions(gtx, dims, theme)
	}

	return dims
}
