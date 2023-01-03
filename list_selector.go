package main

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
)

func ListSelections(state *State, theme *material.Theme, gtx layout.Context) layout.FlexChild {
	list := layout.List{Axis: layout.Vertical}

	return layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
		names := state.GetListKeys()
		return list.Layout(gtx, len(names), func(gtx layout.Context, i int) layout.Dimensions {
			return ListSelector(names[i], theme, gtx)
		})
	})
}

func ListSelector(name string, theme *material.Theme, gtx layout.Context) layout.Dimensions {
	return material.Body1(theme, name).Layout(gtx)
}
