package mooc

import (
	"context"
	"course-api/utils"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

var ErrInvalidCourseID = errors.New("invalid Course ID")

// CourseID represents the course unique identifier.
type CourseID struct {
	value string
}

// NewCourseID instantiate the VO for CourseID
func NewCourseID(value string) (CourseID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return CourseID{}, fmt.Errorf("%w: %s", ErrInvalidCourseID, value)
	}

	return CourseID{
		value: v.String(),
	}, nil
}

// String type converts the CourseID into string.
func (id CourseID) String() string {
	return id.value
}

var ErrEmptyCourseName = errors.New("the field Course Name can not be empty")

// CourseName represents the course name.
type CourseName struct {
	value string
}

// NewCourseName instantiate VO for CourseName
func NewCourseName(value string) (CourseName, error) {
	if value == "" {
		return CourseName{}, ErrEmptyCourseName
	}

	return CourseName{
		value: value,
	}, nil
}

// String type converts the CourseName into string.
func (name CourseName) String() string {
	return name.value
}

// CourseDuration error types
var ErrEmptyDuration = errors.New("the field Duration can not be empty")
var ErrNotEnoughDurationArgs = errors.New("duration has to contain a number and the time unit")
var ErrInvalidArg = errors.New("bad request on argument for duration")

// CourseDuration constants
const (
	ArgsValue = 2
	Magnitude = 0
	TimeUnit  = 1
)

var ValidTimeUnits = [...]string{"months", "days", "weeks", "years"}

// CourseDuration represents the course duration.
type CourseDuration struct {
	value string
}

func NewCourseDuration(value string) (CourseDuration, error) {
	if value == "" {
		return CourseDuration{}, ErrEmptyDuration
	}

	durationArgs := strings.Split(value, " ")
	if len(durationArgs) != ArgsValue {
		return CourseDuration{}, ErrNotEnoughDurationArgs
	}

	if _, err := strconv.Atoi(durationArgs[Magnitude]); err != nil {
		return CourseDuration{}, ErrInvalidArg
	}

	err := utils.Contains(ValidTimeUnits[:], durationArgs[TimeUnit])
	if err != nil {
		return CourseDuration{}, ErrInvalidArg
	}

	return CourseDuration{
		value: value,
	}, nil
}

// String type converts the CourseDuration into string.
func (duration CourseDuration) String() string {
	return duration.value
}

// Data structure which represents a course
type Course struct {
	id       CourseID
	name     CourseName
	duration CourseDuration
}

// NewCourse creates a new course.
func NewCourse(id, name, duration string) (Course, error) {

	idVO, err := NewCourseID(id)
	if err != nil {
		return Course{}, err
	}

	nameVO, err := NewCourseName(name)
	if err != nil {
		return Course{}, err
	}

	durationVO, err := NewCourseDuration(duration)
	if err != nil {
		return Course{}, err
	}

	return Course{
		id:       idVO,
		name:     nameVO,
		duration: durationVO,
	}, nil

}

// CourseRepository define the expected behavior from a course storage
type CourseRepository interface {
	Save(ctx context.Context, course Course) error
}

// ID returns the course unique identifier.
func (c Course) ID() CourseID {
	return c.id
}

// Name returns the course name.
func (c Course) Name() CourseName {
	return c.name
}

// Duration returns the course duration.
func (c Course) Duration() CourseDuration {
	return c.duration
}
