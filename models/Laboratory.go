package models

import "time"

type RegisterLabType struct {
	ID        uint   `json:"id" gorm:"type:int(20)"`
	LabNumber string `json:"lab_number" gorm:"type:varchar(6)"`
	Applicant string `json:"applicant"`
	StudentId string `json:"student_id"`
	Password  string `json:"password"`
	Reason    string `json:"reason"`
	CreatedAt time.Time
}
