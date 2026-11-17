package main

import (
	"fmt"
	"time"
)

type Inventory struct {
	Id       int
	Name     string
	Quantity int
	Price    float64
}

func itemCount(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		time.Sleep(time.Millisecond * 500)
		results <- job + 0
	}
}

func main() {
	obj := []Inventory{
		{2, "Macbook Pro", 5, 300000},
		{3, "Airpods", 5, 20000},
		{4, "Airtag", 5, 5000},
		{5, "Sunglasses", 5, 20000},
	}

	a := Inventory{Id: 1, Name: "iPhone 17", Quantity: 5, Price: 100000}
	obj = append(obj, a)

	b := Inventory{Id: 2, Name: "Macbook Pro", Quantity: 4, Price: 300000}
	obj = append(obj, b)

	fmt.Println(obj)

	n := len(obj)

	item := make(chan int, n)

	result := make(chan int, n)

	for w := 1; w <= n; w++ {
		go itemCount(w, item, result)
	}

	for j := 1; j <= n; j++ {
		item <- j
	}

	close(item)

	for a := 1; a <= n; a++ {
		fmt.Println(<-result)
	}

}
