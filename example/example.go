package main

import (
	"fmt"

	"github.com/breml/sample"
)

func main() {
	s, err := sample.NewLessThan(10)
	if err != nil {
		fmt.Println("Unable to initialize sampler", err)
	}
	for i := 0; i < 100; i++ {
		if s.Sample() {
			fmt.Println(i, "got sampled by LessThan sampler", s)
		}
	}
}
