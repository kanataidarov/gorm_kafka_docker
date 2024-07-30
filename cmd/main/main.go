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

		http.HandleFunc("/applications", handler.CreateApplication(cfg, dbase))

		port := cfg.Handler.Port
		log.Printf("Handler is running on port %d\n", port)
		common.ChkFatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil), "Failed to start web handler")
	}()

	go func() {
		defer wg.Done()

		consumer.Handler(cfg)
	}()

	wg.Wait()

	kfk.Singleton().Producer.Close()
	log.Println("Stopping messaggio_assignment")
}
