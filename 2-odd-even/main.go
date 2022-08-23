package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			fmt.Printf("%d adalah bilangan Genap \n", i)
		} else {
			fmt.Printf("%d adalah bilangan Ganjil \n", i)
		}
	}
}
