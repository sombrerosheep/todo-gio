package main

import (
	"fmt"
	"image/color"

	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

const borderOpacity uint8 = 180

var borderSeed = 0

func getBorderColor() color.NRGBA {
	colors := []color.NRGBA{
		color.NRGBA{R: 200, G: 0, B: 0, A: borderOpacity},
		color.NRGBA{R: 0, G: 200, B: 0, A: borderOpacity},
		color.NRGBA{R: 0, G: 0, B: 200, A: borderOpacity},
	}

	if borderSeed >= len(colors) {
		borderSeed = 0
	}

	c := colors[borderSeed]

	borderSeed = borderSeed + 1

	return c
}

const (
	debugLayout bool    = true
	debugWidth  unit.Dp = 3
)

func DebugLayout() bool {
	return debugLayout
}

func DebugDimensions(gtx layout.Context, dims layout.Dimensions, theme *material.Theme) layout.Dimensions {
	debugColor := getBorderColor()

	return layout.Stack{}.Layout(
		gtx,
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return widget.Border{
				Color:        debugColor,
				Width:        debugWidth,
				CornerRadius: unit.Dp(5),
			}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return dims
			})
		}),
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			return layout.UniformInset(debugWidth).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				debugDims := material.Body1(theme, fmt.Sprintf("(%d x %d)", dims.Size.X, dims.Size.Y))
				debugDims.Alignment = text.Alignment(layout.End)
				debugDims.Color = debugColor
				debugDims.TextSize = unit.Sp(10)

				return debugDims.Layout(gtx)
			})
		}),
	)
}
