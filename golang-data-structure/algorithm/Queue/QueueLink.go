package Queue

import (
	"algorithm/link"
)

type Queue interface {
	IsEmpty() bool
	Size() int
	Enqueue(value interface{})
	Dequeue() (interface{}, bool)
}

type queueLink struct {
	first *link.Node
	last  *link.Node
	len   int
}

func (q *queueLink) IsEmpty() bool {
	return q.len == 0
}

func (q *queueLink) Size() int {
	return q.len
}

func (q *queueLink) Enqueue(value interface{}) {
	if q.IsEmpty() {
		q.first = link.NewLinkedList(value)
		q.last = q.first
	} else {
		q.last.Next = link.NewLinkedList(value)
		q.last = q.last.Next
	}
	q.len++
}

func (q *queueLink) Dequeue() (interface{}, bool) {
	if q.first == nil {
		return nil, false
	} else {
		temp := *q.first
		q.first = q.first.Next
		q.len--
		if q.IsEmpty() {
			q.last = nil
		}
		return temp, true
	}
}

func NewQueueLink() *queueLink {
	return &queueLink{}
}
