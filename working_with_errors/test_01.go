package working_with_errors

import (
	"errors"
	"fmt"
)

type Test01 struct {
}

type NotFoundError struct {
	Name string
}

func (e *NotFoundError) Error() string {
	return e.Name + ": not found\n"
}

func (t *Test01) Foo() {
	err := fmt.Errorf("Parser or Transformer: not found\n")
	fmt.Println(err)

	if e, ok := err.(*NotFoundError); ok {
		fmt.Println(e)
	}

	nfe := NotFoundError{
		Name: "Parser or Transformer",
	}

	if errors.Is(err, &nfe) {
		fmt.Println("here")
	}
}
