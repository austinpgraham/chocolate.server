package user

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/securecookie"
	"golang.org/x/crypto/bcrypt"

	"github.com/austinpgraham/chocolate.server/pkg/user"
)

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func hash(val string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(val), 14)
	return string(bytes), err
}

func checkHash(target string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(target))
	return err == nil
}

func setCookie(_user *user.User, response http.ResponseWriter) error {
	value := map[string]string{
		"name": _user.Username,
	}
	encoded, err := cookieHandler.Encode("session", value)
	if err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
	err = user.SaveSession(_user.UserID, encoded)
	return err
}

func getCookie(request *http.Request) string {
	cookie, err := request.Cookie("session")
	if err != nil {
		return ""
	}
	return cookie.Value
}

func ReqAuth(w http.ResponseWriter, r *http.Request) *user.User {
	var _user user.User
	_ = json.NewDecoder(r.Body).Decode(&_user)
	containedUser := user.GetUser(user.USERNAME, _user.Username)
	cookie := getCookie(r)
	err := user.CheckSession(containedUser.UserID, cookie)
	if err != nil {
		return nil
	}
	return containedUser
}
