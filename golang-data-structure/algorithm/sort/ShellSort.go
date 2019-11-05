package sort

import (
	"fmt"
)

type Shell struct {
}

func (s *Shell) Sort(a []interface{}) {
	h := 1
	cou := len(a)
	for h < (cou / 3) {
		h = 3*h + 1
	}
	for h >= 1 {
		for i := h; i < cou; i++ {
			for j := i; j >= h && s.Less(a[j], a[j-h]); j -= h {
				s.Exch(a, j, j-h)

			}
		}
		h = h / 3
	}
	// array := make([]int, 0)
	// for h < (cou / 3) {
	// 	array = append(array, h)
	// 	h = h*3 + 1
	// }
	// for i := len(array) - 1; i >= 0; i-- {
	// 	for j := i; j < cou; j++ {
	// 		for k := j; k >= i && s.Less(a[k], a[k-i]); k -= i {
	// 			s.Exch(a, k, k-i)
	// 		}
	// 	}
	//
	// }
}

func (*Shell) Less(a, b interface{}) bool {
	switch a.(type) {
	case float64:
		return CompareFloat64(a.(float64), b.(float64)) < 0
	case int:
		return CompareInt(a.(int), b.(int)) < 0
	case int64:
		return CompareInt64(a.(int64), b.(int64)) < 0
	case string:
		return CompareString(a.(string), b.(string)) < 0
	default:
		return false
	}
}

func (s *Shell) Exch(a []interface{}, i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (s *Shell) Show(a []interface{}) {
	for i := 0; i < len(a); i++ {
		fmt.Print(a[i], " ")
	}
	fmt.Println()
}

func (s *Shell) isSorted(a []interface{}) bool {
	for i := 1; i < len(a); i++ {
		if s.Less(a[i], a[i-1]) {
			return false
		}
	}
	return true
}
