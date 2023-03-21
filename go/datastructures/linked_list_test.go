package datastructures

import "testing"

// tests if outcome is "Inserted" > "Original"
func TestLinkedListInsertHead(t *testing.T) {
	n := NewLinkedList("Original")
	n.InsertHead("Inserted")

	// test length
	if n.Size() != 2 {
		t.Errorf("got %d; want %d", n.Size(), 2)
	}

	// test head is "Inserted"
	if n.head.data.(string) != "Inserted" {
		t.Errorf("got %s; want %s", n.head.data.(string), "Inserted")
	}

	// test tail is "Original"
	if n.tail.data.(string) != "Original" {
		t.Errorf("got %s; want %s", n.tail.data.(string), "Original")
	}

	// test head must not have a previous
	if n.head.previous != nil {
		t.Errorf("got %v; want nil", n.head.previous)
	}

	// test tail must not have a next
	if n.tail.next != nil {
		t.Errorf("got %v; want nil", n.tail.next)
	}

	// test head must have "Original" as next
	if n.head.next.data.(string) != "Original" {
		t.Errorf("got %v; want %s", n.head.next.data.(string), "Original")
	}

	// test tail must have "Inserted" as previous
	if n.tail.previous.data.(string) != "Inserted" {
		t.Errorf("got %v; want %s", n.tail.next.data.(string), "Inserted")
	}
}

// tests if outcome is "Original" > "Inserted"
func TestLinkedListInsertTail(t *testing.T) {
	n := NewLinkedList("Original")
	n.InsertTail("Inserted")

	// test length
	if n.Size() != 2 {
		t.Errorf("got %d; want %d", n.Size(), 2)
	}

	// test head is "Original"
	if n.head.data.(string) != "Original" {
		t.Errorf("got %s; want %s", n.head.data.(string), "Original")
	}

	// test tail is "Inserted"
	if n.tail.data.(string) != "Inserted" {
		t.Errorf("got %s; want %s", n.tail.data.(string), "Inserted")
	}

	// test head must not have a previous
	if n.head.previous != nil {
		t.Errorf("got %v; want nil", n.head.previous)
	}

	// test tail must not have a next
	if n.tail.next != nil {
		t.Errorf("got %v; want nil", n.tail.next)
	}

	// test head must have "Inserted" as next
	if n.head.next.data.(string) != "Inserted" {
		t.Errorf("got %v; want %s", n.head.next.data.(string), "Inserted")
	}

	// test tail must have "Original" as previous
	if n.tail.previous.data.(string) != "Original" {
		t.Errorf("got %v; want %s", n.head.next.data.(string), "Original")
	}
}

// from original "First" > "Second" > "Third"
// tests if outcome is "Second" > "Third"
func TestLinkedListDeleteHeadWithThreeItems(t *testing.T) {
	n := NewLinkedList("Second")
	n.InsertHead("First")
	n.InsertTail("Third")

	err := n.DeleteHead()
	if err != nil {
		t.Errorf("got %s; want nil", err.Error())
	}

	// test length
	if n.Size() != 2 {
		t.Errorf("got %d; want %d", n.Size(), 2)
	}

	// test head is "Second"
	if n.head.data.(string) != "Second" {
		t.Errorf("got %s; want %s", n.head.data.(string), "Second")
	}

	// test tail is "Third"
	if n.tail.data.(string) != "Third" {
		t.Errorf("got %s; want %s", n.tail.data.(string), "Third")
	}

	// test head must not have a previous
	if n.head.previous != nil {
		t.Errorf("got %v; want nil", n.head.previous)
	}

	// test tail must not have a next
	if n.tail.next != nil {
		t.Errorf("got %v; want nil", n.tail.next)
	}

	// test head must have "Third" as next
	if n.head.next.data.(string) != "Third" {
		t.Errorf("got %v; want %s", n.head.next.data.(string), "Third")
	}

	// test tail must have "Second" as previous
	if n.tail.previous.data.(string) != "Second" {
		t.Errorf("got %v; want %s", n.tail.next.data.(string), "Second")
	}
}

// from original "First" > "Second" > "Third"
// tests if outcome is "First" > "Second"
func TestLinkedListDeleteTailWithThreeItems(t *testing.T) {
	n := NewLinkedList("Second")
	n.InsertHead("First")
	n.InsertTail("Third")

	err := n.DeleteTail()
	if err != nil {
		t.Errorf("got %s; want nil", err.Error())
	}

	// test length
	if n.Size() != 2 {
		t.Errorf("got %d; want %d", n.Size(), 2)
	}

	// test head is "First"
	if n.head.data.(string) != "First" {
		t.Errorf("got %s; want %s", n.head.data.(string), "First")
	}

	// test tail is "Second"
	if n.tail.data.(string) != "Second" {
		t.Errorf("got %s; want %s", n.tail.data.(string), "Second")
	}

	// test head must not have a previous
	if n.head.previous != nil {
		t.Errorf("got %v; want nil", n.head.previous)
	}

	// test tail must not have a next
	if n.tail.next != nil {
		t.Errorf("got %v; want nil", n.tail.next)
	}

	// test head must have "Second" as next
	if n.head.next.data.(string) != "Second" {
		t.Errorf("got %v; want %s", n.head.next.data.(string), "Second")
	}

	// test tail must have "First" as previous
	if n.tail.previous.data.(string) != "First" {
		t.Errorf("got %v; want %s", n.tail.next.data.(string), "First")
	}
}

func TestDisplayLinkedList(t *testing.T) {
	n := NewLinkedList("Second")
	n.InsertHead("First")
	n.InsertTail("Third")

	n.Display()
}
