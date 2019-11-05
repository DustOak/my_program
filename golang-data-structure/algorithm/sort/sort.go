package sort

type Sort1 interface {
	Sort(a []interface{})
	Less(a, b interface{}) bool
	Exch(a []interface{}, i, j int)
	Show(a []interface{})
	isSorted(a []interface{}) bool
}
