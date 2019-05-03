package sort

import (
	"fmt"
)

type SelectionSort struct {
}

func (s *SelectionSort) Sort(a []interface{}) {
	for i := 0; i < len(a); i++ {
		min := i
		for j := i + 1; j < len(a); j++ {
			if s.Less(a[j], a[min]) {
				min = j
			}
		}
		s.Exch(a, i, min)
	}
}

func (s *SelectionSort) Less(a, b interface{}) bool {
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

func (s *SelectionSort) Exch(a []interface{}, i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (s *SelectionSort) Show(a []interface{}) {
	for i := 0; i < len(a); i++ {
		fmt.Print(a[i], " ")
	}
	fmt.Println()
}

func (s *SelectionSort) isSorted(a []interface{}) bool {
	for i := 1; i < len(a); i++ {
		if s.Less(a[i], a[i-1]) {
			return false
		}
	}
	return true
}
