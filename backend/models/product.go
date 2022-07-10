package models

type Product struct {
	ID   int `gorm:"primary_key"`
	Name string
}
