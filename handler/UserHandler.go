package handler

import (
	"Golang/handler/param"
	"Golang/services"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

//Handler struct :
type UserHandler struct {
}

/*
type response struct {
	Name string `json:"name"`
}
*/

//NewUserHandler is the constructor of UserHandler struct
func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) Handle(rout chi.Router) {
	rout.Route("/{id}", func(router chi.Router) {
		//Url : users/get/id
		router.Get("/id", h.getUserByID)
	})
	rout.Get("/get", h.getUser)
	rout.Post("/post", h.createUser)
	rout.Put("/update", h.updateUser)
	rout.Delete("/delete", h.deleteUser)
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
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write([]byte(`{"essage" : "data created"}`))
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
