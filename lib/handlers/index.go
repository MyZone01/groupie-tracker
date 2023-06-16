package groupie_tracker

import (
	utils "groupie_tracker/lib/utils"
	"log"
	"net/http"
)

func Index(res http.ResponseWriter, req *http.Request) {
	if utils.ValidateRequest(req, res, "/", http.MethodGet) {
		pagePath := "index"
		utils.RenderPage(pagePath, nil, res)
		log.Println("✅ Home page get with success")
	}
}
