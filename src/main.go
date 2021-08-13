package main

import (
	"fmt"
	"time"
	"sync"
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

	//nomor_satu.PrintArrays([]int{2, 3, 5, 7, 9})
	//nomor_dua.PrintThirdAndFifthArray([]int{2, 4, 6, 8, 10})
	//nomor_empat.PrintSquareRoot([]int{32, 20, 1, 2, 4, 8, 16})

	//go count("sheep")
	//go count("fish")

	//go count("sheep")
	//go count("fish")
	//fmt.Scanln()

	var wg sync.WaitGroup
	//
	//wg.Add(1)
	//wg.Add(3)
	////c := make(chan string)
	//go func(f string) {
	//	count(f)
	//	wg.Done()
	//}("goat")

	c := make(chan string)
	foods := []string{"pizza", "rendang", "sop", "nasi", "kue", "croffle"}
	wg.Add(len(foods))
	for _, food := range foods {
		go func(f string) {
			cook(f, c)
			wg.Done()
		}(food)
	}
	msg := <-c
	celebratingNasiMatang(msg)
	wg.Wait()

	//go f("goroutine")
	//f("direct")
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func count(thing string) {
	for i := 1; i <=5; i++ {
		fmt.Println(i, thing)
		//c <- thing
		time.Sleep(time.Millisecond * 500)
	}
	//close(c)
}

func cook(food string, c chan string) {
	messsage := fmt.Sprintf("Cooking %s...\n", food)
	fmt.Println(messsage)

	time.Sleep(time.Second * 5)
	fmt.Println("Finished cooking ", food)
	if food == "nasi" {
		c <- food
	}

}

func celebratingNasiMatang(c string) {
	fmt.Println("yeay nasi matang")
	fmt.Println(c)

}

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
		time.Sleep(time.Millisecond * 500)
	}
}