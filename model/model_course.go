package model

type Course struct {
	Id        string `gorm:"" json:"id"`
	Name      string `gorm:"not null; size:20" json:"name"`
	Price     int    `gorm:"not null; size:50" json:"price"`
	Category  string `gorm:"not null; size:20" json:"category"`
	Thumbnail string `gorm:"not null; size:50" json:"thumbnail"`
}

type Class struct {
	UserId   string `gorm:"size:40" json:"-"`
	Users    User   `gorm:"foreignKey:UserId" json:"user"`
	CourseId string `gorm:"size:40" json:"-"`
	Courses  Course `gorm:"foreignKey:CourseId" json:"course"`
}
