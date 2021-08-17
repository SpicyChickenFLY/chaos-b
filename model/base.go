package model

import "time"

type User struct {
	ID      int
	Name    string
	Pwd     string
	Created time.Time
}

type Cluster struct {
	ID   int
	Host string
	Post int
	User string
	Pwd  string
}

type Test struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    int       `json:"user_id" gorm:"not null"`
	ClusterID int       `json:"cluster_id" gorm:"not null"`
	Type      int       `json:"type" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null;default now()"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null;default now()"`
	Deleted   bool      `json:"deleted" gorm:"not null;default 0"`
	Result    string    `json:"type" gorm:"not null"`
}
