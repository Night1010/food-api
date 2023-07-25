package models

type Food struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Name      string `json:"name"`
	Available bool   `json:"available"`
}
