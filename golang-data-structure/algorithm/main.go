package main

import (
	"algorithm/sort"
	"fmt"
	"time"
)

func main() {
	array1 := []interface{}{
		55, 94, 87, 1, 4, 32, 11, 77, 39, 42, 64, 53, 70, 12, 9,
	}
	current := time.Now().UnixNano()
	c := sort.Shell{}
	c.Sort(array1)
	fmt.Println(time.Now().UnixNano() - current)
}
