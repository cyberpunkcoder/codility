package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

// Solution to problem
func Solution(X int, A []int) int {
	leafs := make(map[int]bool)
	left := X

	for i := 0; i < len(A); i++ {
		if A[i] <= X && !leafs[A[i]] {
			left--
			if left == 0 {
				return i
			}
			leafs[A[i]] = true
		}
	}
	return -1
}
