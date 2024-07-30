package main

import (
	"github.com/kanataidarov/gorm_kafka_docker/internal/handler"
	"github.com/kanataidarov/gorm_kafka_docker/internal/kafka/consumer"
	"github.com/kanataidarov/gorm_kafka_docker/pkg/common"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"sync"
)

func main() {
	dsn := "host=localhost port=44048 user=postgres password=changeme dbname=applications sslmode=disable TimeZone=Asia/Qyzylorda"
	dbase, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	common.ChkFatal(err, "Failed to connect to database")

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		http.HandleFunc("/applications", handler.CreateApplication(dbase))

		addr := ":44049"
		log.Printf("Handler is running on: %s\n", addr)
		common.ChkFatal(http.ListenAndServe(addr, nil), "Failed to start web handler")
	}()

	go func() {
		defer wg.Done()
		consumer.Handler()
	}()

	wg.Wait()
}
