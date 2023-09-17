package domain

type User struct {
	ID       int    `json:"-"`
	Name     string `gorm:"not null" binding:"required"`
	Surname  string `gorm:"not null" binding:"required"`
	Username string `gorm:"unique;not null" binding:"required"`
	Password string `gorm:"not null" binding:"required"`
}
