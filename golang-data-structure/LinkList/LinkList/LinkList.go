package LinkList

import (
	"errors"
	"fmt"
)

type LinkList interface {
	Append(value interface{})
	Delete(value interface{})
	Traversing()
	Search(value interface{}) (interface{}, error)
	Insert(index int, value interface{})
}

type linklist struct {
	//值
	value interface{}
	//下一个节点的地址
	next *linklist
}

// 创建一个新节点
func newNode(value interface{}) linklist {
	return linklist{
		value: value,
		next:  nil,
	}
}

//创建一个链表 value=nil的节点为head节点 next直接指向value等于所传参数的节点
func NewLinkList(value interface{}) linklist {
	node := linklist{
		value: nil,
		next: &linklist{
			value: value,
			next:  nil,
		},
	}
	return node
}

//在尾部添加一个新节点
func (llist *linklist) Append(value interface{}) {
	for current := llist; ; current = current.next {
		if current.next == nil {
			newNode := newNode(value)
			current.next = &newNode
			return
		}
	}
}

//删除一个于所传参数相等值的节点
func (llist *linklist) Delete(value interface{}) {
	for current := llist; ; current = current.next {
		if current.next == nil {
			return
		}
		if current.next.value == value {
			if current.next.next == nil {
				current.next = nil
				return
			}
			current.next = current.next.next
		}

	}
}

//遍历链表
func (llist *linklist) Traversing() {
	for current := llist; ; current = current.next {
		fmt.Print(current.value, " ")
		if current.next == nil {
			break
		}
	}
	fmt.Println()
}

//查找链表
func (llist *linklist) Search(value interface{}) ([]interface{}, error) {
	result := make([]interface{}, 0)
	for current := llist; ; current = current.next {
		if current.value == value {
			result = append(result, current.value)
		}
		if current.next == nil {
			return result, nil
		}
	}
	return nil, errors.New("No result")
}

//在链表中索引号前插入一个节点 启示索引为0  不包含head节点
func (llist *linklist) Insert(index int, value interface{}) {

	for current, i := llist, 0; ; current, i = current.next, i+1 {
		if i == index {
			newNode := newNode(value)
			next := current.next
			newNode.next = next
			current.next = &newNode
		}
		if current.next == nil {
			return
		}
	}
}
