package user

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"

	"github.com/austinpgraham/chocolate.server/pkg/user"
	er "github.com/austinpgraham/chocolate.server/internal/app/error"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser user.User
	_ = json.NewDecoder(r.Body).Decode(&newUser)
	found := user.GetUser("username", newUser.Username)
	if found != nil {
		err := er.Error{Code: "UsernameExists", Message: "Username exists."}
		w.WriteHeader(http.StatusFound)
		json.NewEncoder(w).Encode(err)
		return
	}
	found = user.GetUser("email", newUser.Email)
	if found != nil {
		err := er.Error{Code: "EmailExists", Message: "Email exists."}
		w.WriteHeader(http.StatusFound)
		json.NewEncoder(w).Encode(err)
		return
	}
	found = user.GetUser("f_by_f", newUser.FByF)
	if found != nil {
		err := er.Error{Code: "FByFExsists", Message: "FByF exists."}
		w.WriteHeader(http.StatusFound)
		json.NewEncoder(w).Encode(err)
		return
	}
	user.CreateUser(&newUser)
	w.WriteHeader(http.StatusOK)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	_user := user.GetUser("username", username)
	if _user.Username != username {
		err := er.Error{Code: "UserNotFound", Message: "User not found."}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	} 
	json.NewEncoder(w).Encode(_user)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var _user user.User
	_ = json.NewDecoder(r.Body).Decode(&_user)
	containedUser := user.GetUser(user.USERNAME, _user.Username)
	if containedUser == nil {
		err := er.Error{Code: "UserNotFound.", Message: "User not found."}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}
	if !checkHash(_user.Password, containedUser.Password) {
		err := er.Error{Code: "PasswordIncorrect.", Message: "Password incorrect."}
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(err)
		return
	}
	err := setCookie(containedUser, w) 
	if err != nil {
		logerr := er.Error{Code: "CannotLogin.", Message: "Error logging in."}
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(logerr)
		return
	}
	w.WriteHeader(http.StatusOK)
}
