package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/kanataidarov/gorm_kafka_docker/pkg/common"
)

type Config struct {
	Db      Db
	Kafka   Kafka
	Handler Handler
}

type Db struct {
	Host     string `env:"DB_HOST" envDefault:"localhost"`
	Port     int    `env:"DB_PORT" envDefault:"44048"`
	DbName   string `env:"DB_NAME" envDefault:"applications"`
	User     string `env:"DB_USER" envDefault:"postgres"`
	Password string `env:"DB_PWD"`
	Tz       string `env:"DB_TZ" envDefault:"Asia/Qyzylorda"`
	Ssl      string `env:"DB_SSL" envDefault:"disable"`
}

type Handler struct {
	Port int `env:"SRV_PORT" envDefault:"44049"`
}

type Kafka struct {
	Brokers string `env:"BROKERS" envDefault:"localhost:9092"`
	Topic   string `env:"TOPIC" envDefault:"applications"`
	GroupId string `env:"GROUP_ID" envDefault:"applications"`
}

func Load() *Config {
	var cfg Config
	common.ChkFatal(env.Parse(&cfg), "Could not read config")

	return &cfg
}
