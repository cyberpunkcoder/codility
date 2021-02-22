package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

// Solution to problem
func Solution(A []int) int {
	distinct, nums := 0, make(map[int]bool)

	for _, e := range A {
		if !nums[e] {
			distinct++
		}
		nums[e] = true
	}
	return distinct
}
