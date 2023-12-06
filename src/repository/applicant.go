package repository

import (
	"gorm.io/gorm"
)

type Applicant struct {
	gorm.Model
	Name       string      `json:"name"`
	Email      string      `json:"email" gorm:"unique_index"`
	Phone      string      `json:"phone" gorm:"unique_index"`
	Resume     string      `gorm:"optional" json:"resume"`
	AppliedJob string      `json:"appliedJob"`
	Interviews []Interview `gorm:"foreignkey:ApplicantID"`
}

func (receiver *Applicant) Save() error {
	return db.Create(receiver).Error
}

func (receiver *Applicant) FindById(id uint) ([]Applicant, error) {
	var applicants []Applicant
	err := db.Preload("Interviews").Where("id = ?", id).Find(&applicants).Error
	return applicants, err
}
