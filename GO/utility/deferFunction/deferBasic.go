package deferFunction

import "fmt"

func work() {
	defer fmt.Println("Goodbye!!")
	defer fmt.Println("Bye!!")
	defer fmt.Println("Bye Bye!!")
	fmt.Println("Helo!!")
	fmt.Println("Working...")
}

func DeferBasic() {
	work()
}
