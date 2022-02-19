// Package mymath2 provides ACME inc math solutions.
package mymath2

// Sum adds an unlimited number of values of type int.
func Sum(arr ...int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}
