package groupie_tracker

import (
	"fmt"
	models "groupie_tracker/lib/models"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strings"
	"text/template"
)

func ValidateRequest(req *http.Request, res http.ResponseWriter, url, method string) bool {
	if strings.Contains(url, "*") && path.Dir(url) == path.Dir(req.URL.Path) {
		return true
	}

	if req.URL.Path != url {
		res.WriteHeader(http.StatusNotFound)
		RenderPage("404", nil, res)
		log.Println("404 ❌ - Page not found ", req.URL)
		return false
	}

	if req.Method != method {
		res.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(res, "%s", "Error - Method not allowed")
		log.Println("405 ❌ - Method not allowed")
		return false
	}
	return true
}

func GetAPI(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		log.Println("🚨" + err.Error())
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("🚨" + err.Error())
	}
	defer response.Body.Close()
	return data
}

func RenderPage(pagePath string, data any, res http.ResponseWriter) {
	files := []string{"templates/base.html", "templates/" + pagePath + ".html"}
	tpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Println("🚨" + err.Error())
	} else {
		tpl.Execute(res, data)
	}
}

func FormatLocations(_relation models.RelationModel) []models.Location {
	locations := []models.Location{}
	for _locations, _dates := range _relation.DatesLocations {
		__locations := strings.Split(_locations, "-")
		var l models.Location
		if len(__locations) == 2 {
			l.City = strings.Title(strings.ReplaceAll(__locations[0], "_", " "))
			l.Country = strings.Title(strings.ReplaceAll(__locations[1], "_", " "))
		}
		l.Dates = _dates
		locations = append(locations, l)
	}
	return locations
}
