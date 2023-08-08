package model

import "time"

type User struct {
	ID          string     `gorm:"primaryKey"`
	RoleID      string     `gorm:"not null"`
	Email       string     `gorm:"not null"`
	Username    string     `gorm:"not null"`
	NoHp        string     `gorm:"not null"`
	Fullname    string     `gorm:"not null"`
	Passwd      string     `gorm:"not null"`
	PassVersion int        `gorm:"not null"`
	IsActive    bool       `gorm:"not null"`
	PhotoID     string     `gorm:"not null"`
	LastLoginDt *time.Time `gorm:"null"`
	CreateBy    string     `gorm:"not null"`
	CreateDt    time.Time  `gorm:"not null"`
	UpdateBy    string     `gorm:"not null"`
	UpdateDt    time.Time  `gorm:"not null"`
	DeleteBy    string     `gorm:"not null"`
	DeleteDt    *time.Time `gorm:"null"`
}
