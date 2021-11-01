package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type idService interface {
    // Returns values in ascending order; it should be safe to call
    // getNext() concurrently without any additional synchronization.
    getNext() uint64
}


// STRATEGY 1: No synchronization
type NoSync struct {
	val	uint64
}

// TODO come back to how to get this to start at 1
func (noSync *NoSync) getNext() uint64 {
	var retVal = noSync.val
	noSync.val += 1
	return retVal
}


// STRATEGY 2: Atomic
type AtomicCounter struct {
	counter uint64
}

func (atomicCounter *AtomicCounter) getNext() uint64 {
	return atomic.AddUint64(&atomicCounter.counter, 1)
}


// STRATEGY 3: Mutex
type MutCounter struct {
	mu 		sync.Mutex
	counter uint64
}

func (mutCounter *MutCounter) getNext() uint64 {
	mutCounter.mu.Lock()
	mutCounter.counter += 1
	c := mutCounter.counter
	mutCounter.mu.Unlock()
	return c
}


// STRATEGY 4: Exclusive monitor
func monitor(in chan struct{}, out chan uint64) uint64 {
	var counter uint64
	for {
		<-in
		counter += 1
		out <- counter
	}
}

type MonitoredCounter struct {
	in 		chan struct{}
	out 	chan uint64
}

func (monitoredCounter *MonitoredCounter) getNext() uint64 {
	monitoredCounter.in <- struct{}{}
	return <- monitoredCounter.out
}

func main() {

	myNoSync := new(NoSync)
	fmt.Println(myNoSync.getNext())
	fmt.Println(myNoSync.getNext())
	fmt.Println(myNoSync.getNext())
	fmt.Println()

	atomicCounter := new(AtomicCounter)
	fmt.Println(atomicCounter.getNext())
	fmt.Println(atomicCounter.getNext())
	fmt.Println(atomicCounter.getNext())
	fmt.Println()

	mutCounter := new(MutCounter)
	fmt.Println(mutCounter.getNext())
	fmt.Println(mutCounter.getNext())
	fmt.Println(mutCounter.getNext())
	fmt.Println()

	monitoredCounter := new(MonitoredCounter)
	monitoredCounter.in = make(chan struct{})
	monitoredCounter.out = make(chan uint64)
	go monitor(monitoredCounter.in, monitoredCounter.out)
	fmt.Println(monitoredCounter.getNext())
	fmt.Println(monitoredCounter.getNext())
	fmt.Println(monitoredCounter.getNext())
	fmt.Println()

}