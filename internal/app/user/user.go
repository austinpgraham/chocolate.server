package user

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	er "github.com/austinpgraham/chocolate.server/internal/app/error"
	puser "github.com/austinpgraham/chocolate.server/pkg/user"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser puser.User
	_ = json.NewDecoder(r.Body).Decode(&newUser)
	newUser.Password, _ = hash(newUser.Password)
	found := puser.GetUser("username", newUser.Username)
	if found != nil {
		err := er.Error{Code: "UsernameExists", Message: "Username exists."}
		w.WriteHeader(http.StatusFound)
		json.NewEncoder(w).Encode(err)
		return
	}
	found = puser.GetUser("email", newUser.Email)
	if found != nil {
		err := er.Error{Code: "EmailExists", Message: "Email exists."}
		w.WriteHeader(http.StatusFound)
		json.NewEncoder(w).Encode(err)
		return
	}
	found = puser.GetUser("f_by_f", newUser.FByF)
	if found != nil {
		err := er.Error{Code: "FByFExsists", Message: "FByF exists."}
		w.WriteHeader(http.StatusFound)
		json.NewEncoder(w).Encode(err)
		return
	}
	puser.CreateUser(&newUser)
	addCors(w, r)
	w.WriteHeader(http.StatusOK)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	authUser := ReqAuth(w, r)
	if authUser == nil {
		err := er.Error{Code: "LoginRequired.", Message: "Login Required."}
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(err)
		return
	}
	username := mux.Vars(r)["username"]
	_user := puser.GetUser(puser.USERNAME, username)
	if _user == nil || _user.Username != username {
		err := er.Error{Code: "UserNotFound", Message: "User not found."}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}
	_user.Password = ""
	addCors(w, r)
	json.NewEncoder(w).Encode(_user)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var _user puser.User
	_ = json.NewDecoder(r.Body).Decode(&_user)
	containedUser := puser.GetUser(puser.USERNAME, _user.Username)
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
	addCors(w, r)
	w.WriteHeader(http.StatusOK)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	containedUser := ReqAuth(w, r)
	if containedUser == nil {
		err := er.Error{Code: "LoginRequired.", Message: "Login Required."}
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(err)
		return
	}
	puser.RemoveSession(containedUser.UserID)
	addCors(w, r)
	w.WriteHeader(http.StatusOK)
}
