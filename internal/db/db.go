package db

import (
	"fmt"
	"github.com/kanataidarov/gorm_kafka_docker/pkg/common"
	"gorm.io/gorm"
	"time"
)

type Application struct {
	gorm.Model
	Name           string
	Email          string
	Position       string
	AssignmentID   uint
	Assignment     Assignment `gorm:"foreignKey:AssignmentID;references:ID"`
	AssignmentSent *time.Time
}

type Assignment struct {
	gorm.Model
	Position string
	Version  uint
	Doc      []byte
}

func LastAssignment(dbase *gorm.DB, position string) (*Assignment, error) {
	var assignment Assignment
	if err := dbase.Order("version desc").Last(&assignment, "position = ?", position).Error; err != nil {
		return nil, err
	}

	return &assignment, nil
}

func CreateApplication(dbase *gorm.DB, application Application) (Application, error) {
	err := dbase.Create(&application).Error
	common.ChkWarn(err, "Error during creation of application")

	return application, err
}

func PatchApplication(dbase *gorm.DB, application Application) (Application, error) {
	err := dbase.Model(&application).Update("assignment_sent", time.Now()).Error
	common.ChkWarn(err, fmt.Sprintf("Error during patching of application. ID=%d", application.ID))

	return application, err
}
