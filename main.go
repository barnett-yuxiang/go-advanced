package main

import "github.com/barnett-yuxiang/go-advanced/test_json"

func main() {
	test01()

	added("a + b", func(a, b int) int {
		return a + b
	})

	test_json.Test()
}
