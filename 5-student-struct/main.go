package main

import (
	"fmt"
)

type Student struct {
	name string
	score int
}

func main() {
	students := []*Student{
		{name: "billy.tambunan93@gmail.com", score: 10 },
		{name: "ibrahimker@gmail.com", score: 20 },
		{name: "Iqbal Hamdani", score: 10 },
		{name: "ldwimulya@hacktiv8.com", score: 30 },
		{name: "steve.dewana@gmail.com", score: 50 },
		{name: "teguh.afdilla@mnc-insurance.com", score: 90 },
		{name: "Wicaksana Pratama", score: 30 },
		{name: "Aulia Nurhady", score: 50 },
		{name: "Muhamad Edi Sujarwo", score: 80 },
		{name: "San Sayidul Akdam Augusta", score: 70 },
	}

	getStudent := func (items []*Student)  {
		for _, item := range items {
			fmt.Println(item.name)
		}
	}

	getStudent(students);
}
