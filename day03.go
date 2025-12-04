package main

import (
	"bufio"
	"fmt"
	"os"
)

func digitsToInt(digits []int) int {
	var final int

	for _, n := range digits {
		final = final*10 + n

	}

	return final
}
func main() {

	scanner := bufio.NewScanner(os.Stdin)
	part1 := 0
	part2 := 0

	for scanner.Scan() {
		line := scanner.Text()

		// For part 1 we find the biggest number that's not at the end of the list since the 10s place
		// dominates.
		firstMax := 0
		secondMax := 0
		for i, ch := range line {

			num := int(ch - '0')
			notLast := i != len(line)-1

			if firstMax < num && notLast {
				firstMax = num
				secondMax = 0 // Reset second max to 0 since it needs to always come after firstMax
				// Prevent assignment to secondMax
				continue
			}

			if secondMax < num {
				secondMax = num
			}

		}

		maxJolt := firstMax*10 + secondMax
		part1 += maxJolt

		nums := []int{}
		for _, ch := range line {
			num := int(ch - '0')
			nums = append(nums, num)
		}


		keptJolts := []int{}

		// can keep 12 digits.
		windowSize := len(nums) - 12 + 1
		i := 0
		for {
			if i == len(nums) {
				break
			}

			windowEnd := i + windowSize
			windowEnd = min(len(nums), windowEnd)

			// Find the biggest number and index inside the sliding window.
			window := nums[i:windowEnd]
			biggest := nums[i]
			biggestIndex := 0
			for j, windowNum := range window {
				if biggest < windowNum {
					biggest = windowNum
					biggestIndex = j
				}
			}

			keptJolts = append(keptJolts, biggest)
			windowSize -= biggestIndex

			i += 1 + biggestIndex
		}

		// Remove any trailing numbers if there were no good removal candidates.
		keptJolts = keptJolts[0 : 12]
		joltage := digitsToInt(keptJolts)
		part2 += joltage
	}

	fmt.Println("part1:", part1)
	fmt.Println("part2:", part2)
}
