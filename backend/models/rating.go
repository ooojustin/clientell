package models

type Rating struct {
	ID       uint    `gorm:"primarykey"`
	PersonID uint    `json:"-"`
	Person   Person  `json:"person" gorm:"foreignKey:PersonID;references:ID"`
	Stars    float64 `json:"stars"`
}
