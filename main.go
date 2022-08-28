package main

import "fmt"

type Person struct {
	name string
	usia int
}

func main() {
	var listpersons = []*Person{
		{name: "bram", usia: 42},
		{name: "nugraha", usia: 22},
		{name: "adi", usia: 32},
		{name: "frizky", usia: 40},
		{name: "sigit", usia: 20},
		{name: "agus", usia: 25},
		{name: "satrio", usia: 21},
		{name: "bayu", usia: 24},
		{name: "putu", usia: 24},
		{name: "eka", usia: 24},
	}

	var showlist = func(list []*Person) {
		for _, v := range list {
			fmt.Printf("Nama: %s\n", v.name)
			fmt.Printf("Usia: %d\n", v.usia)
			fmt.Println("----------------")
		}
	}

	fmt.Println("===============All Datas=================")
	showlist(listpersons)
	fmt.Println()
	fmt.Println("==================Slicing================")
	var s = listpersons[3:4]
	showlist(s)

}
