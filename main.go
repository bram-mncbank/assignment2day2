package main

import "fmt"

type Person struct {
	name string
	usia int
}

func main() {
	var listpersons = []*Person{
		{name: "bram", usia: 42},
		{name: "satya", usia: 22},
		{name: "adi", usia: 32},
		{name: "frizky", usia: 40},
		{name: "sigit", usia: 20},
	}

	var showlist = func(list []*Person) {
		for _, v := range list {
			fmt.Printf("Nama: %s\n", v.name)
			fmt.Printf("Usia: %d\n", v.usia)
			fmt.Println("----------------")
		}
	}

	showlist(listpersons)
}
