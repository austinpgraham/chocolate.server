package user

import (
	"net/http"
	"encoding/json"

	"github.com/austinpgraham/chocolate.server/internal/user"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser user.User
	_ = json.NewDecoder(r.Body).Decode(&newUser)
	user.CreateUser(&newUser)
	w.WriteHeader(http.StatusOK)
}