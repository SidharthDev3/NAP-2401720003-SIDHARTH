package main

import "fmt"

type Inventory struct {
	Id       int
	Name     string
	Quantity int
	Price    float64
}

func addItem(obj []Inventory, a Inventory) []Inventory {

	for _, v := range obj {
		if v.Id == a.Id {
			fmt.Println("Item already exists")
			return obj
		}
	}

	if a.Id <= 0 || a.Name == "" || a.Quantity < 0 || a.Price < 0 {
		fmt.Println("Invalid item")
		return obj
	}

	obj = append(obj, a)
	return obj
}

func updateItem(obj []Inventory, id int, quantity int) {

	for i := 0; i < len(obj); i++ {
		if obj[i].Id == id {
			obj[i].Quantity = quantity
			fmt.Println("Item updated")
			return
		}
	}

	fmt.Println("Item not found")
}

func removeItem(obj []Inventory, id int) []Inventory {

	for i := 0; i < len(obj); i++ {
		if obj[i].Id == id {
			obj = append(obj[:i], obj[i+1:]...)
			fmt.Println("Item removed")
			return obj
		}
	}

	fmt.Println("Item not found")
	return obj
}

func itemCount(obj []Inventory, result chan float64) {

	total := 0.0

	for i := 0; i < len(obj); i++ {
		total = total + float64(obj[i].Quantity)*obj[i].Price
	}

	result <- total
}

func main() {

	obj := []Inventory{
		{2, "Macbook Pro", 5, 300000},
		{3, "Airpods", 5, 20000},
		{4, "Airtag", 5, 5000},
		{5, "Sunglasses", 5, 20000},
	}

	a := Inventory{1, "iPhone 17", 5, 100000}
	obj = addItem(obj, a)

	updateItem(obj, 2, 4)

	obj = removeItem(obj, 4)

	fmt.Println("\nINVENTORY DETAILS")

	for i := 0; i < len(obj); i++ {
		fmt.Println(obj[i])
	}

	n := len(obj)
	result := make(chan float64)

	go itemCount(obj, result)

	total := <-result

	fmt.Println("\nINVENTORY SUMMARY")
	fmt.Println("Total Items:", n)
	fmt.Println("Total Value:", total)

}
