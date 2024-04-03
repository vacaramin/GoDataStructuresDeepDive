package stack

type Node struct {
	Value int
	Next  *Node
}

type Stack struct {
	Head *Node
}

func NewStack() *Stack {
	return &Stack{Head: nil}
}

func (s *Stack) Push(val int) {
	newNode := Node{Value: val, Next: nil}
	if s.Head != nil {
		newNode.Next = s.Head
	}
	s.Head = &newNode
}

func (s *Stack) Pop() (int, error) {
	if s.Head == nil {
		return -50, nil
	}
	val := s.Head.Value
	s.Head = s.Head.Next
	return val, nil
}

func (s *Stack) Peek() (int, error) {
	if s.Head == nil {
		return -50, nil
	}
	return s.Head.Value, nil
}

func (s *Stack) IsEmpty() bool {
	return s.Head == nil
}

func (s *Stack) Print() {
	temp := s.Head
	for temp != nil {
		print(temp.Value, " ")
		temp = temp.Next
	}
}

func (s *Stack) Size() int {
	temp := s.Head
	count := 0
	for temp != nil {
		count++
		temp = temp.Next
	}
	return count
}

func (s *Stack) Clear() {
	s.Head = nil
}

func (s *Stack) Search(val int) bool {
	temp := s.Head
	for temp != nil {
		if temp.Value == val {
			return true
		}
		temp = temp.Next
	}
	return false
}

func (s *Stack) Reverse() {
	if s.Head == nil || s.Head.Next == nil {
		return
	}
	var prev, current, next *Node
	current = s.Head
	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	s.Head = prev
}

func (s *Stack) Get(index int) (int, error) {
	temp := s.Head
	for i := 0; i < index; i++ {
		if temp == nil {
			return -50, nil
		}
		temp = temp.Next
	}
	return temp.Value, nil
}

func (s *Stack) Remove(val int) error {
	if s.Head == nil {
		return nil
	}
	if s.Head.Value == val {
		s.Head = s.Head.Next
		return nil
	}
	temp := s.Head
	for temp.Next != nil {
		if temp.Next.Value == val {
			temp.Next = temp.Next.Next
			return nil
		}
		temp = temp.Next
	}
	return nil
}

func (s *Stack) RemoveAll(val int) error {
	if s.Head == nil {
		return nil
	}
	for s.Head.Value == val {
		s.Head = s.Head.Next
	}
	temp := s.Head
	for temp.Next != nil {
		if temp.Next.Value == val {
			temp.Next = temp.Next.Next
		} else {
			temp = temp.Next
		}
	}
	return nil
}

func (s *Stack) RemoveAt(index int) error {
	if s.Head == nil {
		return nil
	}
	if index == 0 {
		s.Head = s.Head.Next
		return nil
	}
	temp := s.Head
	for i := 0; i < index-1; i++ {
		if temp == nil {
			return nil
		}
		temp = temp.Next
	}
	if temp.Next == nil {
		return nil
	}
	temp.Next = temp.Next.Next
	return nil
}

func (s *Stack) RemoveAllAt(index int) error {
	if s.Head == nil {
		return nil
	}
	if index == 0 {
		s.Head = s.Head.Next
		return nil
	}
	temp := s.Head
	for i := 0; i < index-1; i++ {
		if temp == nil {
			return nil
		}
		temp = temp.Next
	}
	if temp.Next == nil {
		return nil
	}
	temp.Next = temp.Next.Next
	return s.RemoveAllAt(index)
}

func (s *Stack) AddAt(index, val int) error {
	newNode := Node{Value: val, Next: nil}
	if index == 0 {
		newNode.Next = s.Head
		s.Head = &newNode
		return nil
	}
	temp := s.Head
	for i := 0; i < index-1; i++ {
		if temp == nil {
			return nil
		}
		temp = temp.Next
	}
	newNode.Next = temp.Next
	temp.Next = &newNode
	return nil
}

func (s *Stack) AddAllAt(index int, vals ...int) error {
	for _, val := range vals {
		newNode := Node{Value: val, Next: nil}
		if index == 0 {
			newNode.Next = s.Head
			s.Head = &newNode
			continue
		}
		temp := s.Head
		for i := 0; i < index-1; i++ {
			if temp == nil {
				return nil
			}
			temp = temp.Next
		}
		newNode.Next = temp.Next
		temp.Next = &newNode
	}
	return nil
}

func (s *Stack) AddAll(vals ...int) {
	for _, val := range vals {
		newNode := Node{Value: val, Next: nil}
		if s.Head != nil {
			newNode.Next = s.Head
		}
		s.Head = &newNode
	}
}

func (s *Stack) AddAllAtEnd(vals ...int) {
	temp := s.Head
	for temp.Next != nil {
		temp = temp.Next
	}
	for _, val := range vals {
		newNode := Node{Value: val, Next: nil}
		temp.Next = &newNode
		temp = temp.Next
	}
}

func (s *Stack) AddAllAtStart(vals ...int) {
	for _, val := range vals {
		newNode := Node{Value: val, Next: nil}
		if s.Head != nil {
			newNode.Next = s.Head
		}
		s.Head = &newNode
	}
}
