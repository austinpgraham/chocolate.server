package course

import (
	"encoding/json"
	"net/http"

	"github.com/austinpgraham/chocolate.server/pkg/course"
	"github.com/austinpgraham/chocolate.server/internal/app/user"
	er "github.com/austinpgraham/chocolate.server/internal/app/error"
)

func CreateCourse(w http.ResponseWriter, r *http.Request) {
	containedUser := user.ReqAuth(w, r)
	if containedUser == nil {
		err := er.Error{Code: "LoginRequired.", Message: "Login Required."}
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(err)
		return
	}
	var newCourse course.Course
	_ = json.NewDecoder(r.Body).Decode(&newCourse)
	found := course.GetCourse(course.COURSE_NUMBER, newCourse.CourseNumber)
	if found != nil {
		err := er.Error{Code: "CourseExists.", Message: "Course exists."}
		w.WriteHeader(http.StatusFound)
		json.NewEncoder(w).Encode(err)
		return
	}
	newCourse.Instructor = containedUser
	course.CreateCourse(&newCourse)
	w.WriteHeader(http.StatusCreated)
}
