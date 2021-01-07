package advanced_usage

import "fmt"

func Added(msg string, a func(a, b int) int) {
	fmt.Println(msg, ":", a(33, 44))
}
