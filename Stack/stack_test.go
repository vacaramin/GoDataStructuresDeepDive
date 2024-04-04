package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	// Create a new stack
	s := NewStack()

	// Test IsEmpty on an empty stack
	if !s.IsEmpty() {
		t.Error("Expected stack to be empty")
	}

	// Test Push and Peek
	s.Push(1)
	s.Push(2)
	s.Push(3)
	if val, _ := s.Peek(); val != 3 {
		t.Errorf("Expected Peek to return 3, got %d", val)
	}

	// Test Size
	if size := s.Size(); size != 3 {
		t.Errorf("Expected Size to return 3, got %d", size)
	}

	// Test Pop
	if val, _ := s.Pop(); val != 3 {
		t.Errorf("Expected Pop to return 3, got %d", val)
	}
	if size := s.Size(); size != 2 {
		t.Errorf("Expected Size to return 2 after Pop, got %d", size)
	}

	// Test Pop on an empty stack
	s.Clear()
	if _, err := s.Pop(); err == nil {
		t.Error("Expected Pop to return an error on an empty stack")
	}

	// Test Peek on an empty stack
	if _, err := s.Peek(); err == nil {
		t.Error("Expected Peek to return an error on an empty stack")
	}
}

func TestStack2(t *testing.T) {
	// Create a new stack
	s := NewStack()

	// Test IsEmpty on an empty stack
	if !s.IsEmpty() {
		t.Error("Expected stack to be empty")
	}

	// Test Push and Peek
	s.Push(1)
	s.Push(2)
	s.Push(3)
	if val, _ := s.Peek(); val != 3 {
		t.Errorf("Expected Peek to return 3, got %d", val)
	}

	// Test Size
	if size := s.Size(); size != 3 {
		t.Errorf("Expected Size to return 3, got %d", size)
	}

	// Test Pop
	if val, _ := s.Pop(); val != 3 {
		t.Errorf("Expected Pop to return 3, got %d", val)
	}
	if size := s.Size(); size != 2 {
		t.Errorf("Expected Size to return 2 after Pop, got %d", size)
	}

	// Test Pop on an empty stack
	s.Clear()
	if _, err := s.Pop(); err == nil {
		t.Error("Expected Pop to return an error on an empty stack")
	}

	// Test Peek on an empty stack
	if _, err := s.Peek(); err == nil {
		t.Error("Expected Peek to return an error on an empty stack")
	}
}
