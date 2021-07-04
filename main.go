package main

import (
	"fmt"
	"time"
)

func go1 (s chan int) {
	fmt.Println ("1")
	s <- 1
}

func go2 (s chan int) {
	fmt.Println ("2")
	s <- 2
}

func go3 (s chan int) {
	//fmt.Println ("3")
	s <- 3
}

func go4 (s chan int) {
	//fmt.Println ("4")
	s <- 4
}

func main() {
	step1 := make (chan int)
	step2 := make (chan int)
	go go1(step1)
	go go2(step2)
	a := <-step2
	fmt.Println (a)
	time.Sleep (time.Second) // here go1() goroutine execute

}