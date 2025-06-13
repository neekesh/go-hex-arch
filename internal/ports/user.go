package ports

import (
	"context"

	adapter_entities "github.com/thapakazi/go-hex-arch/internal/adapter/entities"
	"github.com/thapakazi/go-hex-arch/internal/core/entities"
)

// outgoing hex arch ports
// communication between core and adapter
// as this is need by the core to interact with the database
// we need to define the interface here
type UserRepository interface {
	CreateUser(ctx context.Context, user *entities.User) error
	GetUser(ctx context.Context, id int64) (*entities.User, error)
	UpdateUser(ctx context.Context, user *entities.User) error
	DeleteUser(ctx context.Context, id int64) error
	GetAllUsers(ctx context.Context, params adapter_entities.QueryParams) ([]entities.User, error)
}

// Incoming ports
// called by the service from the http
// calls the service
type UserService interface {
	CreateUser(ctx context.Context, user *adapter_entities.User) error
	GetUser(ctx context.Context, id int64) (*entities.User, error)
	UpdateUser(ctx context.Context, user *entities.User) error
	DeleteUser(ctx context.Context, id int64) error
	GetAllUsers(ctx context.Context, params adapter_entities.QueryParams) ([]entities.User, error)
}
