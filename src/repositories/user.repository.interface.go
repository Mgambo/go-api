package repositories

import "github.com/mgambo/go-api/src/modals"

type UserRepositoryInterface interface {
	FindAll() []modals.User
	Save(user modals.User) modals.User
	FindById(id string) modals.User
	Update(user modals.User) modals.User
	Delete(user modals.User) bool
}
