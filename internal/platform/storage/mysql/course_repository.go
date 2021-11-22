package mysql

import (
	"context"
	mooc "course-api/internal"
	"database/sql"
	"fmt"

	"github.com/huandu/go-sqlbuilder"
)

// go:generate mockery --case snake --outpkg storagemocks --output platform/storage/storagemocks --name CourseRepository

// CourseRepository is a MYSQL mooc.CourseRepository implementation
type CourseRepository struct {
	db *sql.DB
}

// NewCourseRepository initializes a MySQL-based implementation of mooc.CourseRepository.
func NewCourseRepository(db *sql.DB) *CourseRepository {
	return &CourseRepository{
		db: db,
	}
}

// Save implements the mooc.CourseRepository interface and stores a course on database
func (r *CourseRepository) Save(ctx context.Context, course mooc.Course) error {

	courseSQLStruct := sqlbuilder.NewStruct(new(sqlCourse))

	// Build query using sqlbuilder ORM
	query, args := courseSQLStruct.InsertInto(sqlCourseTable, sqlCourse{
		ID:       course.ID().String(),
		Name:     course.Name().String(),
		Duration: course.Duration().String(),
	}).Build()

	// Execute query
	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist course on database: %v", err)
	}

	return nil
}
