package models

type Rating struct {
	ID       uint    `gorm:"primarykey"`
	PersonID uint    `json:"-"`
	Person   *Person `json:"person,omitempty" gorm:"foreignKey:PersonID;references:ID"`
	Stars    int     `json:"stars"`
	Comment  string  `json:"comment" gorm:"size:1000"`
}
