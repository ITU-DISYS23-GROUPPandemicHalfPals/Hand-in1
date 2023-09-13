package main

import (
	"fmt"
)

var channel1 = make(chan int)
var channel2 = make(chan int)

func philosopher2(philosopherPosition int) {
	portionCount := 0
	forkCount := 0

	for portionCount < 3 {
		forkPosition := <-channel1

		if forkPosition == philosopherPosition || forkPosition == (philosopherPosition+1)%5 {
			forkCount++
			fmt.Println("Philosopher", philosopherPosition, "accepted fork", forkPosition)
		} else {
			channel2 <- forkPosition
		}

		if forkCount == 2 {
			portionCount++
			forkCount = 0
			fmt.Println("Philosopher", philosopherPosition, "has eaten", portionCount, "portion")
			channel2 <- philosopherPosition
			channel2 <- (philosopherPosition + 1) % 5
		}
	}

	fmt.Println("Philosopher", philosopherPosition, "is done")
}

func fork2(forkPosition int) {
	forkInUse := false

	for {
		if !forkInUse {
			channel1 <- forkPosition
			forkInUse = true
		} else {
			x := <-channel2

			if x != forkPosition {
				channel2 <- x
			} else {
				forkInUse = false
			}
		}
	}
}

func main2() {
	for i := 0; i < 5; i++ {
		go fork2(i)
		go philosopher2(i)
	}

	for {
	}
}
