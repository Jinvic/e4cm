package model

import (
	"time"

	uuidUtil "e4cm/internal/util/uuid"

	"gorm.io/gorm"
)

type Status string

const (
	StatusPending  Status = "pending"
	StatusApproved Status = "approved"
	StatusRejected Status = "rejected"
)

type SourceType string

const (
	SourceGuest  SourceType = "guest"
	SourceSystem SourceType = "system"
)

type Comment struct {
	ID        string     `gorm:"type:char(36);primaryKey" json:"id"`
	EchoID    string     `gorm:"type:char(36);not null;index" json:"echo_id"`
	UserID    *string    `gorm:"type:char(36);index" json:"user_id,omitempty"`
	Nickname  string     `gorm:"size:100;not null;index" json:"nickname"`
	Email     string     `gorm:"size:255;not null;index" json:"email"`
	Website   string     `gorm:"size:255" json:"website,omitempty"`
	Content   string     `gorm:"type:text;not null" json:"content"`
	Status    Status     `gorm:"type:varchar(20);not null;index" json:"status"`
	Hot       bool       `gorm:"not null;default:false;index" json:"hot"`
	IPHash    string     `gorm:"size:128;index" json:"-"`
	UserAgent string     `gorm:"size:512" json:"-"`
	Source    SourceType `gorm:"type:varchar(20);not null;index" json:"source"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func (c *Comment) BeforeCreate(_ *gorm.DB) error {
	if c.ID == "" {
		c.ID = uuidUtil.MustNewV7()
	}
	return nil
}

type User struct {
	ID       string `gorm:"type:char(36);primaryKey" json:"id"`
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null"        json:"-"`
	IsAdmin  bool   `gorm:"bool"                     json:"is_admin"`
	IsOwner  bool   `gorm:"bool"                     json:"is_owner"`
	Avatar   string `gorm:"size:255"                 json:"avatar"`
}
