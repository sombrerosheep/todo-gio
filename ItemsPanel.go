package main

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
)

type ItemsPanel struct {
	theme *material.Theme

	header *ItemsHeader
	list   *ItemsList
}

func NewItemsPanel(theme *material.Theme) *ItemsPanel {
	lp := ItemsPanel{
		theme:  theme,
		header: NewItemsHeader(theme),
		list:   NewItemsList(theme, state.GetSelected()),
	}

	return &lp
}

func (ip *ItemsPanel) Frame() {
}

func (ip *ItemsPanel) Layout(gtx layout.Context) layout.Dimensions {
	dims := layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return material.Body1(ip.theme, "Items Panel").Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return ip.header.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return ip.list.Layout(gtx)
		}),
	)

	if debugLayout {
		return DebugDimensions(gtx, dims, ip.theme)
	}

	return dims
}
