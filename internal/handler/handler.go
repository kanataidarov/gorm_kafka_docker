package handler

import (
	"encoding/json"
	"fmt"
	"github.com/kanataidarov/gorm_kafka_docker/internal/config"
	"github.com/kanataidarov/gorm_kafka_docker/internal/db"
	"github.com/kanataidarov/gorm_kafka_docker/internal/kafka/producer"
	"github.com/kanataidarov/gorm_kafka_docker/pkg/common"
	"gorm.io/gorm"
	"net/http"
)

type ApplicationRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Position string `json:"position"`
}

func CreateApplication(cfg *config.Config, dbase *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var request ApplicationRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			common.ChkWarn(err, "Invalid request body")
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		if request == (ApplicationRequest{}) {
			common.ChkWarn(err, "Either Name, Email, or Position are empty")
			http.Error(w, "name, email, and position are required", http.StatusBadRequest)
			return
		}

		assignment, err := db.LastAssignment(dbase, request.Position)
		common.ChkWarn(err, fmt.Sprintf("No assignment found for \"%s\"", request.Position))

		application := db.Application{
			Name:       request.Name,
			Email:      request.Email,
			Position:   request.Position,
			Assignment: *assignment,
		}
		err = db.Create(dbase, application)
		if err != nil {
			http.Error(w, "Couldn't process application", http.StatusInternalServerError)
			return
		}

		_ = producer.Push(cfg, application)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Application processed successfully",
		})
		common.ChkWarn(err, "Error encoding response")
	}
}
