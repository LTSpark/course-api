package mysql

import (
	"context"
	mooc "course-api/internal"
	"database/sql"
	"fmt"

	"github.com/huandu/go-sqlbuilder"
)

// go:generate mockery --case snake --outpkg storagemocks --output platform/storage/storagemocks --name UserRepository

// UserRepository is a MYSQL mooc.UserRepository implementation
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository initializes a MySQL-based implementation of mooc.UserRepository.
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// Save implements the mooc.UserRepository interface and stores a course on database
func (r *UserRepository) Save(ctx context.Context, user mooc.User) error {

	courseSQLStruct := sqlbuilder.NewStruct(new(sqlUser))

	// Build query using sqlbuilder ORM
	query, args := courseSQLStruct.InsertInto(sqlUserTable, sqlUser{
		ID:       user.ID().String(),
		Name:     user.Name().String(),
		Email:    user.Email().String(),
		Password: user.Password().String(),
	}).Build()

	// Execute query
	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist course on database: %v", err)
	}

	return nil
}
