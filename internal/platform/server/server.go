package server

import (
	"course-api/internal/creating"
	"course-api/internal/platform/server/handler/courses"
	"course-api/internal/platform/server/handler/health"
	"course-api/internal/platform/server/handler/users"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine
	//dependencyInjection
	creatingCourseService creating.CourseService
	creatingUserService   creating.UserService
}

func New(host string, port uint, creatingCourseService creating.CourseService, creatingUserService creating.UserService) Server {
	srv := Server{
		engine:                gin.New(),
		httpAddr:              fmt.Sprintf("%s:%d", host, port),
		creatingCourseService: creatingCourseService,
		creatingUserService:   creatingUserService,
	}
	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())
	s.engine.POST("/courses", courses.CreateHandler(s.creatingCourseService))
	s.engine.POST("/users", users.CreateHandler(s.creatingUserService))
}
