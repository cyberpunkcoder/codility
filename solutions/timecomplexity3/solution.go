package solution

import "math"

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

// Solution to the problem
func Solution(A []int) int {
	sum, minimum := 0, 0

	for _, value := range A {
		sum += value
	}
	lhs := A[0]
	rhs := sum - lhs

	for i := 1; i < len(A); i++ {
		diff := int(math.Abs(float64(lhs) - float64(rhs)))

		if diff < minimum || i == 1 {
			minimum = diff
		}
		lhs += A[i]
		rhs -= A[i]
	}
	return minimum
}
