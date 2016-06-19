package sol_test

import (
	"reflect"
	"testing"

	"github.com/lambrospetrou/go99/sol"
)

func Test01(t *testing.T) {
	l := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 22, 33, 11, -9}
	if res, _ := sol.P01(l); res != -9 {
		t.Error("Expected -9, got ", res)
	}

	if _, err := sol.P01([]int{}); err == nil {
		t.Error("Expected err, got nil err")
	}
}

func Test02(t *testing.T) {
	l := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 22, 33, 11, -9}
	if res, _ := sol.P02(l); !reflect.DeepEqual([]int{11, -9}, res) {
		t.Error("Expected [11, -9], got ", res)
	}

	if _, err := sol.P02([]int{1}); err == nil {
		t.Error("Expected err with nil list, got nil err")
	}

	if _, err := sol.P02([]int{}); err == nil {
		t.Error("Expected err with nil list, got nil err")
	}
}
