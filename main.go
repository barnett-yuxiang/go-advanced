package main

import (
	"github.com/barnett-yuxiang/go-advanced/utf_8_encoding"
)

func main() {
	// @
	//advanced_usage.Test01()
	//advanced_usage.Added("a + b", func(a, b int) int {
	//	return a + b
	//})
	//test_json.Test()

	// @
	//t := multi_thread.Test01{}
	//t.Run()

	// @
	//t := multi_thread.Test02{}
	//t.Run()

	// @
	//t := multi_thread.Test04{}
	//t.Run()

	//
	t := utf_8_encoding.Test01{}
	t.Run()
}
