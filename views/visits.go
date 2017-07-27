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
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		id, err := strconv.ParseInt(ps.ByName("id"), 10, 32)

		if err != nil {
			NotFoundHandler(w, r)
			return
		}

		visit, err := m.SelectVisit(int(id))
		if err != nil {
			NotFoundHandler(w, r)
			return
		}

		if err = json.NewEncoder(w).Encode(visit); err != nil {
			w.WriteHeader(500)
		}
	}
}

func PostVisitHandler(m *models.Model) httprouter.Handle {
	return func (w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		if ps.ByName("id") == "new" {
			visits := make([]models.Visit, 1)
			err := json.NewDecoder(r.Body).Decode(&visits[0])
			if err != nil || visits[0].VisitedAt == "" {
				BadRequestHandler(w, r)
				return
			}
			if err := m.InsertVisit(visits); err != nil {
				BadRequestHandler(w, r)
				return
			}
		} else {
			id, err := strconv.ParseInt(ps.ByName("id"), 10, 32)

			if err != nil {
				NotFoundHandler(w, r)
				return
			}

			_, err = m.SelectVisit(int(id))
			if err != nil {
				NotFoundHandler(w, r)
				return
			}

			EmptyBodyHandler(w, r)
		}
	}
}
