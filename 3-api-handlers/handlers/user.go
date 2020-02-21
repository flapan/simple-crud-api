package handlers

import (
	"net/http"
	"strconv"
)

/* // User sfsifjiosfj
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
} */

//UserController sfjfj posjf
type UserController struct {
	users []User
}

// Users sfijspfj
func (uc *UserController) Users(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		//Respond(w, uc.users, http.StatusOK)
	case http.MethodPost:

		var user User

		user.Name = "michael" + strconv.Itoa(len(uc.users))
		//user.Email = "michael@person.com"

		uc.users = append(uc.users, user)

		//Respond(w, user, http.StatusOK)
	}
}
