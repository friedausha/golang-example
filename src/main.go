package main

import (
	"fmt"
	"golang-example/nomor_dua"
	"golang-example/nomor_empat"
	"golang-example/nomor_satu"
)

func main() {
	//res, _ := addition.Add(11, 15)
	//fmt.Println(res)

	// Buat slice of int bernama "a" dengan panjang 5
	//a := make([]int, 5, 5)
	//printSlice("a", a)
	//
	//// Buat slice of int bernama "b" dengan panjang 0 kapasitas 5
	//b := make([]int, 0, 5)
	//printSlice("b", b)
	//
	//// Buat variabel c dengan dua data awal b
	//c := b[:2]
	//printSlice("c", c)
	//
	//// Buat variabel d dengan data ke 2 sampai 4
	//d := b[2:5]
	//printSlice("d", d)
	//
	//m := map[string]int{"answer": 42}
	//v, ok := m["answer1"]
	//// gunakan pengecekan: elem, ok = m[key]
	//fmt.Println("The value:", v, "Present?", ok)

	nomor_satu.PrintArrays([]int{2, 3, 5, 7, 9})
	nomor_dua.PrintThirdAndFifthArray([]int{2, 4, 6, 8, 10})
	nomor_empat.PrintSquareRoot([]int{-100, 32, 20, 1, 2, 4, 8, 16})
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
