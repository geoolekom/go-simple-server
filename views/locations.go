package views

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
)

func GetLocation(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Location %s!\n", ps.ByName("id"))
}
