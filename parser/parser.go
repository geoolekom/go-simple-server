package parser

import (
	"encoding/json"
	"path/filepath"
	"os"
	"io/ioutil"
	"strings"
	"github.com/geoolekom/go-simple-server/models"
	"fmt"
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
		fmt.Println("File", path)
		bytes, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		var data Data
		json.Unmarshal(bytes, &data)
		fmt.Println("\tInserting users")
		err = m.InsertUser(data.Users)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("\tInserting locations")
		err = m.InsertLocation(data.Locations)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("\tInserting visits")
		err = m.InsertVisit(data.Visits)
		if err != nil {
			fmt.Println(err)
		}
		return nil
	}); err != nil {
		panic(err)
	}
	fmt.Println("Data is loaded.")
}
