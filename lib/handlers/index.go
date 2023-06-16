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
		utils.RenderPage(pagePath, nil, res)
		log.Println("✅ Home page get with success")
	}
}
