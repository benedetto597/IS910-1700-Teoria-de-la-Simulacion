package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("%s: %s", "Error connecting to RabbitMQ", err)
	}
	defer connection.Close()

	// failOnError(err, "Failed to connect to RabbitMQ")

	channel, err := connection.Channel()
	if err != nil {
		log.Fatalf("%s: %s", "Error connecting to the channel", err)
	}
	defer channel.Close()

	queue, err := channel.QueueDeclare(
		"IS-910", // name of the queue
		false,    // durable or persistent
		false,    // delete when is used
		false,    // exclusive
		false,    // timeout
		nil,      // arguments
	)

	if err != nil {
		log.Fatalf("%s: %s", "Error creating the queue", err)
	}

	message, err := channel.Consume(
		queue.Name, // name of the queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)

	if err != nil {
		log.Fatalf("%s: %s", "Error creating the consume channel", err)
	}

	// Block main goroutine
	forever := make(chan bool)

	// Anonymous function to recive messages
	go func() {
		for queue := range message {
			fmt.Printf("\n%s: %s", "Message received", string(queue.Body))
		}
	}()

	<-forever
}
