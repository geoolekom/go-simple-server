package views

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"github.com/geoolekom/go-simple-server/models"
	"strconv"
)

func GetUserHandler(m *models.Model) httprouter.Handle {
	return func (w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		id, _ := strconv.ParseInt(ps.ByName("id"), 10, 32)
		user, err := m.SelectUser(int(id))
		if err != nil {
			fmt.Fprintf(w, "error: %s!\n", err)
		} else {
			fmt.Fprintf(w, "hello, %s!\n", user.FirstName)
		}
	}
}
