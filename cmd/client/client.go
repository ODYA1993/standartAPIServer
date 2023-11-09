package main

import (
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

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatal("не удалось зарегистрировать потребителя")
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Получено сообщение: %s", d.Body)
		}
	}()

	log.Printf(" [*] Ждем сообщений. Для выхода нажмите: CTRL+C")
	<-forever
}
