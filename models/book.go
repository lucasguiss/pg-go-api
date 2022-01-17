package models

import "time"

type Book struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	MediumPrice string    `json:"medium_price"`
	Author      string    `json:"author"`
	ImgURL      string    `json:"img_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at" gorm:"index"`
}
