package views

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/geoolekom/go-simple-server/models"
	"strconv"
	"encoding/json"
)

func GetUserHandler(m *models.Model) httprouter.Handle {
	return func (w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		id, err := strconv.ParseInt(ps.ByName("id"), 10, 32)

		if err != nil {
			NotFoundHandler(w, r)
			return
		}

		user, err := m.SelectUser(int(id))
		if err != nil {
			NotFoundHandler(w, r)
			return
		}
		if err = json.NewEncoder(w).Encode(user); err != nil {
			w.WriteHeader(500)
		}
	}
}
