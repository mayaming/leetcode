package main

import "reflect"
import "fmt"
import "errors"

// https://leetcode.com/problems/max-points-on-a-line/description/

type Point struct {
	X int
	Y int
}

type Line struct {
	lefter  Point
	righter Point
	pts  []Point
}

func (line *Line) getId() string {
	return fmt.Sprintf("(%d,%d)-(%d,%d)", line.lefter.X, line.lefter.Y, line.righter.X, line.righter.Y)
}

func MakeLine(p1, p2 Point) (Line, error) {
	if reflect.DeepEqual(p1, p2) {
		return Line{}, errors.New("same points")
	}

	if p1.X < p2.X || (p1.X == p2.X && p1.Y <= p2.Y) {
		return Line{p1, p2, make([]Point, 0)}, nil
	} else {
		return Line{p2, p1, make([]Point, 0)}, nil
	}
}

func (l *Line) SameLine(p Point) bool {

	vsLeftX := p.X - l.lefter.X
	vsLeftY := p.Y - l.lefter.Y
	vsRightX := p.X - l.righter.X
	vsRightY := p.Y - l.righter.Y
	return vsLeftX*vsRightY == vsLeftY*vsRightX
}

func (l *Line) AddPts(p Point) {
	l.pts = append(l.pts, p)
}

func maxPoints(points []Point) int {
	lines := make(map[string](*Line))
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			line, err := MakeLine(points[i], points[j])
			if err == nil {
				id := line.getId()
				lines[id] = &line
			}
		}
	}

	for i := 0; i < len(points); i++ {
		p := points[i]
		for _, v := range lines {
			if v.SameLine(p) {
				v.AddPts(p)
			}
		}
	}

	cnt := 0
	for _, v := range lines {
		if len(v.pts) > cnt {
			cnt = len(v.pts)
		}
	}

	if cnt == 0 {
		cnt = len(points)
	}

	return cnt
}

func main() {
	points := []Point{{1,1}, {3,2}, {5,3}, {4,1}, {2,3}, {1,4}}
	fmt.Println(maxPoints(points))
	points = []Point{{1,1}, {2,2}, {3,3}}
	fmt.Println(maxPoints(points))
	points = []Point{{1,1}}
	fmt.Println(maxPoints(points))
	points = []Point{{1,1}, {1,1}}
	fmt.Println(maxPoints(points))
	points = []Point{{1,1}, {1,1}, {1,1}}
	fmt.Println(maxPoints(points))
}