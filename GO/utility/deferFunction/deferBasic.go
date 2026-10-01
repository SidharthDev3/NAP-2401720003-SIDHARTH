package deferFunction

import "fmt"

func work() {
	defer fmt.Println("Goodbye!!")
	fmt.Println("Helo!!")
	fmt.Println("Working...")
}

func DeferBasic() {
	work()
}
