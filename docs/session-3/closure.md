## Closure

Closure is an anonymous function

```go
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
```

### Closure (IIFE)

IIFE (immediately-invoked function expression) is closure that can be execute immediately at the first declaration.

```go
func main() {
  var numbers = []int{4, 93, 77, 10, 52, 22, 34}

  var evenNumbers = func(n ...int) []int {
    var result []int

    for _, v := range n {
      if v % 2 == 0 {
        result = append(result, v)
      }
    }

    return result
  }(numbers...)

  fmt.Println(evenNumbers) // [4, 52, 22, 34]
}
```

- To call **IIFE** we dont need brackets ()

### Closure as a return value

```go
func main() {
  studentList := []string{"Airell", "Nanda", "Mailo", "Marco"}

  find := findStudent(studentLists)

  fmt.Println(find("airell"))
}

func findStudent(students []string) func(string) string {

  return func(s string) string {
    return ...
  }
}
```

### Closure Callback

```go
func main() {
  numbers := []int{4, 93, 77, 10, 52, 22, 34}

  find := findOddNumbers(numbers, func(number int) bool {
    return number%2 != 0
  })
}

func findOddNumbers(numbers []int, callback func(int) bool) int {
  return ...
}
```

Closure Callback with type alias

```go
type isOddNum func(int) bool

func main() {
  numbers := []int{4, 93, 77, 10, 52, 22, 34}

  find := findOddNumbers(numbers, func(number int) bool {
    return number%2 != 0
  })
}

func findOddNumbers(numbers []int, callback isOddNum) int {
  return ...
}
```
