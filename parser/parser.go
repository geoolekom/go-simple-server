package parser

import (
	"encoding/json"
	"path/filepath"
	"os"
	"io/ioutil"
	"strings"
	"github.com/geoolekom/go-simple-server/models"
)

type Data struct {
	Users []models.User `json:"users"`
	Visits []models.Visit `json:"visits"`
	Locations []models.Location `json:"locations"`
}

func LoadData(m *models.Model) {
	dataDir := "/media/data/"
	if err := filepath.Walk(dataDir, func(path string, info os.FileInfo, err error) error {
		if !strings.HasSuffix(path, ".json") {
			return nil
		}
		bytes, err := ioutil.ReadFile(path)
		var data Data
		json.Unmarshal(bytes, &data)

		for _, user := range data.Users {
			m.InsertUser(&user)
		}
		for _, location := range data.Locations {
			m.InsertLocation(&location)
		}
		for _, visit := range data.Visits {
			m.InsertVisit(&visit)
		}
		return err
	}); err != nil {
		panic(err)
	}
}
