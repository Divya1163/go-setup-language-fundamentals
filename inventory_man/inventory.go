package main

import (
	"errors"
	"fmt"
)

// Item structure
type Item struct {
	ID       int
	Name     string
	Quantity int
	Price    float64
}

// Slice to store inventory items
var inventory []Item

// Add a new item
func AddItem(item Item) error {
	if item.ID <= 0 {
		return errors.New("item ID must be positive")
	}

	if item.Name == "" {
		return errors.New("item name cannot be empty")
	}

	if item.Quantity < 0 || item.Price < 0 {
		return errors.New("quantity and price cannot be negative")
	}

	// Check for duplicate ID
	for _, existing := range inventory {
		if existing.ID == item.ID {
			return errors.New("item ID already exists")
		}
	}

	inventory = append(inventory, item)
	return nil
}

// Update an existing item
func UpdateItem(id int, quantity int, price float64) error {
	if quantity < 0 || price < 0 {
		return errors.New("quantity and price cannot be negative")
	}

	for i := range inventory {
		if inventory[i].ID == id {
			inventory[i].Quantity = quantity
			inventory[i].Price = price
			return nil
		}
	}

	return errors.New("item not found")
}

// Remove an item
func RemoveItem(id int) error {
	for i, item := range inventory {
		if item.ID == id {
			inventory = append(inventory[:i], inventory[i+1:]...)
			return nil
		}
	}

	return errors.New("item not found")
}

// Display all items
func DisplayItems() {
	fmt.Println("\n--- Inventory Items ---")

	for _, item := range inventory {
		fmt.Printf("ID: %d | Name: %s | Quantity: %d | Price: %.2f\n",
			item.ID, item.Name, item.Quantity, item.Price)
	}
}

// Summary statistics structure
type Summary struct {
	ItemCount  int
	TotalValue float64
}

// Calculate summary concurrently using goroutine and channel
func CalculateSummary(items []Item) Summary {
	ch := make(chan Summary)

	go func() {
		var result Summary

		for _, item := range items {
			result.ItemCount++
			result.TotalValue += float64(item.Quantity) * item.Price
		}

		ch <- result
	}()

	summary := <-ch
	return summary
}
