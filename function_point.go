package main

import "fmt"

func added(msg string, a func(a, b int) int) {
	fmt.Println(msg, ":", a(33, 44))
}
