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
	err = user.SaveSession(_user.ID, encoded)
	return err
}

func getCookie(request *http.Request) *http.Cookie {
	cookie, err := request.Cookie("session")
	if err != nil {
		return nil
	}
	return cookie
}

func ReqAuth(w http.ResponseWriter, r *http.Request) *user.User {
	cookie := getCookie(r)
	if cookie == nil {
		return nil
	}
	value := make(map[string]string)
	err := cookieHandler.Decode("session", cookie.Value, &value)
	if err != nil {
		return nil
	}
	containedUser := user.GetUser(user.USERNAME, value["name"])
	err = user.CheckSession(containedUser.ID, cookie.Value)
	if err != nil {
		return nil
	}
	return containedUser
}

func addCors(w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")
	if origin != "" {
		w.Header().Set("Access-Control-Allow_origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", 
				   "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token Authorization")
	}	
}
