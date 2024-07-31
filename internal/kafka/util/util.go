package util

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
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

func KafkaConfig(cfg *config.Config) kafka.ConfigMap {
	kc := make(map[string]kafka.ConfigValue)

	if cfg.Kafka.IsLocal {
		kc["bootstrap.servers"] = cfg.Kafka.Brokers
	}

	return kc
}

func Init(cfg *config.Config) {
	mutex.Lock()
	defer mutex.Unlock()

	producerCfg := KafkaConfig(cfg)
	producerCfg["acks"] = "all"
	p, err := kafka.NewProducer(&producerCfg)
	common.ChkFatal(err, "Failed to create producer")

	consumerCfg := KafkaConfig(cfg)
	consumerCfg["group.id"] = cfg.Kafka.GroupId
	consumerCfg["auto.offset.reset"] = "earliest"
	consumerCfg["allow.auto.create.topics"] = true
	c, err := kafka.NewConsumer(&consumerCfg)
	common.ChkFatal(err, "Failed to create consumer")

	_producer = p
	_consumer = c
}

func Singleton() Instance {
	mutex.Lock()
	defer mutex.Unlock()

	return Instance{Producer: _producer, Consumer: _consumer}
}
