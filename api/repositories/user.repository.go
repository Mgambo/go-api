package repositories

import (
	"github.com/mgambo/go-api/api/models"
	shared_repositories "github.com/mgambo/go-api/internal/repositories"
	"gorm.io/gorm"
)

type UserRepository interface {
	Base() *shared_repositories.BaseRepository[models.User]
}

type userRepository struct {
	base *shared_repositories.BaseRepository[models.User]
}

func (r *userRepository) Base() *shared_repositories.BaseRepository[models.User] {
	return r.base
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		base: shared_repositories.NewBaseRepository[models.User](db),
	}
}
