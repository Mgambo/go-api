package repositories

import (
	"github.com/mgambo/go-api/api/models"
	internal_repositories "github.com/mgambo/go-api/internal/repositories"
	"gorm.io/gorm"
)

type UserRepository interface {
	Base() *internal_repositories.BaseRepository[models.User]
}

type userRepository struct {
	base *internal_repositories.BaseRepository[models.User]
}

func (r *userRepository) Base() *internal_repositories.BaseRepository[models.User] {
	return r.base
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		base: internal_repositories.NewBaseRepository[models.User](db),
	}
}
