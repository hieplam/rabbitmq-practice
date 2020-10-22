package main

import (
	"log"
	"os"
	"rabbitmq-practice/utils"

	"github.com/streadway/amqp"
)

//Publisher
func main() {
	//Create Connection
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	utils.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	//Create Channel
	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	//Create Queue
	q, err := ch.QueueDeclare(
		"task_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	utils.FailOnError(err, "Failed to declare a queue")
	body := utils.BodyFrom(os.Args)

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})

	log.Printf(" [x] Sent %s", body)
	utils.FailOnError(err, "Failed to publish a message")
}
