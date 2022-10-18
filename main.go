package main

import (
	"fmt"
)

func main() {
  var evenNumbers = func(n ...int) []int {
    var result []int

    for _, v := range n {
      if v % 2 == 0 {
        result = append(result, v)
      }
    }

    return result
  }

  var numbers = []int{4, 93, 77, 10, 52, 22, 34}

  fmt.Println(evenNumbers(numbers...)) // [4, 52, 22, 34]
}
