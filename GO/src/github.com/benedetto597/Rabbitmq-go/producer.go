package main

import (
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

	message := "Simulating sending messages from queues"

	err = channel.Publish(
		"",         // exchange
		queue.Name, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message), // message body parsed as byte array
		})

	if err != nil {
		log.Fatalf("%s: %s", "Error sending message", err)
	}

	log.Printf("%s: %s", "Message sent", message)
}
