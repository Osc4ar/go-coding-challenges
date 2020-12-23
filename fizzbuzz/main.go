package main

import (
	"fmt"
	"strconv"
)

func printFizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		output := ""
		if i % 3 == 0 {
			output = "fizz"
		}
		if i % 5 == 0 {
			if output != "" {
				output += " "
			}
			output += "buzz"
		}
		if output == "" {
			output = strconv.Itoa(i)
		}
		fmt.Println(output)
	}
}

func main() {
	printFizzBuzz(100000)
}
