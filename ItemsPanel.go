package main

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
)

type ItemsPanel struct {
	theme *material.Theme
}

func NewItemsPanel(theme *material.Theme) *ItemsPanel {
	lp := ItemsPanel{
		theme: theme,
	}

	return &lp
}

func (ip *ItemsPanel) Frame() {
}

func (ip *ItemsPanel) Layout(gtx layout.Context) layout.Dimensions {
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.N.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return material.Body1(ip.theme, "Items Panel").Layout(gtx)
			})
		}),
	)
}
