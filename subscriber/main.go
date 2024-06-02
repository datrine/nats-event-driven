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

	nc.Subscribe("updates", func(msg *nats.Msg) {
		log.Printf("Received Message: %s\n", string(msg.Data))
	})

	log.Println("Subscribed to 'updates' subject")

	select {}
}
