package course

import (
	"github.com/austinpgraham/chocolate.server/pkg/user"

	db "github.com/austinpgraham/chocolate.server/pkg/database"
)

const COURSES_TABLE = "courses"

type Course struct {
	CourseID uint `json: "id" gorm: "AUTO_INCREMENT;unique_index"`
	Instructor string `json: "instructor"`
	CourseNumber string `json: "course_number" gorm: "unique_index"`
	Description string `json: "description"`
	Students []user.User `json: "students" gorm:"-"`
}

func checkTable() {
	db, _ := db.GetConnection()
	defer db.Close()
	if !db.HasTable(COURSES_TABLE) {
		db.CreateTable(&Course{})
	}
}

func CreateCourse(course *Course) error {
	checkTable()
	db, _ := db.GetConnection()
	defer db.Close()
	err := db.Create(course).Error
	return err
}

func GetCourse(att string, val string) *Course {
	checkTable()
	db, _ := db.GetConnection()
	defer db.Close()
	var course Course
	if db.First(&course, fmt.Sprintf("%v = ?", att), val).Error != nil {
		return nil
	}
	course.Students = getStudentRoster(course.CourseID)
	return &user
}
