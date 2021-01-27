package race_condition

import (
	"fmt"
	"sync"
)

type Test01 struct {
}

var Wait sync.WaitGroup
var Counter int = 0

func (t *Test01) Run() {
	for routine := 1; routine <= 2; routine++ {

		Wait.Add(1)
		go Routine(routine)
	}

	Wait.Wait()
	fmt.Printf("Final Counter: %d\n", Counter)
}

func Routine(id int) {

	for count := 0; count < 2; count++ {

		value := Counter
		value++
		Counter = value
	}

	Wait.Done()
}
