package sol

import (
	"errors"
	"fmt"
)

// P01 returns the last element of a list or error
func P01(l []int) (int, error) {
	if len(l) == 0 {
		return -1, errors.New("No elements in list")
	}
	return l[len(l)-1], nil
}

// P02 returns the last two elements of a list or error
func P02(l []int) ([]int, error) {
	if len(l) < 2 {
		return nil, errors.New("List has less than 2 elements")
	}
	return l[len(l)-2:], nil
}

// P03 returns the n-th element
func P03(l []int, n int) (int, error) {
	if n < 0 {
		return -1, errors.New("No negative n allowed")
	}
	if len(l)-1 < n {
		return -1, errors.New(fmt.Sprintf("List has less than %d elements", n))
	}
	return l[n], nil
}

// P04 returns the length of a list
func P04(l []int) int {
	return len(l)
}

// P05 reverses a list and returns a new copy of the reversed
func P05(l []int) []int {
	if l == nil {
		return nil
	}
	rev := make([]int, len(l))
	revIdx := len(l) - 1
	for _, v := range l {
		rev[revIdx] = v
		revIdx -= 1
	}
	return rev
}

// P05_1 reverses the list in place
func P05_1(l []int) []int {
	length := len(l) - 1
	for i := 0; i < length/2; i++ {
		l[i], l[length-i] = l[length-i], l[i]
	}
	return l
}
