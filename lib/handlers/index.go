package groupietracker

import (
	utils "groupietracker/lib/utils"
	"log"
	"net/http"
)

func Index(res http.ResponseWriter, req *http.Request) {
	if utils.ValidateRequest(req, res, "/", http.MethodGet) {
		pagePath := "index"
		res.WriteHeader(http.StatusOK)
		url := "https://groupietrackers.herokuapp.com/api/locations"
		_data, err := utils.GetAPI(url)
		if err != nil {
			utils.RenderPage("500", nil, res)
			return
		}
		data := string(_data)
		cells := utils.FormatMap(data)

		utils.RenderPage(pagePath, &cells, res)
		log.Println("✅ Map page get with success")
	}
}
