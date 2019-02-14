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

	pub.Work()
	web.Serve()
	work.Work()
	pipe.Work()

	g := dflt.EnvString("greet", "hello")
	for {
		fmt.Println("skaffold-hello", g)
		time.Sleep(5 * time.Second)
	}
}
