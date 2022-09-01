## Function

```go
func main() {
  greet("Iqbal", 24) // My name Iqbal, i`m 24 years old.
}

func greet(name string, age int8) {
  fmt.printf("My name %s, i`m %d years old.", name, age)
}
```

- Use `func` and follow by function name and parameter that used.
- This function is a function that doesn't return anything.
- We must init data type when the function accept parameter.

### With same data type

```go
func main() {
  greet("Iqbal", "24") // My name Iqbal, i`m 24 years old.
}

func greet(name, age string) {
  fmt.printf("My name %s, i`m %s years old.", name, age)
}
```

### With return

```go
func main() {
  greeting := greet("Iqbal", "24")
  fmt.printfgreeting() // My name Iqbal, i`m 24 years old.
}

func greet(name string, age int8) string {
  return fmt.Sprintf("My name %s, i`m %s years old.", name, age)
}
```

- Data type of return value must be define.
- `Sprintf` is the same with `Printf`, but `Sprintf` will return value, and `Printf` was not.

### With return multiple value

```go
func main() {
  diameter := 15

  area, circumference = calculate(diameter)

  fmt.Printfln("Area: ", area) // Area: 176.714.....
  fmt.Printfln("Circumference: ", circumference) // Circumference: 47.123.....
}

func calculate(name float) (float64, float64) {
  area := math.Pi * math.Pow(d/2, 3)

  circumference := math.Pi * d

  return area, circumference
}
```

### Predefined return value

```go
func calculate(name float) (area float64, circumference float64) {
  area := math.Pi * math.Pow(d/2, 3)

  circumference := math.Pi * d

  return
}
```

- The area and circumference variables that are not a new variable, but the variable is the variable that is used as the return value.
- If we use the predefined return value technique, then we need to reassign the variable used as the return value with a value that will produce a return value.
- But keep in mind that we still need the return keyword at the end of the function line

### Variadic function

Function where the function can accept an unlimited number of arguments. There is a dot three times before the description of the parameter data type it receives `(...students)`

```go
func main() {
  numbers := []int{1, 2, 3, 4, 5}
  result := sum(numbers)
}

func sum(n ...int) int {
  total := 0

  for _, v := range n {
    total += v
  }

  return total
}
```

Data type slice as parameter

```go
func main() {
  profile("Iqbal", "Apple", "Mango", "Orange")
}

func profile(name string, foods ...string) {
  ...
}
```

- We can combine regular parameters with variable parameters in a variadic function.
- But keep in mind here that the variadic parameter needs to be placed at the end of the parameter
