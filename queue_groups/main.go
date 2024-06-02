package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 3; i++ {
		_, err := nc.QueueSubscribe("updates", "workers", func(workerId int) nats.MsgHandler {
			return func(msg *nats.Msg) {
				log.Printf("Worker Id: %v:: Recieved message: %s", workerId, string(msg.Data))
			}
		}(i+1))

		if err != nil {
			log.Fatal(err)
		}
	}
	select {}
}
