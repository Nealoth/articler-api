package domain

import "time"

type UserEntity struct {
	Version        uint64    `db:"version"`
	Created        time.Time `db:"created"`
	Updated        time.Time `db:"updated"`
	ID             uint64    `db:"id"`
	Email          string    `db:"email"`
	HashedPassword string    `db:"hashed_password"`
}

type UserCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserProfile struct {
	ID    uint64 `json:"-"`
	Email string `json:"email"`
}
