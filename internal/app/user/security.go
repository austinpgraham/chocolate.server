package user

import (
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
	value := map[string]string {
		"name": _user.Username,
	}
	encoded, err := cookieHandler.Encode("session", value)
	if err == nil {
		cookie := &http.Cookie{
			Name: "session",
			Value: encoded,
			Path: "/",
		}
		http.SetCookie(response, cookie)
	}
	err = user.SaveSession(_user.UserID, encoded)
	return err
}