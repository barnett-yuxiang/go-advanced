package multi_thread

import (
	"fmt"
	"time"
)

type Test02 struct {
}

func DoSomethingAfter1Second() {
	time.Sleep(1 * time.Second)
	fmt.Println("1. do some things")
}

func DoSomethingAfter2Second() {
	time.Sleep(2 * time.Second)
	fmt.Println("2. do some things")
}

func (x *Test02) Run() {
	go DoSomethingAfter1Second()
	go DoSomethingAfter2Second()
	select {
	default:
		fmt.Println("!.. Default case ..!")
	}
}
