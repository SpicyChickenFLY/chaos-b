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
	ID        int
	UserID    int
	ClusterID int
	Type      int
	Created   time.Time
	Result    string
}
