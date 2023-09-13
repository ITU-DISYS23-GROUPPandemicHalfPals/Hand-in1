package main

import (
	"fmt"
	"sync"
	"time"
)

const portions = 3
const count = 5
const wait = time.Second / 100

var feast sync.WaitGroup

type channels struct {
	in  []chan bool
	out []chan bool
}

func philosopher(position int, c *channels) {
	defer feast.Done()
	fmt.Println("Philosopher", position, "is thinking.")

	eatCount := 0
	for eatCount < portions {
		<-c.in[position]

		if len(c.in[(position+1)%count]) == 1 {
			<-c.in[(position+1)%count]
		} else {
			c.out[position] <- true
			continue
		}

		eatCount++
		fmt.Println("Philosopher", position, "is eating. Eat count =", eatCount)
		time.Sleep(wait)

		c.out[position] <- true
		c.out[(position+1)%count] <- true

		fmt.Println("Philosopher", position, "is thinking.")
		time.Sleep(wait)
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
		feast.Add(1)
	}

	feast.Wait()
}
