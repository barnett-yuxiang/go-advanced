package multi_thread

import (
	"fmt"
	"time"
)

type Test04 struct {
}

func (t *Test04) DoSomething(channel chan int) {
	time.Sleep(10 * time.Second)
	channel <- 1
}

func (t *Test04) Run() {
	// creating channel
	channel := make(chan int)

	go t.DoSomething(channel)

	// channel is not ready
	// and no default case
	for {
		time.Sleep(1 * time.Second)
		select {
		case op1 := <-channel:
			fmt.Println(op1)
			//return
		default:
			fmt.Println("Not Found")
		}
	}
}
