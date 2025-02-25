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
	log.Println("Connected to NATS server")
	defer nc.Close()

	err = nc.Publish("updates", []byte("Hello, NATS!"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Published message to 'updates' subject")
}
