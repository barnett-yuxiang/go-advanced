package main

import "github.com/barnett-yuxiang/go-advanced/multi_thread"

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
	t := multi_thread.Test03{}
	t.Run()
}
