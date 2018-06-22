package main

import (
	"fmt"
)

func main() {
	var sum, sumOfSquare int
	for i := 1; i <= 100; i++ {
		sum = sum + i
		sumOfSquare = sumOfSquare + (i * i)
	}

	squareOfSum := sum * sum
	diff := squareOfSum - sumOfSquare
	fmt.Println(diff)
}
