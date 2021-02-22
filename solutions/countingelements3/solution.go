package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

// Solution to problem
func Solution(A []int) int {
	nums := make(map[int]bool)

	for i := 0; i < len(A); i++ {
		nums[A[i]] = true
	}
	for j := 1; j <= len(A)+1; j++ {
		if !nums[j] {
			return j
		}
	}
	return 1
}
