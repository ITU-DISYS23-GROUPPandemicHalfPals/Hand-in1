/*
// A deadlock occurs in a case where every philosopher takes a fork to the same side of them.
// Therefore we have to figure out how to avoid that problem.
//
// The way we chose to solve this problem, is for every philosopher to take the fork to the left of them.
// Then they check if the right fork is available, if it is the philosopher will eat.
// When the right fork is not available, the philosopher will put the left fork back.
//
// This will make sure that a fork will never be locked with one philosopher.
*/

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
	for {
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

		if eatCount == portions {
			break
		}

		fmt.Println("Philosopher", position, "is thinking.")
		time.Sleep(wait)
	}

	fmt.Println("Philosopher", position, "is done.")
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

	fmt.Println("Every philosopher is done.")
}
