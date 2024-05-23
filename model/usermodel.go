package model

import "time"

type User struct {
	ID        string    `json:"id" gorm:"primaryKey;type:varchar(55);index"`
	Username  string    `json:"username" gorm:"type:varchar(55);not null"`
	Email     string    `json:"email" gorm:"type:varchar(55);not null;unique"`
	Password  string    `json:"password" gorm:"type:varchar(255);not null"`
	Address   string    `json:"address" gorm:"type:text"`
	Photo     string    `json:"photo" gorm:"type:varchar(255);default:'https://www.weact.org/wp-content/uploads/2016/10/Blank-profile.png'"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp;default:current_timestamp;autoUpdateTime"`
	IsDeleted bool      `json:"is_deleted" gorm:"type:boolean;default:false"`
}
