package main

import (
	"fmt"
	"github.com/kanataidarov/gorm_kafka_docker/internal/config"
	"github.com/kanataidarov/gorm_kafka_docker/internal/handler"
	"github.com/kanataidarov/gorm_kafka_docker/internal/kafka/consumer"
	kfk "github.com/kanataidarov/gorm_kafka_docker/internal/kafka/util"
	"github.com/kanataidarov/gorm_kafka_docker/pkg/common"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

func main() {
	log.Println("Starting messaggio_assignment")
	cfg := config.Load()

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s TimeZone=%s sslmode=%s",
		cfg.Db.Host, cfg.Db.Port, cfg.Db.User, cfg.Db.Password, cfg.Db.DbName, cfg.Db.Tz, cfg.Db.Ssl)
	dbase, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	common.ChkFatal(err, "Failed to connect to database")

	kfk.Init(cfg)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()

		http.HandleFunc("/applications", handler.ApplicationsHandler(cfg, dbase))

		var addr string
		if cfg.Kafka.IsLocal {
			addr = fmt.Sprintf(":%d", getPort(cfg))
		} else {
			addr = fmt.Sprintf("%s:%d", cfg.Handler.Host, getPort(cfg))
		}

		common.ChkFatal(http.ListenAndServe(addr, nil), "Failed to start web handler")
		log.Println("Handler is running on " + addr)
	}()

	go func() {
		defer wg.Done()

		consumer.Handler(cfg, dbase)
	}()

	wg.Wait()

	kfk.Singleton().Producer.Close()
	log.Println("Stopping messaggio_assignment")
}

func getPort(cfg *config.Config) int {
	if envVar, ok := os.LookupEnv("PORT"); ok {
		port, err := strconv.Atoi(envVar)
		if err != nil {
			common.ChkWarn(err, fmt.Sprintf("Failed to convert PORT envvar. Setting handler port: %d", cfg.Handler.Port))
			return cfg.Handler.Port
		}
		return port

	}
	return cfg.Handler.Port
}
