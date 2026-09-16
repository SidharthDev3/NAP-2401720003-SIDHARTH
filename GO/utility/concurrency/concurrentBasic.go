package concurrency

import (
	"fmt"
	"time"
)

func count(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(i, msg)
		time.Sleep(time.Millisecond * 500)

	}
}

// var wg sync.WaitGroup
// wg.Add(1)
// go func(){
// 	defer wg.Done()
// 	work()
// }()
// wg.wait()

func ConcurrentBasic() {
	go count("Sidharth") //Concurrent
	go count("Hello")    //Concurrent
	count("Hi")          //main
}
