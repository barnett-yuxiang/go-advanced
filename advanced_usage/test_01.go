package advanced_usage

import "fmt"

type Test02 struct {
}

func (t Test02) Foo() {
	bundle := make(map[string]string)
	t.Bar(bundle)
	bundle["service"] = "startup"
	fmt.Println(bundle)
}

func (t Test02) Bar(data map[string]string) {
	data["tag"] = "lark"
}

func (t Test02) dynamic() {
	var generator TableGenerator
	generator = ThreadStateGenerator{}
	_ = generator.Compute()
}
