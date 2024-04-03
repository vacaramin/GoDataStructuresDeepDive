package linklist

import (
	"errors"
	"fmt"
)

type Node struct {
	val  int
	Next *Node
}

type LinkList struct {
	head *Node
}

func (l *LinkList) Add(val int) {
	newNode := Node{val: val, Next: nil}
	if l.head != nil {
		temp := l.head
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = &newNode
	} else {
		l.head = &newNode
	}

}
func (l *LinkList) Remove(val int) error {
	if l.head == nil {
		return errors.New("can't remove, the linked list is empty")
	}
	if l.head.val == val {
		l.head = l.head.Next
		return nil
	}

	if l.head != nil {
		temp := l.head
		for temp.Next != nil {
			if temp.Next.val == val {
				temp.Next = temp.Next.Next
				return nil
			}
			temp = temp.Next
		}
		return errors.New("element not found in link list")
	} else {
		return errors.New("the list is empty can't remove the element")
	}
}
func (l *LinkList) Print(val int) {
	fmt.Print("[ ")
	if l.head != nil {
		temp := l.head
		for temp.Next != nil {
			fmt.Print(temp.val, ", ")
			temp = temp.Next
		}

	}
	fmt.Print(" ]")

}
