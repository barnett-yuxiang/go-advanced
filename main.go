package main

import (
	"github.com/barnett-yuxiang/go-advanced/advanced_usage"
	"github.com/barnett-yuxiang/go-advanced/test_json"
)

func main() {
	advanced_usage.Test01()

	advanced_usage.Added("a + b", func(a, b int) int {
		return a + b
	})

	test_json.Test()
}
