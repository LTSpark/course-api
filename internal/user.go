package mooc

import (
	"context"
	"course-api/utils"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

// UserID Value Object Definition and Validations
type UserID struct {
	value string
}

var ErrInvalidUserID = errors.New("invalid Course ID")

// NewUserID instantiate the VO for UserID
func NewUserID(value string) (UserID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return UserID{}, fmt.Errorf("%w: %s", ErrInvalidUserID, value)
	}

	return UserID{
		value: v.String(),
	}, nil
}

// String type converts the UserID into string.
func (id UserID) String() string {
	return id.value
}

// UserName Value Object Definition and Validations
type UserName struct {
	value string
}

var ErrEmptyUserName = errors.New("user name cannot be an empty value")
var ErrInvalidUserName = errors.New("user name cannot have more than 30 characters")

// NewUserName instantiate the VO for UserName
func NewUserName(value string) (UserName, error) {
	if value == "" {
		return UserName{}, ErrEmptyUserName
	}

	if len(value) > 30 {
		return UserName{}, ErrInvalidUserName
	}

	return UserName{
		value: value,
	}, nil
}

// String type converts the UserName into string.
func (name UserName) String() string {
	return name.value
}

// UserEmail Value Object Definition and Validations
type UserEmail struct {
	value string
}

var ErrInvalidUserEmail = errors.New("has to be a valid email")

// NewUserEmail instantiate the VO for UserName
func NewUserEmail(value string) (UserEmail, error) {

	if err := utils.IsEmailValid(value); err != nil {
		return UserEmail{}, ErrInvalidUserEmail
	}

	return UserEmail{
		value: value,
	}, nil
}

// String type converts the UserEmail into string.
func (email UserEmail) String() string {
	return email.value
}

// UserPassword Value Object Definition and Validations
type UserPassword struct {
	value string
}

var ErrEmptyUserPassword = errors.New("user password cannot be an empty value")

// NewUserPassword instantiate the VO for UserName
func NewUserPassword(value string) (UserPassword, error) {
	if len(value) < 8 {
		return UserPassword{}, ErrEmptyUserPassword
	}

	return UserPassword{
		value: value,
	}, nil
}

// String type converts the UserPassword into string.
func (password UserPassword) String() string {
	return password.value
}

type User struct {
	id       UserID
	name     UserName
	email    UserEmail
	password UserPassword
}

// NewUser creates a new course.
func NewUser(id, name, email, password string) (User, error) {

	idVO, err := NewUserID(id)
	if err != nil {
		return User{}, err
	}

	nameVO, err := NewUserName(name)
	if err != nil {
		return User{}, err
	}

	emailVO, err := NewUserEmail(email)
	if err != nil {
		return User{}, err
	}

	passwordVO, err := NewUserPassword(password)
	if err != nil {
		return User{}, err
	}

	return User{
		id:       idVO,
		name:     nameVO,
		email:    emailVO,
		password: passwordVO,
	}, nil

}

// UserRepository define the expect behavior from a user storage
type UserRepository interface {
	Save(ctx context.Context, user User) error
}

func (u User) ID() UserID {
	return u.id
}

func (u User) Name() UserName {
	return u.name
}

func (u User) Email() UserEmail {
	return u.email
}

func (u User) Password() UserPassword {
	return u.password
}
