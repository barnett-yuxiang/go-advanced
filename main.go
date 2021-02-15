package main

import "github.com/barnett-yuxiang/go-advanced/cool_shell"

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
	//t := utf_8_encoding.Test01{}
	//t.Run()

	//t := working_with_errors.Test01{}
	//t.Foo()

	//t := advanced_usage.Test02{}
	//t.Foo()

	//t := file_lock.Test01{}
	//t.Run()

	///////////////////////
	//c := make(chan bool)
	//m := make(map[string]string)
	//go func() {
	//	m["1"] = "a" // First conflicting access.
	//	c <- true
	//}()
	//m["2"] = "b" // Second conflicting access.
	//<-c
	//for k, v := range m {
	//	fmt.Println(k, v)
	//}
	///////////////////////

	//
	//t := race_condition.Test01{}
	//t.Run()

	cool_shell.Test01()
	cool_shell.Test01Fixed()
	cool_shell.TestDeepEqual()
}
