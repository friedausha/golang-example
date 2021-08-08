package main

import (
	"fmt"
	"golang-example/addition"
)

func main() {
	res, _ := addition.Add(11, 15)
	fmt.Println(res)
}
