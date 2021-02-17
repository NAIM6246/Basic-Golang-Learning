package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

type UploadsHandler struct {
}

func NewUploadsHandler() *UploadsHandler {
	return &UploadsHandler{}
}

func (h *UploadsHandler) UploadHandler(router chi.Router) {
	router.Post("/", h.uploadFile)
}

func (h *UploadsHandler) uploadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error uploading file")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Println(handler.Filename)

	//creating file
	dst, err := os.Create("./image/" + handler.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Successfully uploaded file")
}
