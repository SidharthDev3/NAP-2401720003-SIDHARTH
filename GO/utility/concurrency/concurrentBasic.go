package concurrency

import (
	"fmt"
	"time"
)

func print(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(i, msg)
		time.Sleep(time.Millisecond * 500)

	}
}

func ConcurrentBasic() {
	go print("Sidharth") //Concurrent
	go print("Hello")    //Concurrent
	print("Hi")          //main
}
