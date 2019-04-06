package logic

import "math"

func SearchFibonacci(data []int, size int) {
}

// IsSquare returns true if the given integer is a perfect square.
func IsSquare(x int) bool {
	root := int(math.Sqrt(float64(x)))
	return root*root == x
}

// IsFibonacci returns true if the given integer is part of the fibonacci sequence.
//
// Source: https://www.geeksforgeeks.org/check-number-fibonacci-number/
func IsFibonacci(x int) bool {
	if x == 0 {
		return false
	}
	return IsSquare(5*x*x+4) || IsSquare(5*x*x-4)
}
