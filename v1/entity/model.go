package entity

import "time"

type User struct {
	ID        uint16    `json:"idz" gorm:"primaryKey;AUTO_INCREMENT"`
	Username  string    `json:"username" gorm:"unique;not null"`
	Password  string    `json:"password" gorm:"not null"`
	Name      string    `json:"name" gorm:"not null; type:varchar(50)"`
	Email     string    `json:"email" gorm:"type:email"`
	CreatedAt time.Time `json:"created_at" sql:"default:CURRENT_TIMESTAMP()"`
}
