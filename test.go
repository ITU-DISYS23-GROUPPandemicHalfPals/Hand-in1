package main

import (
	"fmt"
	"time"
)

const count = 5

type channels struct {
	in  []chan bool
	out []chan bool
}

func philosopher(position int, c *channels) {
	fmt.Println("Philosopher", position, "is thinking.")

	eats := 0
	for eats < 3 {
		time.Sleep(1000)

		leftFork := false
		rightFork := false

		if len(c.in[position]) == 1 {
			leftFork = <-c.in[position]
		}

		if !leftFork {
			continue
		}

		if len(c.in[(position+1)%count]) == 1 {
			rightFork = <-c.in[(position+1)%count]
		}

		if !rightFork {
			c.out[position] <- true
			continue
		}

		if leftFork && rightFork {
			eats++
			fmt.Println("Philosopher", position, "is eating.")
			time.Sleep(1000)
			fmt.Println("Philosopher", position, "is thinking.")

			c.out[position] <- true
			c.out[(position+1)%count] <- true
		}
	}
}

func fork(position int, c *channels) {
	for {
		c.in[position] <- true
		<-c.out[position]
	}
}

func main() {
	c := new(channels)

	c.in = make([]chan bool, count)
	c.out = make([]chan bool, count)

	for i := 0; i < count; i++ {
		c.in[i] = make(chan bool, 1)
		c.out[i] = make(chan bool, 1)
	}

	for i := 0; i < count; i++ {
		go philosopher(i, c)
		go fork(i, c)
	}

	for {
	}
}
