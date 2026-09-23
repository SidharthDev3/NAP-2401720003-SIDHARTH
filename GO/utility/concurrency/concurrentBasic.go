package concurrency

import (
	"fmt"
	"time"
)

// func print(msg string) {
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(i, msg)
// 		time.Sleep(time.Millisecond * 500)

// 	}
// }

func print(msg string) {
	fmt.Println(msg)

}

func ConcurrentBasic() {
	go print("Sidharth")
	// time.Sleep(time.Millisecond * 400)
	go print("Tom")
	time.Sleep(time.Millisecond * 500)

}
