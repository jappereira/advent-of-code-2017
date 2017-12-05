package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	input := 289326

	side := int(math.Ceil(math.Sqrt(float64(input))))

	if side % 2 == 0 {
		side++
	}

	total := side * side
	radius := side / 2

	findNearestCenterPoint(input, side, radius, total - (side / 2))
}

func findNearestCenterPoint(input, side, radius, point int) {
	distanceToPoint := int(math.Abs(float64(input - point)))

	if distanceToPoint <= side / 2 {
		fmt.Println(distanceToPoint + radius)
		os.Exit(0)
	} else {
		findNearestCenterPoint(input, side, radius, point - (side - 1))
	}
}
