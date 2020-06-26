package register

import "time"

type User struct {
	ID        uint64
	UniqueID  string
	Firstname string
	Lastname  string
	Address   string
	Email     string
	CreatedAt time.Time
	UpdateAt  time.Time
	DeletedAt time.Time
}
