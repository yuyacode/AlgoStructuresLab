package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

// 開始点と終了点をそれぞれ p1 と p2 で受け取る
// depthは深さ
func KochCurve(p1, p2 Point, depth int) []Point {
	if depth == 0 {
		return []Point{p1, p2}
	}

	dx := p2.x - p1.x  // 100 - 0
	dy := p2.y - p1.y  // 0 - 0

	a := Point{p1.x + dx/3, p1.y + dy/3}  // 3分割した際の１点目の座標
	c := Point{p1.x + 2*dx/3, p1.y + 2*dy/3}  // 3分割した際の2点目の座標

	rad := math.Pi * 60.0 / 180.0  // 60度の角をラジアン単位で表す
	b := Point{
		x: (c.x-a.x)*math.Cos(rad)-(c.y-a.y)*math.Sin(rad) + a.x,
		y: (c.x-a.x)*math.Sin(rad)+(c.y-a.y)*math.Cos(rad) + a.y,
	}

	points := KochCurve(p1, a, depth-1)
	points = append(points, KochCurve(a, b, depth-1)...)
	points = append(points, KochCurve(b, c, depth-1)...)
	points = append(points, KochCurve(c, p2, depth-1)...)
	return points
}

func KochCurveMain() {
	p1 := Point{0, 0}
	p2 := Point{100, 0}

	depth := 1

	points := KochCurve(p1, p2, depth)
	for _, point := range points {
		fmt.Printf("(%.2f, %.2f)\n", point.x, point.y)
	}
}
