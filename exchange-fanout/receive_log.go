package main

import (
	"log"
	"rabbitmq-practice/utils"

	"github.com/streadway/amqp"
)

func main() {
	//Create conn
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	utils.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	//create channel
	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
	defer conn.Close()

	//declare exchange
	err = ch.ExchangeDeclare(
		"logs",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	utils.FailOnError(err, "Failed to declare an exchange")

	//declare queue
	q, err := ch.QueueDeclare(
		"",    // name -> random queue name.
		false, // durable
		false, // delete when unused
		true,  // exclusive -> used by only one connection and the queue will be deleted when that connection closes
		false, // no-wait
		nil,   // arguments
	)
	utils.FailOnError(err, "Failed to declare a queue")

	//bind queue
	err = ch.QueueBind(
		q.Name, // queue name
		"",     // routing key
		"logs", // exchange
		false,
		nil,
	)
	utils.FailOnError(err, "Failed to bind a queue")

	//consume queue
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	utils.FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
