package sol

import "errors"

func P01(l []int) (int, error) {
	if len(l) == 0 {
		return -1, errors.New("No elements in list")
	}
	return l[len(l)-1], nil
}
