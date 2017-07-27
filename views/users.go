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
func PostUserHandler(m *models.Model) httprouter.Handle {
	return func (w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil || user.BirthDate == "" {
			BadRequestHandler(w, r)
			return
		}

		if ps.ByName("id") == "new" {
			users := make([]models.User, 1)
			users[0] = user
			if err := m.InsertUser(users); err != nil {
				BadRequestHandler(w, r)
				return
			}
		} else {

			_, err = strconv.ParseInt(ps.ByName("id"), 10, 32)
			if err != nil {
				NotFoundHandler(w, r)
				return
			}

			err = m.UpdateUser(user)
			if err != nil {
				NotFoundHandler(w, r)
				return
			}

			EmptyBodyHandler(w, r)
		}
	}
}
