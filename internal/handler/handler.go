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

type ApplicationResponse struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Position  string `json:"position"`
	Processed bool   `json:"processed"`
}

func ApplicationsHandler(cfg *config.Config, dbase *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getApplications(dbase)(w, r)
		case http.MethodPost:
			createApplication(cfg, dbase)(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func createApplication(cfg *config.Config, dbase *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
		application, err = db.CreateApplication(dbase, application)
		if err != nil {
			http.Error(w, "Couldn't process application at creation stage", http.StatusInternalServerError)
			return
		}

		err = producer.Push(cfg, application)
		if err != nil {
			http.Error(w, "Couldn't process application at publication stage", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Application processed successfully",
			"assignment": map[string]interface{}{
				"position": application.Assignment.Position,
				"version":  application.Assignment.Version},
		})
		common.ChkWarn(err, "Error encoding response")
	}
}

func getApplications(dbase *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var applications []db.Application
		err := dbase.Find(&applications).Error
		if err != nil {
			common.ChkWarn(err, "Failed to retrieve applications")
			http.Error(w, "Failed to retrieve applications", http.StatusInternalServerError)
			return
		}

		var response []ApplicationResponse
		for _, app := range applications {
			response = append(response, ApplicationResponse{
				Name:      app.Name,
				Email:     app.Email,
				Position:  app.Position,
				Processed: app.AssignmentSent != nil,
			})
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(response)
		common.ChkWarn(err, "Error encoding response")
	}
}
