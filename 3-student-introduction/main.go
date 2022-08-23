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

	for _, student := range students {
		fmt.Println(student)
	}
}
