package main

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
)

type ListRow struct{}

func (ls ListRow) Layout(gtx layout.Context, name string, theme *material.Theme) layout.Dimensions {
	return material.Body1(theme, name).Layout(gtx)
}
