package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

type Line struct {
	Begin, End Point
}

type Path []Point


func (l Line) Distance() float64{
	return math.Hypot(l.End.X - l.Begin.X, l.End.Y - l.Begin.Y)
}

func (l *Line) ScaleByX(f float64){
	l.End.X += (f-1)*(l.End.X - l.Begin.X)
	l.End.Y += (f-1)*(l.End.Y - l.Begin.Y)
}

func main(){
	side := Line{Point{1,2}, Point{4,6}}
	fmt.Println(side.Distance())
	side.ScaleByX(2)
	fmt.Println(side.Distance())

}