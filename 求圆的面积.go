package main

import (
	"fmt"
	"math"
)

func main() {
	var radius float64
	fmt.Scan(&radius)
	area := math.Pi * radius * radius
	fmt.Printf("半径为%.2f的圆的面积是%.2f\n", radius, area)
}
