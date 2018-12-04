// Package aoclib contains utility functions for working with strings.
package aoclib

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// Contains tells whether a contains x.
func Contains(a []int, x int) bool {
	// TODO: general type? not just int?
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
