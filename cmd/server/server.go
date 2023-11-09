package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("не удалось подключиться к rabbitmq")
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("не удалось открыть канал")
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		log.Fatal("не удалось объявить очередь")
	}

	for i := 0; i < 10; i++ {
		msg := fmt.Sprintf("Привет Мир %d", i)

		if err := ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(msg),
			}); err != nil {
			log.Fatal("не удалось объявить очередь")
		}
	}
}
