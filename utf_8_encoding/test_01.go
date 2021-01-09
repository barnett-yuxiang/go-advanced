package utf_8_encoding

import "fmt"

type Test01 struct {
}

func (Test01) Run() {
	byteVal := byte('a')
	runeVal := 'a' //rune as default type of char is rune
	fmt.Printf("%c = %d %c = %U", byteVal, byteVal, runeVal, runeVal)
}
