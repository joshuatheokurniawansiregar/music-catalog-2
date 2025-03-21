package memberships

import (
	"github.com/joshuatheokurniawansiregar/music_catalog_2/internal/configs"
	"github.com/joshuatheokurniawansiregar/music_catalog_2/internal/models/memberships"
)

//go:generate mockgen -source=service.go -destination=service_mocken_test.go -package=memberships
type repositoryInterface interface{
	CreateUser(model memberships.User) error
	GetUser(email string, username string, id uint)(*memberships.User, error)
}

type Service struct {
	cfg *configs.Config
	repositoryInterface repositoryInterface
}

func NewService(cfg *configs.Config, repositoryInterface repositoryInterface)*Service{
	return &Service{
		cfg: cfg,
		repositoryInterface: repositoryInterface,
	}
}