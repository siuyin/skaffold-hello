package main

import (
	"fmt"
	"time"
)

// aWork comprises aBoss and one or more
// aWorkers.
func aWork() {
	go func() {
		aCh, aDone := aBoss()
		for i := 0; i < 3; i++ {
			aWorker(aCh, aDone)
		}
		select {}
	}()
}

// aBoss returns c, a work assignment channel
// and d, a work done (result) channel.
func aBoss() (chan int, chan int) {
	c := make(chan int)
	d := make(chan int)
	fmt.Println("aBoss starting")
	go func() {
		go func() {
			i := 0
			for {
				c <- i
				fmt.Printf("aBoss sent %d\n", i)
				i++
			}
		}()
		for {
			fmt.Printf("aBoss received %d\n", <-d)
		}
	}()
	return c, d
}

func aWorker(c, d chan int) {
	fmt.Println("aWorker starting")
	go func() {
		for {
			i := <-c
			time.Sleep(time.Second)
			d <- i * 2
		}
	}()
}
