package repositories

import "github.com/kahunacohen/repo-pattern/models"

type UserRepository interface {
	GetOne(id int) (models.User, error)
}
