package producer

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/kanataidarov/gorm_kafka_docker/internal/config"
	"github.com/kanataidarov/gorm_kafka_docker/internal/db"
	kfk "github.com/kanataidarov/gorm_kafka_docker/internal/kafka/util"
	"github.com/kanataidarov/gorm_kafka_docker/pkg/common"
)

func Push(cfg *config.Config, application db.Application) error {
	data, err := json.Marshal(application)
	if err != nil {
		common.ChkWarn(err, "Error marshalling application to kafka message")
		return err
	}

	producer := kfk.Singleton().Producer

	err = producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &cfg.Kafka.Topic, Partition: kafka.PartitionAny},
		Key:            []byte(fmt.Sprintf("%d", application.ID)),
		Value:          []byte(fmt.Sprintf("%s", string(data))),
	}, nil)
	if err != nil {
		common.ChkWarn(err, fmt.Sprintf("Produce of the message failed. ID=%d", application.ID))
		return err
	}

	producer.Flush(9999)

	return nil
}
