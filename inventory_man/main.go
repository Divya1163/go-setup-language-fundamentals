package main

import "fmt"

func main() {

	// Adding items
	fmt.Println("Adding Items:")

	err := AddItem(Item{1, "Laptop", 5, 50000})
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Laptop added successfully")
	}

	err = AddItem(Item{2, "Mouse", 10, 500})
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Mouse added successfully")
	}

	err = AddItem(Item{3, "Keyboard", 8, 1200})
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Keyboard added successfully")
	}

	// Display inventory
	DisplayItems()

	// Updating an item
	fmt.Println("\nUpdating Laptop:")
	err = UpdateItem(1, 7, 52000)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Laptop updated successfully")
	}

	// Removing an item
	fmt.Println("\nRemoving Mouse:")
	err = RemoveItem(2)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Mouse removed successfully")
	}

	// Display updated inventory
	DisplayItems()

	// Concurrent summary calculation
	fmt.Println("\n--- Inventory Summary ---")

	summary := CalculateSummary(inventory)

	fmt.Println("Total Items:", summary.ItemCount)
	fmt.Printf("Total Inventory Value: Rs. %.2f\n", summary.TotalValue)
}
