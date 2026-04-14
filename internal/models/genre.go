package models

type Genre struct {
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"unique"`
}
