package main

import (
	"fmt"
	"math"
	"reflect"
)

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	width  float64
	height float64
}

func main() {
	r := rectangle{width: 3, height: 4}
	c := circle{radius: 10}
	fmt.Printf("Area of rectangle: %f\n", getArea(r))
	fmt.Printf("Area of circle: %f\n", getArea(c))

	showInfo(r)
	showInfo(c)

	castToRectangle(c)
	castToRectangle(r)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s shape) float64 {
	return s.area()
}

func showInfo(s shape) {
	t := reflect.TypeOf(s).Name()
	switch t {
	case "circle":
		c := s.(circle)
		fmt.Printf("This is a circle with radius %f\n", c.radius)
		break
	case "rectangle":
		r := s.(rectangle)
		fmt.Printf("This is a rectangle with width %f and height %f\n", r.width, r.height)
		break
	}
}

func castToRectangle(s shape) {
	instanceCast, statusCast := s.(rectangle)
	if !statusCast {
		fmt.Println("Cast failed")
	} else {
		fmt.Println("Cast success")
		fmt.Printf("This is a rectangle with width %f and height %f\n", instanceCast.width, instanceCast.height)
	}
}
