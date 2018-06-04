package user

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"

	"github.com/austinpgraham/chocolate.server/internal/user"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser user.User
	_ = json.NewDecoder(r.Body).Decode(&newUser)
	potUser := user.GetUser(newUser.Username)
	if potUser != nil {
		w.WriteHeader(http.StatusFound)
		return
	}
	user.CreateUser(&newUser)
	w.WriteHeader(http.StatusOK)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	_user := user.GetUser(username)
	if _user == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	} 
	json.NewEncoder(w).Encode(_user)
}