package course

import (
	"fmt"

	"github.com/jinzhu/gorm"

	db "github.com/austinpgraham/chocolate.server/pkg/database"
	"github.com/austinpgraham/chocolate.server/pkg/user"
)

const COURSE_ID = "id"
const COURSE_NUMBER = "course_number"
const COURSES_TABLE = "courses"

type Course struct {
	gorm.Model
	Instructor   user.User `json:"instructor" gorm:"foreignkey:InstructorID"`
	InstructorID uint      `json:"instructor_id"`
	CourseNumber string    `json:"course_number" gorm:"unique_index"`
	CourseTitle  string    `json:"course_title"`
	Description  string    `json:"description"`
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
	return &course
}
