package main

import (
	"sort"

	"gioui.org/layout"
	"gioui.org/widget/material"
)

type ListPanel struct {
	theme *material.Theme

	listHeader   *ListHeader
	listSelector *ListSelector
}

func NewListPanel(rowSelected chan<- string, theme *material.Theme) *ListPanel {
	names := state.GetListKeys()
	sort.Strings(names)

	lp := ListPanel{
		theme:        theme,
		listHeader:   NewListHeader(theme),
		listSelector: NewListSelector(names, rowSelected, theme),
	}

	return &lp
}

func (lp ListPanel) Layout(gtx layout.Context) layout.Dimensions {
	dims := layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.N.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return lp.listHeader.Layout(gtx)
			})
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return lp.listSelector.Layout(gtx)
		}),
	)

	if DebugLayout() {
		return DebugDimensions(gtx, dims, lp.theme)
	}

	return dims
}
