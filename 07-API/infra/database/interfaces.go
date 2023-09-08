package database

import "github.com/Jhon-Henkel/go_lang_examples/tree/main/07-API/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
