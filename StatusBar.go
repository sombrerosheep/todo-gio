package main

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
)

type StatusBar struct {
	theme *material.Theme
}

func (stat StatusBar) Layout(gtx layout.Context) layout.Dimensions {
	return layout.Inset{
		Top:    2,
		Bottom: 2,
		Left:   5,
		Right:  5,
	}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return material.Body2(stat.theme, state.GetStatus()).Layout(gtx)
	})
}
