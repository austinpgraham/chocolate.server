package neighborhood

import (
	"encoding/json"
	"net/http"

	er "github.com/austinpgraham/chocolate.server/internal/app/error"
	"github.com/austinpgraham/chocolate.server/internal/app/user"
	"github.com/austinpgraham/chocolate.server/pkg/neighborhood"
)

func CreateNeighborhood(w http.ResponseWriter, r *http.Request) {
	containedUser := user.ReqAuth(w, r)
	if containedUser == nil {
		err := er.Error{Code: "LoginRequired.", Message: "Login Required."}
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(err)
		return
	}
	var nb neighborhood.Neighborhood
	_ = json.NewDecoder(r.Body).Decode(&nb)
	found := neighborhood.GetNeighborhood(neighborhood.NAME, nb.Name)
	if found != nil {
		err := er.Error{Code: "NeihborhoodExists.", Message: "Neighborhood exists."}
		w.WriteHeader(http.StatusFound)
		json.NewEncoder(w).Encode(err)
		return
	}
	nb.Admin = *containedUser
	neighborhood.CreateNeighborhood(&nb)
	w.WriteHeader(http.StatusCreated)
}
