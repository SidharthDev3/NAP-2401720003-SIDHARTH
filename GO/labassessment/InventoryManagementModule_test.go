package main

import "testing"

func TestAddItem(t *testing.T) {

	obj := []Inventory{}

	a := Inventory{1, "Laptop", 5, 50000}

	obj = addItem(obj, a)

	if len(obj) != 1 {
		t.Error("Item not added")
	}
}

func TestUpdateItem(t *testing.T) {

	obj := []Inventory{
		{1, "Laptop", 5, 50000},
	}

	updateItem(obj, 1, 10)

	if obj[0].Quantity != 10 {
		t.Error("Item not updated")
	}
}

func TestRemoveItem(t *testing.T) {

	obj := []Inventory{
		{1, "Laptop", 5, 50000},
		{2, "Mouse", 3, 1000},
	}

	obj = removeItem(obj, 1)

	if len(obj) != 1 {
		t.Error("Item not removed")
	}
}

func TestTotalValue(t *testing.T) {

	obj := []Inventory{
		{1, "Laptop", 2, 50000},
		{2, "Mouse", 3, 1000},
	}

	result := make(chan float64)

	go itemCount(obj, result)

	total := <-result

	if total != 103000 {
		t.Error("Total value is incorrect")
	}
}
