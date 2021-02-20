package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

// Solution to problem
func Solution(A []int) int {
	nums := make(map[int]int)
	var retval int

	for i := 0; i < len(A); i++ {
		nums[A[i]]++
		if nums[A[i]]%2 == 0 {
			retval -= A[i]
			continue
		}
		retval += A[i]
	}
	return retval
}
