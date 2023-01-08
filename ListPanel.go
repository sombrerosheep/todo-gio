package main

import (
	"sort"

	"gioui.org/layout"
	"gioui.org/widget/material"
)

type ListPanel struct {
	theme *material.Theme

	listHeader *ListHeader
}

func NewListPanel(theme *material.Theme) *ListPanel {
	lp := ListPanel{
		theme:      theme,
		listHeader: NewListHeader(theme),
	}

	return &lp
}

func (lp *ListPanel) Frame() {
}

func (lp ListPanel) Layout(gtx layout.Context) layout.Dimensions {
	dims := layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.N.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return lp.listHeader.Layout(gtx)
			})
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			names := state.GetListKeys()
			sort.Strings(names)

			return ListSelector{}.Layout(gtx, names, lp.theme)

		}),
	)

	if DebugLayout() {
		return DebugDimensions(gtx, dims, lp.theme)
	}

	return dims
}
