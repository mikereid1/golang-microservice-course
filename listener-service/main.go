package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"math"
	"os"
	"time"
)

func main() {
	// connect to rabbitmq
	conn, err := connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer conn.Close()
	log.Println("Connected to RabbitMQ")

	// start listening for messages

	// create consumer

	// watch the queue and consume events

	// do something with it
}

func connect() (*amqp.Connection, error) {
	var count int64
	var backoff = 1 * time.Second
	var connection *amqp.Connection

	for {
		c, err := amqp.Dial("amqp://guest:guest@localhost")
		if err != nil {
			fmt.Println("RabbitMQ not yet ready")
			count++
		} else {
			connection = c
			break
		}

		if count > 5 {
			fmt.Println(err)
			return nil, err
		}

		backoff = time.Duration(math.Pow(float64(count), 2)) * time.Second
		log.Println("backing off...")
		time.Sleep(backoff)
	}

	return connection, nil
}
