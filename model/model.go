package model

import (
	"time"

	"github.com/google/uuid"
)

// Score model
type Score struct {
	ID        uuid.UUID
	CreatedAt time.Time
	User
	Game
}

type Game struct {
	Name string
}

type User struct {
	ID   uuid.UUID
	Name string
}
