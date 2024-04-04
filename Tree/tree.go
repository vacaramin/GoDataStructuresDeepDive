package tree

import "errors"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

type Tree struct {
	Root *Node
}

func (t *Tree) Insert(val int) {
	newNode := Node{Val: val, Left: nil, Right: nil}
	if t.Root == nil {
		t.Root = &newNode
		return
	}

}

func (t *Tree) Remove(val int) error {
	if t.Root == nil {
		return errors.New("remove failed the tree is empty")
	}

	temp := t.Root
	for {
		if temp.Val == val {
			if temp.Right != nil {

			}
			if temp.Left != nil {

			}

		}
	}
}
