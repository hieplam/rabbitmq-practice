package main

import (
	"fmt"
	"log"
	"os"
	"rabbitmq-practice/utils"

	"github.com/streadway/amqp"
)

func main() {
	//Create Conn
	fmt.Println("env:", os.Getenv("AMQP_URL"))
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	defer conn.Close()
	utils.FailOnError(err, "Failed to connect to rabbitmq server")

	//Create channel
	ch, err := conn.Channel()
	defer ch.Close()

	//Declare Exchange
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

	//Publish message to exchange.
	body := utils.BodyFrom(os.Args)
	err = ch.Publish(
		"logs", // exchange
		"",     // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	utils.FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)

}
