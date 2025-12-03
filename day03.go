package main

import (
	"bufio"
	"fmt"
	"os"
)




func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		// For part 1 we find the biggest number that's not at the end of the list since the 10s place
		// dominates.
		firstMax := 0
		secondMax := 0
		for i, ch := range(line) {


			num := int(ch - '0')
			notLast := i != len(line) - 1

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

		fmt.Println(firstMax, secondMax)



	}

	

}
