package lab3

import (
	"bufio"
	"fmt"
	"os"
)

type Person struct {
	Name   string
	Age    int
	Job    string
	Salary float64
}

func (p Person) readData() Person {
	var reader = bufio.NewReader(os.Stdin)

	fmt.Print("Enter Name of the person: ")
	fmt.Scan(&p.Name)
	reader.ReadLine()

	fmt.Print("Enter Age of the person: ")
	fmt.Scan(&p.Age)
	reader.ReadLine()

	fmt.Print("Enter Job of the person: ")
	fmt.Scan(&p.Job)
	reader.ReadLine()

	fmt.Print("Enter Salary of the person: ")
	fmt.Scan(&p.Salary)
	reader.ReadLine()

	return p
}

func (p Person) printData() {
	fmt.Printf("Name: %v \n", p.Name)
	fmt.Printf("Age: %v \n", p.Age)
	fmt.Printf("Job: %v \n", p.Job)
	fmt.Printf("Salary: %v \n", p.Salary)
}

func main() {
	var p1, p2 Person

	p1 = p1.readData()
	p1.printData()

	p2 = p2.readData()
	p2.printData()
}
