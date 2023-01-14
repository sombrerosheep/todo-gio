package main

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
)

type ListSelector struct {
	theme *material.Theme

	rows []*ListRow
}

func NewListSelector(listNames []string, rowSelected chan<- string, theme *material.Theme) *ListSelector {
	selector := ListSelector{
		theme: theme,
	}

	selector.rows = make([]*ListRow, len(listNames))

	for i, v := range listNames {
		selector.rows[i] = NewListRow(v, rowSelected, theme)
	}

	return &selector
}

func (ls ListSelector) Layout(gtx layout.Context) layout.Dimensions {
	list := layout.List{Axis: layout.Vertical}

	dims := list.Layout(gtx, len(ls.rows), func(gtx layout.Context, i int) layout.Dimensions {
		return ls.rows[i].Layout(gtx)
	})

	if DebugLayout() {
		return DebugDimensions(gtx, dims, ls.theme)
	}

	return dims
}
