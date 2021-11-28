package entity

import "time"

type User struct {
	ID        uint16    `gorm:"primaryKey;AUTO_INCREMENT"`
	Username  string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	Name      string    `gorm:"not null; type:varchar(50)"`
	Email     string    `gorm:"type:email"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP()"`
}
