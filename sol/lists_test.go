package sol_test

import (
	"fmt"
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

func Test03(t *testing.T) {
	l := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 22, 33, 11, -9}
	if res, _ := sol.P03(l, 3); res != 4 {
		t.Error("Expected 4, got ", res)
	}

	if _, err := sol.P03([]int{}, 1); err == nil {
		t.Error("Expected err, got nil err")
	}

	if _, err := sol.P03([]int{1, 2, 3, 4}, -1); err == nil {
		t.Error("Expected err, got nil err")
	}
}

func Test04(t *testing.T) {
	l := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 22, 33, 11, -9}
	if sol.P04(l) != 13 {
		t.Error("Expected 13, got ", sol.P04(l))
	}

	if sol.P04([]int{}) != 0 {
		t.Error("Expected 0, got ", sol.P04([]int{}))
	}

	if sol.P04(nil) != 0 {
		t.Error("Expected 0, got ", sol.P04(nil))
	}
}

func Test05(t *testing.T) {
	l := []int{1, 2, 3, 4, 5, 11, -9}
	lrev := []int{-9, 11, 5, 4, 3, 2, 1}
	if !reflect.DeepEqual(lrev, sol.P05(l)) {
		t.Error(fmt.Sprintf("Expected %v, got %v", lrev, sol.P05(l)))
	}

	if !reflect.DeepEqual([]int{}, sol.P05([]int{})) {
		t.Error(fmt.Sprintf("Expected %v, got %v", []int{}, sol.P05([]int{})))
	}

	if len(sol.P05(nil)) != 0 {
		t.Error(fmt.Sprintf("Expected len 0, got %v", len(sol.P05(nil))))
	}
}

func Test05_1(t *testing.T) {
	l := []int{1, 2, 3, 4, 5, 11, -9}
	lrev := []int{-9, 11, 5, 4, 3, 2, 1}
	if !reflect.DeepEqual(lrev, sol.P05_1(l)) {
		t.Error(fmt.Sprintf("Expected %v, got %v", lrev, sol.P05_1(l)))
	}

	if !reflect.DeepEqual([]int{}, sol.P05_1([]int{})) {
		t.Error(fmt.Sprintf("Expected %v, got %v", []int{}, sol.P05_1([]int{})))
	}

	if len(sol.P05_1(nil)) != 0 {
		t.Error(fmt.Sprintf("Expected len 0, got %v", len(sol.P05_1(nil))))
	}
}

func Test06(t *testing.T) {
	lyes := []int{1, 2, 3, 4, 3, 2, 1}
	lno := []int{-9, 11, 5, 4, 3, 2, 1}

	if !sol.P06(lyes) {
		t.Error(fmt.Sprintf("Expected %v, to be true", lyes))
	}

	if !sol.P06([]int{4}) {
		t.Error(fmt.Sprintf("Expected single element to be true"))
	}

	if !sol.P06([]int{4, 5, 5, 4}) {
		t.Error(fmt.Sprintf("Expected even length list to be true"))
	}

	if sol.P06(lno) {
		t.Error(fmt.Sprintf("Expected %v, to be false", lno))
	}

	if !sol.P06([]int{}) {
		t.Error(fmt.Sprintf("Expected empty list to be true"))
	}

	if !sol.P06(nil) {
		t.Error(fmt.Sprintf("Expected nil to be true"))
	}
}
