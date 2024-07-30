package consumer

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/kanataidarov/gorm_kafka_docker/pkg/common"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Handler() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9093",
		"group.id":          "applications02",
		"auto.offset.reset": "earliest"})
	common.ChkFatal(err, "Failed to create consumer")

	topic := "applications"
	err = c.SubscribeTopics([]string{topic}, nil)

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	run := true
	for run {
		select {
		case sig := <-sigchan:
			log.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev := c.Poll(100)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				log.Printf("Message on %s: %s\n", e.TopicPartition, string(e.Value))
			case kafka.Error:
				log.Printf("Error: %v: %v\n", e.Code(), e)
				if e.Code() == kafka.ErrAllBrokersDown {
					run = false
				}
			default:
				log.Printf("Ignored %v\n", e)
			}
		}
	}

	err = c.Close()
	common.ChkFatal(err, "Error closing consumer")
}
