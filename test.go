package main

import (
	"fmt"
	"time"
)

const portions = 3
const count = 5
const wait = 100

type channels struct {
	in  []chan bool
	out []chan bool
}

func philosopher(position int, c *channels) {
	fmt.Println("Philosopher", position, "is thinking.")

	eatCount := 0
	for eatCount < portions {
		time.Sleep(wait)

		if len(c.in[position]) == 1 {
			<-c.in[position]
		} else {
			continue
		}

		if len(c.in[(position+1)%count]) == 1 {
			<-c.in[(position+1)%count]
		} else {
			c.out[position] <- true
			continue
		}

		eatCount++
		fmt.Println("Philosopher", position, "is eating. Eat count =", eatCount)
		time.Sleep(wait)
		fmt.Println("Philosopher", position, "is thinking.")

		c.out[position] <- true
		c.out[(position+1)%count] <- true
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
		go philosopher(i, c)
		go fork(i, c)
	}

	for {
	}
}
