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
	g := dflt.EnvString("greet", "hi")
	publisher()
	subscriber()
	webServer()
	for {
		fmt.Println("skaffold-hello", g)
		time.Sleep(3 * time.Second)
	}
}
func publisher() {
	go func() {
		url := dflt.EnvString("NATS", "nats://sk-t-nats-streaming:4222")
		c, err := stan.Connect("test-cluster", "stan-test-publisher", stan.NatsURL(url))
		if err != nil {
			log.Fatalf("nats-streaming conn: %s: %v", url, err)
		}
		defer c.Close()
		for {
			c.Publish("topic1", []byte(time.Now().Format("15:04:05.000")))
			time.Sleep(time.Second)
		}
	}()
}

func subscriber() {
	go func() {
		url := dflt.EnvString("NATS", "nats://sk-t-nats-streaming:4222")
		c, err := stan.Connect("test-cluster", "stan-test-subscriber", stan.NatsURL(url))
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
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "hello this is from go")
		})
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()
}
