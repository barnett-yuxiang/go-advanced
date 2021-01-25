package advanced_usage

import "fmt"

type ThreadStateGenerator struct {
}

func (tsg ThreadStateGenerator) Compute() error {
	fmt.Println("tsg Compute...")
	return nil
}
