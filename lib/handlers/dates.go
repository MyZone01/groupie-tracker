package groupietracker

import (
	models "groupietracker/lib/models"
	utils "groupietracker/lib/utils"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

func DatesList(res http.ResponseWriter, req *http.Request) {
	if utils.ValidateRequest(req, res, "/events/", http.MethodGet) {
		pagePath := "events"
		_datesList, err := utils.GetDates(res)
		if err {
			return
		}
		datesList := []string{}
		datesFormatList := []models.DateFormat{}

		for _, dm := range _datesList.Index {
			id := dm.ID
			for _, date := range dm.Dates {
				datesList = append(datesList, strings.ReplaceAll(date, "*", "")+"/"+strconv.Itoa(id))
			}
		}

		sort.SliceStable(datesList, func(i, j int) bool {
			date1, _ := time.Parse("02-01-2006", strings.Split(datesList[i], "/")[0])
			date2, _ := time.Parse("02-01-2006", strings.Split(datesList[j], "/")[0])
			return date1.Before(date2)
		})

		for _, _date := range datesList {
			_dateFormat := strings.Split(_date, "/")
			date := strings.Split(utils.FormatDates(_dateFormat[0]), " ")
			dateFormat := models.DateFormat{
				Id: _dateFormat[1],
				Day: date[0],
				Month: date[1],
				Year: date[2],
			}
			datesFormatList = append(datesFormatList, dateFormat)
		}

		utils.RenderPage(pagePath, &datesFormatList, res)
		log.Println("✅ Date page get with success")
	}
}
