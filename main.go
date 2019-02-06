package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	stan "github.com/nats-io/go-nats-streaming"
	"github.com/siuyin/dflt"
)

func main() {
	g := dflt.EnvString("greet", "hello")
	publisher()
	subscriber()
	webServer()
	aWork()
	for {
		fmt.Println("skaffold-hello", g)
		time.Sleep(3 * time.Second)
	}
}
func publisher() {
	go func() {
		url := dflt.EnvString("NATS", "nats://sk-t-nats-streaming:4222")
		c, err := stan.Connect("my-clust", "stan-test-publisher", stan.NatsURL(url))
		if err != nil {
			log.Fatalf("nats-streaming conn: %s: %v", url, err)
		}
		defer c.Close()
		for {
			c.Publish("topic1", []byte(time.Now().Format("15:04:05.000")))
			time.Sleep(5 * time.Second)
		}
	}()
}

func subscriber() {
	go func() {
		url := dflt.EnvString("NATS", "nats://sk-t-nats-streaming:4222")
		c, err := stan.Connect("my-clust", "stan-test-subscriber", stan.NatsURL(url))
		if err != nil {
			log.Fatalf("nats-streaming conn: %s: %v", url, err)
		}
		defer c.Close()

		_, err = c.Subscribe("topic1", func(m *stan.Msg) {
			fmt.Printf("%s\n", m.Data)

		}, stan.DurableName("sub1"), stan.DeliverAllAvailable())
		//		}, stan.DeliverAllAvailable())
		select {}
	}()
}

func webServer() {
	go func() {
		log.Println("webServer starting")
		http.HandleFunc("/ans", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "The answer is 42.")
		})
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Hello, this is from a go webserver running in kubernetes.")
		})
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()
}
