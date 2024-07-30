package util

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/kanataidarov/gorm_kafka_docker/internal/config"
	"github.com/kanataidarov/gorm_kafka_docker/pkg/common"
	"sync"
)

var (
	_producer *kafka.Producer
	_consumer *kafka.Consumer
	mutex     sync.Mutex
)

type Instance struct {
	Producer *kafka.Producer
	Consumer *kafka.Consumer
}

func Init(cfg *config.Config) {
	mutex.Lock()
	defer mutex.Unlock()

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":        cfg.Kafka.Brokers,
		"acks":                     "all",
		"allow.auto.create.topics": true})
	common.ChkFatal(err, "Failed to create producer")

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.Kafka.Brokers,
		"group.id":          cfg.Kafka.GroupId,
		"auto.offset.reset": "earliest"})
	common.ChkFatal(err, "Failed to create consumer")

	_producer = p
	_consumer = c
}

func Singleton() Instance {
	mutex.Lock()
	defer mutex.Unlock()

	return Instance{Producer: _producer, Consumer: _consumer}
}
