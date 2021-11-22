package mysql

const (
	sqlCourseTable = "courses"
)

// Map course object with a struct
type sqlCourse struct {
	ID       string `db:"id"`
	Name     string `db:"name"`
	Duration string `db:"duration"`
}
