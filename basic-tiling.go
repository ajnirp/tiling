package main

import (
	"math/rand"
	"time"

	"github.com/fogleman/gg"
)

const (
	canvasSide = 900
	cellSide   = 45
)

func main() {
	rand.Seed(time.Now().UnixNano())

	context := gg.NewContext(canvasSide, canvasSide)

	context.SetRGB(1, 1, 1)
	context.Clear()

	const numRows = canvasSide / cellSide

	context.SetRGB(0, 0, 0)
	context.SetLineWidth(1.0)

	for row := 0; row < numRows; row++ {
		for col := 0; col < numRows; col++ {
			// In each cell, the line goes from left to right. It either starts
			// from the top left or the bottom left, and ends respectively at
			// the bottom right or the top right.
			startsFromTop := true
			if rand.Intn(2) == 1 {
				startsFromTop = false
			}

			leftX := float64(col * cellSide)
			rightX := float64((col + 1) * cellSide)
			topY := float64(row * cellSide)
			bottomY := float64((row + 1) * cellSide)

			if startsFromTop {
				context.DrawLine(leftX, topY, rightX, bottomY)
			} else {
				context.DrawLine(leftX, bottomY, rightX, topY)
			}

			context.Stroke()
		}
	}

	context.SavePNG("tiling.png")
}
