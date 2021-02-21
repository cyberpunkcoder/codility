package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

// Solution to the problem
func Solution(X int, Y int, D int) int {
	return ((Y - X) + (D - 1)) / D
}
