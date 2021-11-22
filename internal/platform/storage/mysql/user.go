package mysql

const (
	sqlUserTable = "users"
)

// Map course object with a struct
type sqlUser struct {
	ID       string `db:"id"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	Password string `db:"password"`
}
