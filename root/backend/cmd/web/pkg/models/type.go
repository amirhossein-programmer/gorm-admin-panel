package models

type Type struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
