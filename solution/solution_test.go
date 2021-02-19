package solution

import (
	"fmt"
	"testing"
	"time"
)

// Add "go.testFlags": ["-v"], to settings.json to view output

func TestSolution(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		args args
		want int
	}{
		// Put tests in here, example below:
		//{args{1}, 0},
	}
	for i, tt := range tests {
		// Generate  a name for the test
		name := fmt.Sprint("Test", i+1)

		// Run test
		t.Run(name, func(t *testing.T) {

			// Record runtime
			start := time.Now()
			got := Solution(tt.args.N)
			duration := time.Since(start)

			// Print exceptions to expected value
			if got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}

			// Print the arguments, the returns and runtime
			fmt.Println("Gave:", tt.args.N)
			fmt.Println("Got:", got)
			fmt.Println("Time:", duration)
		})
	}
}

func TestSolution2(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		args args
		want int
	}{
		// Put tests in here, example below:
		//{args{1}, 0},
	}
	for i, tt := range tests {
		// Generate  a name for the test
		name := fmt.Sprint("Test", i+1)

		// Run test
		t.Run(name, func(t *testing.T) {

			// Record runtime
			start := time.Now()
			got := Solution2(tt.args.N)
			duration := time.Since(start)

			// Print exceptions to expected value
			if got != tt.want {
				t.Errorf("Solution2() = %v, want %v", got, tt.want)
			}

			// Print the arguments, the returns and runtime
			fmt.Println("Gave:", tt.args.N)
			fmt.Println("Got:", got)
			fmt.Println("Time:", duration)
		})
	}
}
