package groupietracker

import (
	"encoding/json"
	utils "groupietracker/lib/utils"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

type DateModel struct {
	Index []struct {
		ID    int      `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}

type DateFormat struct {
	Id    string
	Day   string
	Month string
	Year  string
}

func DatesList(res http.ResponseWriter, req *http.Request) {
	if utils.ValidateRequest(req, res, "/events/", http.MethodGet) {
		pagePath := "events"
		res.WriteHeader(http.StatusOK)
		url := "https://groupietrackers.herokuapp.com/api/dates"
		data, err := utils.GetAPI(url)
		if err != nil {
			utils.RenderPage("500", nil, res)
			log.Println("❌ Internal Server Error ", err)
			return
		}
		var _datesList DateModel
		err = json.Unmarshal(data, &_datesList)
		if err != nil {
			utils.RenderPage("500", nil, res)
			log.Println("❌ Internal Server Error ", err)
			return
		}
		datesList := []string{}
		datesFormatList := []DateFormat{}

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
			dateFormat := DateFormat{
				_dateFormat[1],
				date[0],
				date[1],
				date[2],
			}
			datesFormatList = append(datesFormatList, dateFormat)
		}

		utils.RenderPage(pagePath, &datesFormatList, res)
		log.Println("✅ Date page get with success")
	}
}
