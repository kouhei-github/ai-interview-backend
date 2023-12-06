package repository

import (
	"gorm.io/gorm"
)

type Interview struct {
	gorm.Model
	ApplicantID         uint   `json:"applicant_id"`
	InterviewType       string `json:"interview_type"`
	InterviewStatus     int    `json:"interview_status"`
	Applicant           Applicant
	InterviewEvaluation InterviewEvaluation
}

func (receiver *Interview) Save() error {
	return db.Create(receiver).Error
}

func (receiver *Interview) FindById(id uint) ([]Interview, error) {
	var interviews []Interview
	err := db.Preload("Applicant").Preload("InterviewEvaluation").Where("id = ?", id).Find(&interviews).Error
	return interviews, err
}
