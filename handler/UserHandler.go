package handler

import (
	"Golang/handler/param"
	"Golang/models"
	"Golang/services"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

//Handler struct :
type UserHandler struct {
	userService services.UserService
}

//NewUserHandler is the constructor of UserHandler struct
func NewUserHandler() *UserHandler {
	return &UserHandler{
		userService: services.UserService{},
	}
}

func (h *UserHandler) Handle(rout chi.Router) {

	rout.Route("/{id}", func(router chi.Router) {
		//Url : users/get/id
		router.Get("/id", h.getUserByID)
	})
	rout.Get("/", h.getUser)
	rout.Post("/", h.createUser)
	rout.Put("/", h.updateUser)
	rout.Delete("/", h.deleteUser)
}

func (h *UserHandler) getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Cotent-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte(`{"message" : "data fetch"}`))
}

func (h *UserHandler) getUserByID(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("getuserbyid 1")
	id := param.UInt(r, "id")
	userService := services.NewUserService()
	d, e := userService.GetUserByID(id)
	if e != nil {
		w.Header().Add("Content-type", "application/json")
		w.WriteHeader(200)
		w.Write([]byte(`{"name" :"Not found"}`))
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte("Name : " + d.NAME))
	//	resp := response{Name: d}
	//	byteArray, err := json.Marshal(resp)
	/*	if err == nil {
			w.Write([]byte(byteArray))
		}
	*/

	fmt.Println(d.NAME)
	fmt.Println(d.ID)

}

func (h *UserHandler) createUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}
	d, e := h.userService.CreateUser(&user)
	if e != nil {
		panic(e)
	} else {
		fmt.Println(d)
	}
	fmt.Fprintf(w, "User create")
}

func (h *UserHandler) updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Cotent-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte(`{"message" : "data updated"}`))
}

func (h *UserHandler) deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Cotent-Type", "application/json")
	w.WriteHeader(204)
}
