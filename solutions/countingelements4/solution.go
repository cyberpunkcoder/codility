package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

// Solution to problem
func Solution(A []int) int {
	n := len(A)
	occurs := make([]bool, n)

	for _, e := range A {
		if e < 1 || e > n {
			return 0
		}
		if occurs[e-1] {
			return 0
		}
		occurs[e-1] = true
	}
	return 1
}
