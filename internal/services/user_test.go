package services

import (
	"testing"

	"github.com/kahunacohen/repo-pattern/models"
)

type InMemoryUserRepository struct{}

func (r *InMemoryUserRepository) GetOne(id int) (models.User, error) {
	return models.User{ID: 1, Email: "aharonc@matav.org.il"}, nil
}

func TestUserService(t *testing.T) {
	r := InMemoryUserRepository{}
	user, _ := r.GetOne(1)
	if user.Email != "aharonc@matav.org.il" {
		t.Fail()
	}
}
