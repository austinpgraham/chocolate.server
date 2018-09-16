package neighborhood

import (
	"encoding/json"
	"fmt"
	"net/http"

	er "github.com/austinpgraham/chocolate.server/internal/app/error"
	"github.com/austinpgraham/chocolate.server/internal/app/user"
	"github.com/austinpgraham/chocolate.server/pkg/neighborhood"
	muser "github.com/austinpgraham/chocolate.server/pkg/user"
	"github.com/gorilla/mux"
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
	nb.Password, _ = hash(nb.Password)
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

func GetNeighborhood(w http.ResponseWriter, r *http.Request) {
	authUser := user.ReqAuth(w, r)
	if authUser == nil {
		err := er.Error{Code: "LoginRequired.", Message: "Login Required."}
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(err)
		return
	}
	nb := mux.Vars(r)["neighborhood"]
	neigh := neighborhood.GetNeighborhood(neighborhood.NAME, nb)
	if neigh == nil {
		err := er.Error{Code: "NeighborhoodNotFound", Message: "Neighborhood not found."}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}
	neigh.Password = ""
	neigh.Admin = *muser.GetUser("ID", fmt.Sprintf("%d", authUser.Model.ID))
	neigh.Admin.Password = ""
	json.NewEncoder(w).Encode(neigh)
}

func GetOwnedNeighborhoods(w http.ResponseWriter, r *http.Request) {
	authUser := user.ReqAuth(w, r)
	if authUser == nil {
		err := er.Error{Code: "LoginRequired.", Message: "Login Required."}
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(err)
		return
	}
	owned := neighborhood.GetOwnedNeighborhoods(authUser)
	if owned == nil {
		err := er.Error{Code: "GetError.", Message: "Get Error."}
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(owned)
}
