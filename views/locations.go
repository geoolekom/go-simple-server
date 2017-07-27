package views

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/geoolekom/go-simple-server/models"
	"strconv"
	"encoding/json"
)

func GetLocationHandler(m *models.Model) httprouter.Handle {
	return func (w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		id, err := strconv.ParseInt(ps.ByName("id"), 10, 32)

		if err != nil {
			NotFoundHandler(w, r)
			return
		}

		location, err := m.SelectLocation(int(id))
		if err != nil {
			NotFoundHandler(w, r)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if err = json.NewEncoder(w).Encode(location); err != nil {
			w.WriteHeader(500)
		}
	}
}

func PostLocationHandler(m *models.Model) httprouter.Handle {
	return func (w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		if ps.ByName("id") == "new" {
			locations := make([]models.Location, 1)
			err := json.NewDecoder(r.Body).Decode(&locations[0])
			if err != nil {
				BadRequestHandler(w, r)
				return
			}
			if err := m.InsertLocation(locations); err != nil {
				BadRequestHandler(w, r)
				return
			}
		} else {
			id, err := strconv.ParseInt(ps.ByName("id"), 10, 32)

			if err != nil {
				NotFoundHandler(w, r)
				return
			}

			_, err = m.SelectLocation(int(id))
			if err != nil {
				NotFoundHandler(w, r)
				return
			}

			EmptyBodyHandler(w, r)
		}
	}
}
