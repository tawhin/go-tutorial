package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"github.com/pluralsight/webservice/models"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

// Define a method by binding a function to a type
// This func implicitly matches the Handler interface within the Handler package, go will detect this without the
// need to explicity define an implementing statement against the interface
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/users" {
		switch r.Method {
		case http.MethodGet:
			uc.getAll(w, r)
		case http.MethodPost:
			uc.post(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		matches := uc.userIDPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid or Missing User ID"))
			return
		}
		// Convert index 1 of the matches collection to an int
		id, err := strconv.Atoi(matches[1])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid User ID"))
			return
		}

		switch r.Method {
		case http.MethodGet:
			uc.get(id, w)
		case http.MethodPut:
			uc.put(id, w, r)
		case http.MethodDelete:
			uc.delete(id, w)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}

}

func (uc userController) getAll(w http.ResponseWriter, r *http.Request) {
	encodeResponseAsJson(models.GetUsers(), w)
}

func (uc userController) get(id int, w http.ResponseWriter) {
	user, err := models.GetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// As we are writing to the ResponseWriter then the header will be implicitly set
	// to StatusOK
	encodeResponseAsJson(user, w)
}

func (uc userController) post(w http.ResponseWriter, r *http.Request) {
	pendingUser, parseErr := uc.parseRequest(w, r)
	if parseErr == nil {
		newUser, err := models.AddUser(pendingUser)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		encodeResponseAsJson(newUser, w)
	}
}

func (uc userController) put(id int, w http.ResponseWriter, r *http.Request) {
	user, parseErr := uc.parseRequest(w, r)
	user.ID = id
	if parseErr == nil {
		_, err := models.UpdateUser(user)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		encodeResponseAsJson(user, w)
	}
}

func (uc userController) delete(id int, w http.ResponseWriter) {
	err := models.RemoveUser(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (uc userController) parseRequest(w http.ResponseWriter, r *http.Request) (models.User, error) {
	dec := json.NewDecoder(r.Body)
	var u models.User
	err := dec.Decode(&u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Could not parse request body as a user object"))
		return models.User{}, err
	}

	return u, nil
}

// Constructor function
func newUserController() *userController {
	// userController is local scope but returning an address.
	// go will automatically promote the local variable to a global scope when returning it's address.
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
