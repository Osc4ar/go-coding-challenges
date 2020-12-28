package main

import "fmt"

func main() {
	count := 0
	for i := 1000; i <= 9999; i++ {
		for j := i; j <= 9999; j++ {
			result := fmt.Sprintf("%d", i * j);
			if result[0] == result[len(result)-1] {
				count++
			}
		}
	}
	fmt.Printf("There are %d even-ended numbers between 1000 and 9999", count)
}

