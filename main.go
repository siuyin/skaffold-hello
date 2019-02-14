package main

import (
	"fmt"
	"time"

	"hello/pipe"
	"hello/pub"
	"hello/web"
	"hello/work"

	"github.com/siuyin/dflt"
)

func main() {

	pub.Work()  // publish/subscribe system using nats streaming
	web.Serve() // a simple web server in go
	work.Work() // a worker system with aBoss and one or more workers
	pipe.Work() // a pipeline work system

	// good old main-line code
	g := dflt.EnvString("greet", "hello")
	for {
		fmt.Println("skaffold-hello", g)
		time.Sleep(5 * time.Second)
	}
}
