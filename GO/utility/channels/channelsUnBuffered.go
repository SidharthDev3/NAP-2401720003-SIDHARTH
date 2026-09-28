package channels

import "fmt"

func ChannelsUnBuffered() {
	ch := make(chan string)

	go func() {
		ch <- "Hello I am Sidharth"
	}()

	msg := <-ch

	fmt.Println(msg)
}
