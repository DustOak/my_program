package Stack

import (
	"algorithm/link"
)

type stacklink struct {
	top *link.Node
	len int
}

func (s *stacklink) Push(value interface{}) {
	temp := s.top
	a := link.NewLinkedList(value)
	a.Next = temp
	s.top = a
	s.len++
}

func (s *stacklink) Pop() (interface{}, bool) {
	temp := s.top
	if temp.Next != nil {
		s.top = s.top.Next
		s.len--
		return temp, true
	}
	return nil, false
}

func (s *stacklink) IsEmpty() bool {
	return s.len == 0
}

func (s *stacklink) Size() int {
	return s.len
}

func (s *stacklink) GetSlice() []interface{} {
	slice := make([]interface{}, 0)
	for current := s.top; current.Next != nil; current = current.Next {
		slice = append(slice, *current)
	}
	return slice
}
func NewStackLink() *stacklink {
	return &stacklink{
		top: link.NewLinkedList(nil),
		len: 0,
	}
}
