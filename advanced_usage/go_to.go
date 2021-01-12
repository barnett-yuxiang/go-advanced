package advanced_usage

import (
	"errors"
	"fmt"
)

var condition1 = false

var condition2 = true

func returnError() (int, error) {
	return 222, errors.New("333")
}

func TestGoTo() error {

	var ret error

	if condition1 {
		ret = errors.New("111")
		goto Done
	}

	if condition2 {
		num, err := returnError()
		ret = err
		fmt.Println(num)
		goto Done
	}

Done:
	return ret
}

func Run() {
	ret := TestGoTo()
	fmt.Println(ret)
}
