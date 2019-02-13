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
		for i := 0; i < 1; i++ {
			aWorker(aCh, aDone)
		}
		select {}
	}()
}

type aInstr struct {
	num int
	inp int
	rem string
}
type aResp struct {
	num, wNum int
	outp      int
	rem       string
}

// aBoss returns c, a work assignment channel
// and d, a work done (result) channel.
func aBoss() (chan *aInstr, chan *aResp) {
	c := make(chan *aInstr)
	d := make(chan *aResp)
	fmt.Println("aBoss starting")
	go func() {
		go func() {
			i := 0
			for {
				ins := aInstr{i, i, "I want it yesterday!"}
				c <- &ins
				fmt.Printf("aBoss sent %v\n", ins)
				i++
			}
		}()
		for {
			fmt.Printf("aBoss received %v\n", <-d)
		}
	}()
	return c, d
}

func aWorker(c chan *aInstr, d chan *aResp) {
	fmt.Println("aWorker starting")
	go func() {
		for {
			ins := <-c
			res := aResp{ins.num, ins.num, ins.num * 2, ""}
			res.rem = fmt.Sprintf("The result for request %v is %v", res.wNum, res.outp)
			time.Sleep(time.Second)
			d <- &res
		}
	}()
}
