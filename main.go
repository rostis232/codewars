package main

import (
	"fmt"
	"math"
)

func ListSquared(m, n int) [][]int {
	var answer [][]int
	for a := m; a <= n; a++ {
		var preanswer []int
		sumOfSquares := 0
		for b := 1; b <= a; b++ {
			if a%b == 0 {
				sumOfSquares = sumOfSquares + (b * b)
			}
		}
		if sumOfSquares != 0 && math.Mod(math.Sqrt(float64(sumOfSquares)), 1.0) == 0 {
			preanswer = append(preanswer, a, sumOfSquares)
			answer = append(answer, preanswer)
		}
	}
	if len(answer) == 0 {
		fmt.Println("Return nil")
		return nil

	}
	fmt.Println("Return answer")
	return answer

}

func main() {
	m, n := 300, 600

	fmt.Println(ListSquared(m, n))

}
