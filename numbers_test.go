package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	numbers := numbers{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	holder := numbers.evenOrOdd()
	if len(holder) != 10 {
		t.Errorf("Expected length 10, got %d", len(holder))
	}
	if holder[0] != "odd" {
		t.Errorf("Expected 'odd', got %s", holder[0])
	}
	if holder[1] != "even" {
		t.Errorf("Expected 'even', got %s", holder[1])
	}

}
