package bootstrap

import (
	"course-api/internal/platform/server"
	"course-api/internal/platform/storage/mysql"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	host   = "localhost"
	port   = 8080
	dbUser = "root"
	dbPass = "ubuntuleo"
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "courseapi"
)

func Run() error {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	courseRepository := mysql.NewCourseRepository(db)

	srv := server.New(host, port, courseRepository)
	return srv.Run()
}
