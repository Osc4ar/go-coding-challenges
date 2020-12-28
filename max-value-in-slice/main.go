package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 12, 443, 51, 1, 0, 323}

	max := nums[0]

	for _, num := range nums[1:] {
		if num > max {
			max = num
		}
	}

	fmt.Printf("The max value in the slice is %d", max)
}
