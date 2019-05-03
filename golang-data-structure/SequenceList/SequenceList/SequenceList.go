package SequenceList

import (
	"errors"
	"fmt"
)

type SequenceList interface {
	Compare(element interface{}) ([]int, error)
	Insert(index int, value interface{}) error
	Append(value interface{})
	Delete(index int) error
	Sort()
	Merge(a sequenceList)
	Println()
}

//任何类型的线性表
type sequenceList struct {
	//数组
	elem []interface{}
}

//返回一个默认长度的空线性表
func NewDefaultSizeSequenceList() sequenceList {
	return sequenceList{
		elem: make([]interface{}, 0),
	}
}

//返回一个自定长度的空线性表
func NewSequenceList(size int) sequenceList {
	return sequenceList{
		elem: make([]interface{}, size),
	}
}

//判断线性表是否为空 为空返回true 否则范围false
func (slist *sequenceList) IsEmpty() bool {
	if len(slist.elem) == 0 {
		return true
	}
	return false
}

//返回当前线性表长度
func (slist *sequenceList) ListSize() int {
	return len(slist.elem)
}

//返回索引元素，如果索引大于当前数组长度或者小于0 则返回-1和一个error
func (slist *sequenceList) GetElement(index int) (interface{}, error) {
	if index > len(slist.elem) || index < 0 {
		return -1, errors.New("The index is invalid")
	}
	return slist.elem[index], nil
}

//返回线性表中与参数值相等的一个或多个的索引号，返回一个索引数组，如果没有相等则返回nil和error
func (slist *sequenceList) Compare(element interface{}) ([]int, error) {
	index := make([]int, 0)
	for key, value := range slist.elem {
		if value == element {
			index = append(index, key)
		}
	}
	if len(index) == 0 {
		return nil, errors.New("No equal value")
	}
	return index, nil
}

//在某个索引前插入一个值 若想追加一个值或者在空线性表中添加值请使用Append()
func (slist *sequenceList) Insert(index int, value interface{}) error {
	if index <= 0 || index > len(slist.elem) {
		return errors.New("The index is invalid")
	}
	temp := make([]interface{}, 0)
	for i := 0; i < slist.ListSize(); i++ {
		if i == index {
			temp = append(temp, value)
		}
		temp = append(temp, slist.elem[i])
	}
	slist.elem = temp
	return nil
}

// 在线性表中追加一个值
func (slist *sequenceList) Append(value interface{}) {
	slist.elem = append(slist.elem, value)
}

//在线性表中删除对应索引的值
func (slist *sequenceList) Delete(index int) error {
	if index < 0 || index > len(slist.elem) {
		return errors.New("The index is invalid")
	}
	temp := make([]interface{}, 0)
	for i := 0; i < slist.ListSize(); i++ {
		if i == index {
			continue
		}
		temp = append(temp, slist.elem[i])
	}
	slist.elem = temp
	return nil

}

//升序排列线性表int类型
func (slist *sequenceList) Sort() {
	for i := 0; i < slist.ListSize(); i++ {
		for j := 0; j < slist.ListSize(); j++ {
			if slist.elem[i].(int) < slist.elem[j].(int) {
				temp := slist.elem[i]
				slist.elem[i] = slist.elem[j]
				slist.elem[j] = temp
			}
		}
	}
}

//合并两个线性表
func (slist *sequenceList) Merge(a sequenceList) {
	for value := range a.elem {
		a.elem = append(a.elem, value)
	}
	slist.elem = a.elem
	slist.Sort()
}

func (slist *sequenceList) Println() {
	fmt.Println(slist.elem[:])
}
