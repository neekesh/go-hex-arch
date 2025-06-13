package entities

import (
	"time"

	adapter_entities "github.com/thapakazi/go-hex-arch/internal/adapter/entities"
)

type User struct {
	ID string `json:"id"`
	adapter_entities.User
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
