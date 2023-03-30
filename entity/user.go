package entity

import "time"

type User struct {
	ID             int64      `json:"id" gorm:"type:int;primaryKey;autoIncrement:false"`
	Name           string     `json:"name" gorm:"type:varchar(100);not null"`
	ProfilePicture string     `json:"profile_picture" gorm:"type:text;not null"`
	CreatedAt      *time.Time `json:"created_at" gorm:"type:timestamp;not null"`
	UpdatedAt      *time.Time `json:"updated_at" gorm:"type:timestamp;null"`
}
