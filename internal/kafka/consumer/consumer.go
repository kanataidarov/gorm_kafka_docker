package consumer

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/kanataidarov/gorm_kafka_docker/internal/config"
	"github.com/kanataidarov/gorm_kafka_docker/internal/db"
	kfk "github.com/kanataidarov/gorm_kafka_docker/internal/kafka/util"
	"github.com/kanataidarov/gorm_kafka_docker/pkg/common"
	"gorm.io/gorm"
	"log"
)

func Handler(cfg *config.Config, dbase *gorm.DB) {
	consumer := kfk.Singleton().Consumer

	err := consumer.SubscribeTopics([]string{cfg.Kafka.Topic}, nil)
	common.ChkWarn(err, fmt.Sprintf("Error subscribing to topic \"%s\"", cfg.Kafka.Topic))

	sigChan := common.SysInterrupt()

	run := true
	for run {
		select {
		case sig := <-sigChan:
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
				markSent(dbase, e.Value)
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

func markSent(dbase *gorm.DB, msg []byte) {
	var application *db.Application
	if err := json.Unmarshal(msg, &application); err != nil {
		common.ChkWarn(err, "Error unmarshalling kafka message to application")
		return
	}

	updated, _ := db.PatchApplication(dbase, *application)
	log.Printf("Application ID=%d marked sent\n", updated.ID)
}
