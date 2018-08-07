package course

import (
	"github.com/austinpgraham/chocolate.server/pkg/user"
)

type Course struct {
	CourseID uint `json: "id" gorm: "AUTO_INCREMENT;unique_index"`
	Instructor string `json: "instructor"`
	CourseNumber string `json: "course_number" gorm: "unique_index"`
	Description string `json: "description"`
	Students []user.User `json: "students" gorm:"-"`
}