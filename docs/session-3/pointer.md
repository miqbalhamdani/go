## Pointer

Use for saving variable memory address. For example, we have a variable with type integer and value is 3, the pointer is an address where the memory of 3 is, not the value 3 itself.

We can use asterisk * before data type to define pointer.

```go
func main {
  var firstNumber int  = 3
  var secondNumber *int = &firstNumber

  fmt.Printfln("firstNumber (value) :", firstNumber) // 3
  fmt.Printfln("firstNumber (memory) :", &firstNumber) // 0xc0000ad007

  fmt.Printfln("secondNumber (value) :", *secondNumber) // 3
  fmt.Printfln("secondNumber (memory) :", secondNumber) // 0xc0000ad007
}
```

- firstNumber is value.
- secondNumber is pointer that containing the memory address of firstNumber.
- to assign a value to the pointer variable, we can use ampersand sign `&`.
- to display the memory address of the firstNumber, we can use ampersand sign `&` too.
- to display real value in variable pointer, we can use asterisk `*`.

<br>

```go
func main {
  var firstNumber int  = 3
  var secondNumber *int = &firstNumber

  fmt.Printfln("firstNumber (value) :", firstNumber) // 3
  fmt.Printfln("firstNumber (memory) :", &firstNumber) // 0xc0000ad007

  *secondNumber = 8

  fmt.Printfln("secondNumber (value) :", firstNumber) // 8
  fmt.Printfln("secondNumber (memory) :", *secondNumber) // 8
}
```

Pointers are used to store memory addresses, so when we change the value of a pointer, other variables that have the same memory address will also be changed in value.

### Pointer as a parameter

```go
func main {
  a int := 5
  fmt.Println("Before:", a) // 10

  changeValue(&a)

  fmt.Println("After:", a) // 20
}

func changeValue(number *int) {
  *number = 20
}
```
