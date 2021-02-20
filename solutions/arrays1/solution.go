package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

// Solution to problem
func Solution(A []int, K int) []int {
	B := make([]int, len(A))
	for i := 0; i < len(A); i++ {
		B[(i+K)%len(A)] = A[i]
	}
	return B
}
