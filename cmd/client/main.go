package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

func nat() {
	nc, _ := nats.Connect(nats.DefaultURL)

	// Simple Publisher
	nc.Publish("foo", []byte("Hello World"))

	// Simple Async Subscriber
	nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	// Responding to a request message
	nc.Subscribe("request", func(m *nats.Msg) {
		m.Respond([]byte("answer is 42"))
	})

	// Simple Sync Subscriber
	sub, _ := nc.SubscribeSync("foo")
	//m, err := sub.NextMsg(timeout)

	// Channel Subscriber
	ch := make(chan *nats.Msg, 64)
	sub, _ = nc.ChanSubscribe("foo", ch)
	//msg := <-ch

	// Unsubscribe
	sub.Unsubscribe()

	// Drain
	sub.Drain()

	// Requests
	//msg, _ := nc.Request("help", []byte("help me"), 10*time.Millisecond)

	// Replies
	nc.Subscribe("help", func(m *nats.Msg) {
		nc.Publish(m.Reply, []byte("I can help!"))
	})

	// Drain connection (Preferred for responders)
	// Close() not needed if this is called.
	nc.Drain()

	// Close connection
	nc.Close()
}

func main() {
	//nat()

	natsConn, err := nats.Connect("nats://localhost:4222")
	//natsConn, err := nats.Connect("demo.nats.io")
	if err != nil {
		log.Fatal("連不上nats")
	}
	defer natsConn.Close()

	err = natsConn.Publish("subject", []byte("hello world"))
	if err != nil {
		log.Fatal("送不出去")
	}

	err = natsConn.Flush()
	if err != nil {
		log.Fatal("清空吃敗")
	}
	/*
		_, err = natsConn.Subscribe("subject", func(msg *nats.Msg) {
			fmt.Println("收到", string(msg.Data))
		})
		if err != nil {
			log.Fatal("訂閱失敗")
		}
	*/

}
