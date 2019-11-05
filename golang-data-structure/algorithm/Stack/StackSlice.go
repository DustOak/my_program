package Stack

type stackslice struct {
	values []interface{}
	len    int
}

func (s *stackslice) Push(v interface{}) {
	s.values = append(s.values, v)
	s.len++
}

func (s *stackslice) Pop() interface{} {
	s.len--
	temp := s.values[len(s.values)-1]
	s.values[len(s.values)-1] = nil
	return temp
}

func (s *stackslice) IsEmpty() bool {
	return s.len == 0
}

func (s *stackslice) Size() int {
	return s.len
}

func (s *stackslice) GetSlice() []interface{} {
	return s.values
}
func NewFixedCapacityStack(cap int) *stackslice {
	return &stackslice{
		values: make([]interface{}, cap),
		len:    0,
	}
}
