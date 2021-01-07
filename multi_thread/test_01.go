package multi_thread

import (
	"fmt"
	"sync"
	"time"
)

type Test01 struct {
}

func execute(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Start")
	cnt := 0
	for {
		time.Sleep(1 * time.Second)
		fmt.Printf("execute %d\n", cnt)
		cnt++
		if cnt > 5 {
			break
		}
	}
	fmt.Println("End")
}

func (t01 *Test01) Run() {
	var wg sync.WaitGroup
	wg.Add(1)
	go execute(&wg)
	wg.Wait()
}
