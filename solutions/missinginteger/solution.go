package solution

// you can also use imports, for example:
// import "fmt"
// import "os"
import "sort"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

// Solution to problem
func Solution(A []int) int {
	retval := 0
	sort.Ints(A)

	for i := 0; i < len(A); i++ {
		if A[i] > retval {
			retval++
			if A[i] == retval {
				continue
			}
			return retval
		}
	}
	return retval + 1
}
