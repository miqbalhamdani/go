package main

import (
	"fmt"
)

func main() {
	students := []string{
		"billy.tambunan93@gmail.com",
		"ibrahimker@gmail.com",
		"Iqbal Hamdani",
		"ldwimulya@hacktiv8.com",
		"steve.dewana@gmail.com",
		"teguh.afdilla@mnc-insurance.com",
		"Wicaksana Pratama",
		"Aulia Nurhady",
		"Muhamad Edi Sujarwo",
		"San Sayidul Akdam Augusta",
	}

	setStudentsPointer := func(items []string) []*string {
		pointerStudent := []*string{}

		for i := 0; i < len(items); i++ {
			pointerStudent = append(pointerStudent, &items[i])
		}

		return pointerStudent
	}

	getStudent := func (items []*string)  {
		for _, item := range items {
			fmt.Println(*item)
		}
	}

	studentsPointer := setStudentsPointer(students);

	getStudent(studentsPointer);
}
