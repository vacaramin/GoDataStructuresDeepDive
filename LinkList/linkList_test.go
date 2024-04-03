package linklist

import (
	"testing"
)

func TestAdd(t *testing.T) {
	ll := &LinkList{}
	ll.Add(1)
	if ll.head.val != 1 {
		t.Errorf("Expected head value of 1, got %d", ll.head.val)
	}

	ll.Add(2)
	if ll.head.Next.val != 2 {
		t.Errorf("Expected second node value of 2, got %d", ll.head.Next.val)
	}
}

func TestRemove(t *testing.T) {
	ll := &LinkList{}
	ll.Add(1)
	ll.Add(2)
	ll.Remove(1)
	if ll.head.val != 2 {
		t.Errorf("Expected head value of 2 after removal, got %d", ll.head.val)
	}

	err := ll.Remove(3)
	if err == nil {
		t.Error("Expected an error when trying to remove a non-existent element, got nil")
	}
}

func TestEmptyRemove(t *testing.T) {
	ll := &LinkList{}
	err := ll.Remove(1)
	if err == nil {
		t.Error("Expected an error when trying to remove from an empty list, got nil")
	}
}
