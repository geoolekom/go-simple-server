package views

import (
	"github.com/geoolekom/go-simple-server/models"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"encoding/json"
)

func GetVisitHandler(m *models.Model) httprouter.Handle {
	return func (w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		id, err := strconv.ParseInt(ps.ByName("id"), 10, 32)
		if err != nil {
			w.WriteHeader(404)
			return
		}

		visit, err := m.SelectVisit(int(id))
		if err != nil {
			w.WriteHeader(404)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if err = json.NewEncoder(w).Encode(visit); err != nil {
			w.WriteHeader(500)
		}
	}
}
