package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

// Solution to problem
func Solution(N int) int {
	max, count := 0, -1
	for N > 0 {
		if N%2 == 1 {
			if max < count {
				max = count
			}
			count = 0
		} else if count != -1 {
			count++
		}
		N /= 2
	}
	return max
}
