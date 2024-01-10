package models

type Developer struct {
	ID        int    `json:"id" gorm:"primaryKey;autoIncrement"`
	FirstName string `json:"first_name" gorm:"column:first_name;not null"`
	LastName  string `json:"last_name" gorm:"column:last_name;not null"`
	Email     string `json:"email" gorm:"column:email;unique;not null"`
}
