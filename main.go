package main

import (
	"fmt"
)

var ch1 = make(chan int)
var ch2 = make(chan int)

func philosopher(seat int) {
	count := 0
	forks := 0

	for count < 3 {
		fork := <-ch1

		if fork == seat || fork == (seat+1)%5 {
			forks++
			fmt.Println("Philosopher", seat, "get fork", fork)
		} else {
			ch2 <- fork
		}

		if forks == 2 {
			count++
			fmt.Println("Philosopher", seat, "has eaten", count, "portion")
			ch2 <- seat
			ch2 <- (seat + 1) % 5
		}
	}
}

func fork(seat int) {
	inUse := false

	for {
		if !inUse {
			ch1 <- seat
			inUse = true
		} else {
			x := <-ch2

			if x != seat {
				ch2 <- x
			} else {
				inUse = false
			}
		}
	}

}

func main() {
	for i := 0; i < 5; i++ {
		go fork(i)
		go philosopher(i)
	}

	for {

	}
}
