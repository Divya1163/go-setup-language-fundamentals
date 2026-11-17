package main

import "testing"

// Reset inventory before each test
func resetInventory() {
	inventory = []Item{}
}

// Test adding an item
func TestAddItem(t *testing.T) {
	resetInventory()

	item := Item{1, "Laptop", 5, 50000}

	err := AddItem(item)

	if err != nil {
		t.Errorf("Expected item to be added, got error: %v", err)
	}

	if len(inventory) != 1 {
		t.Errorf("Expected 1 item, got %d", len(inventory))
	}
}

// Test duplicate ID validation
func TestDuplicateItem(t *testing.T) {
	resetInventory()

	AddItem(Item{1, "Laptop", 5, 50000})

	err := AddItem(Item{1, "Mouse", 10, 500})

	if err == nil {
		t.Error("Expected error for duplicate ID")
	}
}

// Test updating an item
func TestUpdateItem(t *testing.T) {
	resetInventory()

	AddItem(Item{1, "Laptop", 5, 50000})

	err := UpdateItem(1, 10, 55000)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if inventory[0].Quantity != 10 {
		t.Errorf("Expected quantity 10, got %d", inventory[0].Quantity)
	}
}

// Test removing an item
func TestRemoveItem(t *testing.T) {
	resetInventory()

	AddItem(Item{1, "Laptop", 5, 50000})

	err := RemoveItem(1)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(inventory) != 0 {
		t.Errorf("Expected empty inventory, got %d items", len(inventory))
	}
}
