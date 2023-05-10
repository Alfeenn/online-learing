package model

type User struct {
	Id       string `gorm:"primary_key; column:id"`
	Username string `gorm:"not null; unique; size:32"`
	Password string `gorm:"not null; size:70"`
	Name     string `gorm:"not null; size:20"`
	Age      int64  `gorm:"not null; size:20"`
	Phone    int64  `gorm:"not null; size:20"`
	Role     string `gorm:"not null; size:10"`
}
