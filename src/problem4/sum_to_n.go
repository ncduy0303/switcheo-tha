package main

import "fmt"

// Implementation A: Recursive approach
// Time complexity: O(n)
// Space complexity: O(n)
func sum_to_n_a(n int) int {
	if n == 0 {
        return 0
    }
    return n + sum_to_n_c(n - 1)
}

// Implementation B: Iterative approach with a single loop
// Time complexity: O(n)
// Space complexity: O(1)
func sum_to_n_b(n int) int {
    sum := 0
    for i := 1; i <= n; i++ {
        sum += i
    }
    return sum
}

// Implementation C: Mathematical approach with arithmetic progression
// Time complexity: O(1)
// Space complexity: O(1)
func sum_to_n_c(n int) int {
	return n * (n + 1) / 2
}

func main() {
	// Test the implementations, should all print 15
    n := 5
    fmt.Println("sum_to_n_a:", sum_to_n_a(n))
    fmt.Println("sum_to_n_b:", sum_to_n_b(n))
    fmt.Println("sum_to_n_c:", sum_to_n_c(n))
}