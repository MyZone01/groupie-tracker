package groupietracker

import (
	utils "groupietracker/lib/utils"
	"log"
	"net/http"
)

func Index(res http.ResponseWriter, req *http.Request) {
	if utils.ValidateRequest(req, res, "/", http.MethodGet) {
		pagePath := "index"
		_data, shouldReturn := utils.GetLocations(res)
		if shouldReturn {
			return
		}
		data := string(_data)
		cells := utils.FormatMap(data)

		utils.RenderPage(pagePath, &cells, res)
		log.Println("✅ Map page get with success")
	}
}
