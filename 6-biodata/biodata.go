package main

import (
	"fmt"
	"os"
	"strconv"
)

type Student struct {
	name string
	score int
	job string
}

func printStudent(num int) {
	students := []*Student{
		{name: "Wicaksana Pratama", score: 30, job: "FE" },
		{name: "Aulia Nurhady", score: 50, job: "BE" },
		{name: "Muhamad Edi Sujarwo", score: 80, job: "BE" },
		{name: "San Sayidul Akdam Augusta", score: 70, job: "BE" },
	}

	fmt.Println(students[num])
}

func main() {
	arg, arg2 := strconv.Atoi(os.Args[1])

	fmt.Println(arg, arg2)

	printStudent(arg)
}


