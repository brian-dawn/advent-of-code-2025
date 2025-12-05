package main

import (
	"bufio"
	"fmt"
	"os"
)



type Point struct {
	X int
	Y int
}


func part1(world map[Point]bool, width int, height int) int {

	var totalAbleToMove int
	for x:= range width {
		for y := range height  {

			spot, found := world[Point{x, y}]
			if !found {
				continue
			}

			// No paper roll
			if !spot {
				continue
			}

			var surrounding int
			for i:= range 3 {
				for j := range 3  {

					if i == 1 && j == 1 {
						continue
					}

					neighborX, neighborY := x + i - 1, y + j - 1
					neighbor, neighborPresent := world[Point{neighborX, neighborY}]

					if !neighborPresent || !neighbor {
						continue
					}

					surrounding++

				}
			}

			if surrounding < 4 {
				totalAbleToMove++
			}


		}
	}

	return totalAbleToMove
}


func main() {




	scanner := bufio.NewScanner(os.Stdin)

	world := make(map[Point]bool)
	var width, y int

	for scanner.Scan() {
		line := scanner.Text()

		
		width = len(line)

		for x, ch := range line {
			switch ch {
case '.':
				world[Point{x, y}] = false
			case '@':
				world[Point{x, y}] = true
			}

			if width < x {
				width = x
			}
		}
		y++


	}


	p1Answer := part1(world, width, y)

	fmt.Println("part1:", p1Answer)


}
