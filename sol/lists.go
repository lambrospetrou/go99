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
