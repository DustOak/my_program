package Stack

type Stack interface {
	Push(interface{})
	Pop() (interface{}, bool)
	IsEmpty() bool
	Size() int
	GetSlice() []interface{}
}
