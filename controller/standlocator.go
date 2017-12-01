package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"psgoweb/viewmodel"
)

type standLocator struct {
	standLocatorTemplate *template.Template
}

func (sl standLocator) registerRoutes() {
	http.HandleFunc("/stand-locator", sl.handleStandLocator)
	http.HandleFunc("/api/stands", sl.handleApiStands)
}

func (sl standLocator) handleStandLocator(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewStandLocator()
	w.Header().Add("Content-Type", "text/html")
	sl.standLocatorTemplate.Execute(w, vm)
}

func (sl standLocator) handleApiStands(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var loc struct {
		ZipCode string `json:"zipCode"`
	}
	err := dec.Decode(&loc)
	if err != nil {
		log.Println(fmt.Errorf("Error retrieving location: %v", err))
		enc := json.NewEncoder(w)
		enc.Encode([]viewmodel.StandCoordinate{})
		return
	}
	log.Println("location:", loc)
	vm := coords
	enc := json.NewEncoder(w)
	enc.Encode(vm)
}

var coords []viewmodel.StandCoordinate = []viewmodel.StandCoordinate{
	viewmodel.StandCoordinate{
		Latitude:  37.409,
		Longitude: -122.06,
		Title:     "Bobby's stand",
	},
	viewmodel.StandCoordinate{
		Latitude:  37.4092,
		Longitude: -122.061,
		Title:     "Macy's stand",
	},
	viewmodel.StandCoordinate{
		Latitude:  37.4094,
		Longitude: -122.06,
		Title:     "Juan's stand",
	},
	viewmodel.StandCoordinate{
		Latitude:  37.41,
		Longitude: -122.065,
		Title:     "Allison's stand",
	},
	viewmodel.StandCoordinate{
		Latitude:  37.415,
		Longitude: -122.07,
		Title:     "Chen's stand",
	},
	viewmodel.StandCoordinate{
		Latitude:  37.4217,
		Longitude: -122.075,
		Title:     "Matthew's stand",
	},
	viewmodel.StandCoordinate{
		Latitude:  37.4206,
		Longitude: -122.08,
		Title:     "Alice's stand",
	},
	viewmodel.StandCoordinate{
		Latitude:  37.4205,
		Longitude: -122.083,
		Title:     "Kara's stand",
	},
	viewmodel.StandCoordinate{
		Latitude:  37.414,
		Longitude: -122.09,
		Title:     "Fred's stand",
	},
	viewmodel.StandCoordinate{
		Latitude:  37.412,
		Longitude: -122.09,
		Title:     "Jake's stand",
	},
	viewmodel.StandCoordinate{
		Latitude:  37.41,
		Longitude: -122.093,
		Title:     "Wallace's stand",
	},
	viewmodel.StandCoordinate{
		Latitude:  37.416,
		Longitude: -122.095,
		Title:     "Gromit's stand",
	},
	viewmodel.StandCoordinate{
		Latitude:  37.41,
		Longitude: -122.1,
		Title:     "Kirk's stand",
	},
	viewmodel.StandCoordinate{
		Latitude:  37.41,
		Longitude: -122.102,
		Title:     "Lorelei's stand",
	},
	viewmodel.StandCoordinate{
		Latitude:  37.412,
		Longitude: -122.099,
		Title:     "Rebecca's stand",
	},
	viewmodel.StandCoordinate{
		Latitude:  37.407,
		Longitude: -122.1025,
		Title:     "Chris's stand",
	},
	viewmodel.StandCoordinate{
		Latitude:  37.423,
		Longitude: -122.1025,
		Title:     "Carson's stand",
	},
}
