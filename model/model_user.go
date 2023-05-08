package model

type User struct {
	Id        string `gorm:"primary_key; column:id"`
	Email     string `gorm:"not null; unique; size:32"`
	Password  string `gorm:"not null; size:20"`
	Role      string `gorm:"not null; size:10"`
	CreatedAt string `gorm:"column:created_at"`
	UpdatedAt string `gorm:"column:updated_at"`
}
