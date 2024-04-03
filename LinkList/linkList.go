package linklist

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
