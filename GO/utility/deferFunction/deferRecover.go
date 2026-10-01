package deferFunction

import "fmt"

func safeRun() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from Panic: ", r)
		}
	}()

	fmt.Println("Starting...")
	panic("Something went wrong!!")
	// fmt.Println("Program executed succesfully")
}

func DeferRecover() {
	safeRun()
}
