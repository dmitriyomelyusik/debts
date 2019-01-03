// package rest - implements REST API.

package rest

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/dmitriyomelyusik/debts/backend/domain"

	"github.com/dmitriyomelyusik/debts/backend/service"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Controller bla-bla
type Controller struct {
	service *service.Service
}

// NewController bla-bla
func NewController(service *service.Service) Controller {
	if service == nil {
		log.Fatal("service can't be nil")
	}
	return Controller{service: service}
}

func (c Controller) addUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondWithError(w, http.StatusOK, err.Error())
		return
	}
	var user domain.User

	err = json.Unmarshal(data, &user)
	if err != nil {
		respondWithError(w, http.StatusOK, err.Error())
		return
	}

	user, err = c.service.AddUser(user)
	if err != nil {
		respondWithError(w, http.StatusOK, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, user)
}

func (c Controller) editUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondWithError(w, http.StatusOK, err.Error())
		return
	}
	var user domain.User

	err = json.Unmarshal(data, &user)
	if err != nil {
		respondWithError(w, http.StatusOK, err.Error())
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusOK, err.Error())
		return
	}

	err = c.service.UpdateUser(id, user)
	if err != nil {
		respondWithError(w, http.StatusOK, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, nil)
}

func (c Controller) deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusOK, err.Error())
		return
	}
	err = c.service.DeleteUser(id)
	if err != nil {
		respondWithError(w, http.StatusOK, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, nil)
}

func (c Controller) getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusOK, err.Error())
		return
	}
	user, err := c.service.GetUser(id)
	if err != nil {
		respondWithError(w, http.StatusOK, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, user)
}

func (c Controller) getUsers(w http.ResponseWriter, r *http.Request) {
	users, err := c.service.GetUsers()
	if err != nil {
		respondWithError(w, http.StatusOK, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, users)
}

func (c Controller) addDebt(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondWithError(w, http.StatusOK, err.Error())
		return
	}
	var debt domain.Debt

	err = json.Unmarshal(data, &debt)
	if err != nil {
		respondWithError(w, http.StatusOK, err.Error())
		return
	}

	debt, err = c.service.AddDebt(debt)
	if err != nil {
		respondWithError(w, http.StatusOK, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, debt)
}

func (c Controller) editDebt(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondWithError(w, http.StatusOK, err.Error())
		return
	}
	var debt domain.Debt

	err = json.Unmarshal(data, &debt)
	if err != nil {
		respondWithError(w, http.StatusOK, err.Error())
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusOK, err.Error())
		return
	}

	err = c.service.UpdateDebt(id, debt)
	if err != nil {
		respondWithError(w, http.StatusOK, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, nil)
}

func (c Controller) deleteDebt(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusOK, err.Error())
		return
	}
	err = c.service.DeleteUser(id)
	if err != nil {
		respondWithError(w, http.StatusOK, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, nil)
}

func (c Controller) getDebt(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusOK, err.Error())
		return
	}
	debt, err := c.service.GetDebt(id)
	if err != nil {
		respondWithError(w, http.StatusOK, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, debt)
}

func (c Controller) getDebts(w http.ResponseWriter, r *http.Request) {
	debts, err := c.service.GetDebts()
	if err != nil {
		respondWithError(w, http.StatusOK, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, debts)
}

// NewRouter returns new instance of server router
func NewRouter(ctrl *Controller) *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/users", ctrl.getUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", ctrl.getUser).Methods(http.MethodGet)
	router.HandleFunc("/users", ctrl.addUser).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/users/{id}", ctrl.editUser).Methods(http.MethodPatch, http.MethodOptions)
	router.HandleFunc("/users/{id}", ctrl.deleteUser).Methods(http.MethodDelete, http.MethodOptions)

	router.HandleFunc("/debts", ctrl.getDebts).Methods(http.MethodGet)
	router.HandleFunc("/debts/{id}", ctrl.getDebt).Methods(http.MethodGet)
	router.HandleFunc("/debts", ctrl.addDebt).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/debts/{id}", ctrl.editDebt).Methods(http.MethodPatch, http.MethodOptions)
	router.HandleFunc("/debts/{id}", ctrl.deleteDebt).Methods(http.MethodDelete, http.MethodOptions)

	originsOK := handlers.AllowedOrigins([]string{"*"})
	headersOK := handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Content-Length",
		"Accept-Encoding", "X-CSRF-Token", "Authorization"})
	methodsOK := handlers.AllowedMethods([]string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"})
	router.Use(handlers.CORS(originsOK, headersOK, methodsOK))

	return router
}

// respondWithError - responds with an error.
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

// respondWithJSON - responds with a json.
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Print(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(response)
	if err != nil {
		log.Print(err)
	}
}
