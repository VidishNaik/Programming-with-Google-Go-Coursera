package main

import (
	"fmt"
	"sync"
	"time"
)

type fork struct{ sync.Mutex }

type philosopher struct {
	id                  int
	left, right *fork
}

// Goes from thinking to hungry to eating and done eating then starts over.
// Adapt the pause values to increased or decrease contentions
// around the forks.
func (p philosopher) eat() {
	for j := 0; j < 3; j++ {
		p.left.Lock()
		p.right.Lock()

		fmt.Println("Philosopher", p.id + 1, "is starting to eat.")
		time.Sleep(time.Second)
		
		p.right.Unlock()
		p.left.Unlock()

		fmt.Println("Philosopher", p.id + 1, "is finishing eating.")
		time.Sleep(time.Second)
	}
	wg.Done()
}

func printStatement(action string, id int) {
	fmt.Printf("Philosopher #%d is %s\n", id+1, action)
}

var wg sync.WaitGroup

func main() {
	// How many philosophers and forks

	count := 5

	// Create forks
	forks := make([]*fork, count)
	for i := 0; i < count; i++ {
		forks[i] = new(fork)
	}

	// Create philospoher, assign them 2 forks and send them to the dining table
	philosophers := make([]*philosopher, count)
	for i := 0; i < count; i++ {
		philosophers[i] = &philosopher{
			id: i, left: forks[i], right: forks[(i+1)%count]}
		wg.Add(1)
		go philosophers[i].eat()

	}
	wg.Wait()

}
