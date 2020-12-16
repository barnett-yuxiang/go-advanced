package main

func main() {
	test01()

	added("a + b", func(a, b int) int {
		return a + b
	})
}
