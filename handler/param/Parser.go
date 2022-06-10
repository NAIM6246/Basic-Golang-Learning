package param

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

//Int 	:
func Int(r *http.Request, p string) int {
	idString := chi.URLParam(r, p)
	/*
		//	checking the id
		fmt.Println(p, 2)
		fmt.Println(idString)
	*/
	if len(idString) == 0 {
		//	fmt.Println("Entered")
		id, err := strconv.Atoi(idString)
		/*
			//cheching id should be commented
			fmt.Println(id, 3)
		*/
		if err != nil {
			//fmt.Println(err, 4)

			return 0
		}
		return id
	}
	return 0
}

//UInt 	:
func UInt(r *http.Request, p string) uint {
	return uint(Int(r, p))
}
