package main

/**
Suppose we have five philosophers around a table and each of them holds only one fork.
But each one requires two forks in hands simultaneously to eat.
If each philosopher is competitive, they are not happy to borrow their forks to their neighbors.
The dining will not finished, which can be treated as a deadlock in programming.
There is an another situation where each philosophers is willing to lend their forks to others. But if they give up their own forks and pick up their neighbors’ forks at exactly same time, they still cannot eat, which is treated as a live lock in programming. The solution to finish this meal is allowing each philosophers to hold their own forks for a random length of time. After timing out, if they still cannot borrow their neighbors’ forks, they have to give up their own
forks letting their neighbors to use.
Under this way, these philosophers can finish this meal finally.
*/

import (
	"fmt"
	"math/rand"
	"time"
)

// Philosopher struct contains name, fork channel, neighbor philosopher
// fork channel is used for demonstrate if the fork of the philosopher available
type Philosopher struct {
	name     string
	fork     chan bool
	neighbor *Philosopher
}

// Make channel Philosopher
func makePhilosopher(name string, neighbor *Philosopher) *Philosopher {
	phil := &Philosopher{name, make(chan bool, 1), neighbor}
	phil.fork <- true
	return phil
}

// Simulate eating
func (phil *Philosopher) eat() {
	fmt.Printf("%v is eating \n", phil.name)
	time.Sleep(time.Duration(rand.Int63n(1e9)))
}

// Simulate thinking
func (phil *Philosopher) think() {
	fmt.Printf("%v is thinking.\n", phil.name)
	time.Sleep(time.Duration(rand.Int63n(1e9)))
}

// Get forks
func (phil *Philosopher) getForks() {
	timeOut := make(chan bool, 1)
	go func() { time.Sleep(1e9); timeOut <- true }()
	// fork is not available
	<-phil.fork
	fmt.Printf("%v now got his fork. \n", phil.name)
	select {
	// if neighbor's fork is available
	case <-phil.neighbor.fork:
		fmt.Printf("%v got %v's fork.\n", phil.name, phil.neighbor.name)
		fmt.Printf("%v has two fork.\n", phil.name)
		return
	case <-timeOut:
		phil.fork <- true
		phil.think()
		phil.getForks()
	}
}

func (phil *Philosopher) returnForks() {
	// after a philosopher finish eating, making his fork channel
	phil.fork <- true
	phil.neighbor.fork <- true
}

func (phil *Philosopher) dine(announce chan *Philosopher) {
	// the whole procedure of dining
	phil.think()
	phil.getForks()
	phil.eat()
	phil.returnForks()
	announce <- phil
}

func main() {
	names := []string{"#1", "#2", "#3", "#4",
		"#5", "#6"}
	philosophers := make([]*Philosopher, len(names))
	var phil *Philosopher
	// link all philosophers together
	for i, name := range names {
		phil = makePhilosopher(name, phil)
		philosophers[i] = phil
	}
	// let the first philosopher to be the neighbor of the last philosopher
	philosophers[0].neighbor = phil
	fmt.Printf("There are %v philosophers sitting at a table.\n", len(names))
	fmt.Println("They each have one fork, and must borrow from their neighbor to eat.")
	announce := make(chan *Philosopher)
	for _, phil := range philosophers {
		// dine occur concurrently
		go phil.dine(announce)
	}
	// the announce channel will get the philosophers who finish dining sequentially in concurrent dine()
	// print out them concurrently
	for i := 0; i < len(names); i++ {
		phil := <-announce
		fmt.Printf("%v is done dining. \n", phil.name)
	}
}
