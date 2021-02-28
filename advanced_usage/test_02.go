package advanced_usage

import "fmt"

type Test03 struct {
}

type Sample struct {
	Symbols []string
}

func (t *Test03) Foo() {
	symbols := []string{"main", "hello", "world"}
	fmt.Println(symbols)
	symbols = symbols[1:2:2]
	fmt.Println(symbols)

	symbols = symbols[:0]
	fmt.Println(symbols)
	fmt.Println(symbols)
	fmt.Println(symbols)

	var batchSamples []Sample
	batchSamples = append(
		batchSamples,
		Sample{Symbols: []string{"main", "100", "200"}},
		Sample{Symbols: []string{"main"}},
		Sample{Symbols: []string{"main", "10000"}},
	)

	for _, sample := range batchSamples {
		fmt.Printf("1.%p, %s\n", &sample, sample)
		if sample.Symbols[0] == "main" {
			sample.Symbols[0] = "new_main"
		}

		if len(sample.Symbols) > 1 {
			sample.Symbols = sample.Symbols[1:]
		}
	}

	fmt.Println("======")

	for _, sample := range batchSamples {
		fmt.Printf("2.%p, %s\n", &sample, sample)
	}

	fmt.Println("======")

	for _, sample := range batchSamples {
		if len(sample.Symbols) == 1 {
			sample.Symbols = sample.Symbols[1:]
			//batchSamples[ids] = sample
			fmt.Println(sample)
		} else {
			sample.Symbols = sample.Symbols[1:]
			fmt.Println(sample)
		}
	}

	fmt.Println("======")
	for index, sample := range batchSamples {
		fmt.Println(index, sample)
	}
	fmt.Println("======")
	for index, sample := range batchSamples {
		fmt.Println(index, sample)
	}
}

func (t *Test03) Boo() {
	v := []int{1, 2, 3}
	for i := range v {
		v = append(v, i)
	}
	fmt.Println(v)
}
