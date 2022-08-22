package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/nats-io/nats.go"
)

func main() {
	natsConn, err := nats.Connect("nats://localhost:4222")

	if err != nil {
		log.Fatal("連不上nats")
	}
	defer natsConn.Close()

	wg := sync.WaitGroup{}
	wg.Add(1)

	if _, err := natsConn.Subscribe("subject", func(msg *nats.Msg) {
		wg.Done()
		fmt.Println("收到", string(msg.Data))
	}); err != nil {
		log.Fatal(err)
	}

	wg.Wait()

	/*
		_, err = natsConn.QueueSubscribe("subject", "queue", func(msg *nats.Msg) {
			fmt.Println("收到", string(msg.Data))
		})
		if err != nil {
			log.Fatal("訂閱失敗")
		}
	*/
}
