package models

import "time"

type Movie struct {
	ID        int64     `json:"id" gorm:"primary_key;AUTO_INCREMENT;column:ID"`
	Title     string    `json:"title" gorm:"column:Title"`
	Completed bool      `json:"completed" gorm:"column:Completed"`
	CreatedAt time.Time `json:"created_at" gorm:"column:CreatedAt"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:UpdatedAt"`
}

func (Movie) TableName() string {
	return "Movie"
}
