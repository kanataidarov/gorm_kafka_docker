package consumer

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/kanataidarov/gorm_kafka_docker/internal/config"
	kfk "github.com/kanataidarov/gorm_kafka_docker/internal/kafka/util"
	"github.com/kanataidarov/gorm_kafka_docker/pkg/common"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Handler(cfg *config.Config) {
	consumer := kfk.Singleton().Consumer

	err := consumer.SubscribeTopics([]string{cfg.Kafka.Topic}, nil)

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	run := true
	for run {
		select {
		case sig := <-sigchan:
			log.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev := consumer.Poll(99)
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

	err = consumer.Close()
	common.ChkFatal(err, "Error closing consumer")
}
