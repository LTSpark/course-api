package creating

import (
	"context"
	mooc "course-api/internal"
	"course-api/utils"
)

// UserService is the default UserService interface
// implementation returned by creating.NewUserService.
type UserService struct {
	UserRepository mooc.UserRepository
}

// NewUserService returns the default Service interface implementation.
func NewUserService(UserRepository mooc.UserRepository) UserService {
	return UserService{
		UserRepository: UserRepository,
	}
}

// CreateUser implements the creating.UserService interface.
func (s UserService) CreateUser(ctx context.Context, id, name, email, password string) error {

	password, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	User, err := mooc.NewUser(id, name, email, password)
	if err != nil {
		return err
	}
	return s.UserRepository.Save(ctx, User)
}
