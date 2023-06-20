package groupietracker

import (
	"bufio"
	"fmt"
	models "groupietracker/lib/models"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"text/template"
)

type MapLocation struct {
	IsCity bool
	Name   string
	Number int
}

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

func GetAPI(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		log.Println("Failed to make API Request:", err.Error())
		return nil, err
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("🚨 Failed to read API response body:", err.Error())
		return nil, err
	}

	return data, nil
}

func RenderPage(pagePath string, data any, res http.ResponseWriter) {
	files := []string{"templates/base.html", "templates/" + pagePath + ".html"}
	tpl, err := template.ParseFiles(files...)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		files := []string{"templates/base.html", "templates/500.html"}
		tpl, err := template.ParseFiles(files...)
		tpl.Execute(res, data)
		log.Println("🚨 " + err.Error())
	} else {
		tpl.Execute(res, data)
	}
}

func FormatLocations(_relation models.RelationModel) []models.Event {
	locations := []models.Event{}
	for _locations, _dates := range _relation.DatesLocations {
		__locations := strings.Split(_locations, "-")
		var l models.Event
		if len(__locations) == 2 {
			l.City = strings.Title(strings.ReplaceAll(__locations[0], "_", " "))
			l.Country = strings.Title(strings.ReplaceAll(__locations[1], "_", " "))
		}
		for i, date := range _dates {
			_dates[i] = FormatDates(date)
		}
		l.Dates = _dates
		locations = append(locations, l)
	}
	return locations
}

func LoadEnv(path string) error {
	file, err := os.Open(path)
	if err != nil {
		log.Println("🚨 " + err.Error())
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			log.Println("🚨 Your env file must be set")
		}
		key := parts[0]
		value := parts[1]
		os.Setenv(key, value)
	}
	return scanner.Err()
}

func FormatDates(date string) string {
	months := map[string]string{
		"01": "January",
		"02": "February",
		"03": "March",
		"04": "April",
		"05": "May",
		"06": "June",
		"07": "July",
		"08": "August",
		"09": "September",
		"10": "October",
		"11": "November",
		"12": "December",
	}
	_date := strings.Split(date, "-")
	if len(_date) == 3 {
		return fmt.Sprintf("%s %s %s", _date[0], months[_date[1]], _date[2])
	} else {
		return ""
	}
}

func CreateMap() []MapLocation {
	cells := make([]MapLocation, 1925)
	northAmerica := []int{
		19, 20, 21, 73, 74, 75, 76, 77, 123, 124, 127, 128, 129, 130, 131, 132, 177, 178, 179, 181, 182, 183, 184, 185, 186, 187, 188, 228, 229, 230, 231, 232, 233, 234, 237, 238, 239, 240, 241, 242, 243, 283, 284, 285, 286, 287, 288, 289, 293, 294, 295, 296, 297, 332, 333, 334, 337, 339, 340, 341, 342, 343, 348, 349, 350, 351, 352, 387, 388, 389, 390, 391, 392, 393, 394, 395, 396, 397, 398, 399, 404, 405, 406, 407, 441, 442, 443, 444, 445, 446, 447, 448, 449, 450, 451, 452, 453, 454, 455, 459, 460, 497, 498, 499, 500, 501, 502, 503, 504, 505, 506, 507, 509, 510, 511, 514, 552, 553, 554, 555, 556, 557, 558, 559, 560, 561, 562, 564, 565, 566, 607, 610, 611, 612, 613, 614, 615, 616, 619, 620, 621, 622, 666, 667, 668, 669, 670, 671, 672, 674, 675, 676, 722, 723, 724, 725, 726, 727, 728, 729, 730, 731, 778, 779, 780, 781, 782, 783, 784, 785, 786, 834, 835, 836, 837, 838, 839, 889, 890, 891, 892, 893, 894, 944, 945, 946, 947, 948, 949, 1000, 1001, 1002, 1005, 1056, 1057, 1112, 1113, 1114, 1169,
	}
	southAmerica := []int{
		1225, 1226, 1227, 1228, 1229, 1281, 1282, 1283, 1284, 1285, 1335, 1336, 1337, 1338, 1339, 1340, 1341, 1342, 1390, 1391, 1392, 1393, 1394, 1395, 1396, 1397, 1446, 1447, 1448, 1449, 1450, 1451, 1502, 1503, 1504, 1505, 1506, 1557, 1558, 1559, 1560, 1612, 1613, 1614, 1666, 1667, 1668, 1721, 1722, 1776, 1777, 1831, 1832, 1887,
	}
	asia := []int{
		154, 208, 209, 210, 262, 263, 264, 265, 314, 315, 316, 317, 318, 319, 320, 321, 322, 323, 325, 326, 369, 370, 371, 372, 373, 374, 375, 376, 377, 378, 379, 380, 381, 382, 423, 424, 425, 426, 427, 428, 429, 430, 431, 432, 433, 434, 435, 436, 437, 438, 439, 440, 478, 479, 480, 481, 482, 483, 484, 485, 486, 487, 488, 489, 490, 491, 492, 493, 494, 495, 532, 533, 534, 535, 536, 537, 538, 539, 540, 541, 542, 543, 544, 545, 546, 547, 548, 549, 550, 587, 588, 589, 590, 591, 592, 593, 594, 595, 596, 597, 598, 599, 600, 601, 603, 642, 643, 644, 645, 646, 647, 648, 649, 650, 651, 652, 653, 654, 657, 658, 697, 698, 699, 700, 701, 702, 703, 704, 705, 706, 707, 708, 709, 712, 752, 753, 754, 755, 756, 757, 758, 759, 760, 761, 762, 763, 764, 807, 808, 809, 810, 811, 812, 813, 814, 815, 816, 817, 818, 819, 821, 859, 860, 861, 862, 863, 864, 865, 866, 867, 868, 869, 870, 871, 872, 873, 876, 913, 914, 916, 917, 918, 919, 920, 921, 922, 923, 924, 925, 926, 928, 930, 967, 968, 969, 970, 972, 973, 974, 975, 976, 977, 978, 979, 980, 981, 982, 1022, 1023, 1024, 1025, 1026, 1027, 1029, 1030, 1031, 1034, 1035, 1036, 1037, 1079, 1080, 1081, 1085, 1089, 1090, 1135, 1144, 1145, 1200, 1202, 1203, 1255, 1256, 1257, 1311,
	}
	africa := []int{
		961, 962, 963, 965, 1020, 1021, 1019, 1018, 1017, 1016, 1015, 1069, 1070, 1071, 1072, 1073, 1074, 1075, 1076, 1077, 1124, 1125, 1126, 1127, 1128, 1129, 1130, 1131, 1132, 1133, 1179, 1180, 1181, 1182, 1183, 1184, 1185, 1186, 1187, 1188, 1189, 1235, 1236, 1237, 1238, 1239, 1240, 1241, 1242, 1243, 1244, 1293, 1294, 1295, 1296, 1297, 1298, 1348, 1349, 1350, 1351, 1352, 1404, 1405, 1406, 1407, 1459, 1460, 1461, 1462, 1514, 1515, 1516, 1518, 1573, 1569, 1570, 1571, 1624, 1625,
	}
	europe := []int{
		361, 362, 415, 416, 417, 469, 465, 470, 471, 472, 523, 524, 525, 527, 582, 580, 579, 578, 632, 634, 636, 631, 311, 312, 251, 519, 520, 686, 687, 689, 688, 690, 691, 692, 637, 638, 583, 528, 473, 418, 474, 475, 365, 421, 476, 477, 531, 530, 529, 584, 585, 586, 639, 640, 641, 693, 694, 695, 696, 795, 796, 850, 851, 797, 742, 743, 744, 745, 746, 747, 748, 749, 750, 751, 805, 804, 803, 802, 801, 799, 798, 800, 854, 856, 857, 911,
	}
	oceania := []int{
		1206, 1262, 1263, 1317, 1369, 1370, 1372, 1422, 1423, 1424, 1425, 1426, 1427, 1428, 1476, 1477, 1478, 1479, 1480, 1481, 1482, 1483, 1531, 1532, 1533, 1534, 1535, 1536, 1537, 1538, 1586, 1587, 1590, 1591, 1592, 1645, 1646, 1650, 1705,
	}
	for _, index := range northAmerica {
		cells[index-1].Name = "NorthAmerica"
	}
	for _, index := range southAmerica {
		cells[index-1].Name = "SouthAmerica"
	}
	for _, index := range africa {
		cells[index-1].Name = "Africa"
	}
	for _, index := range asia {
		cells[index-1].Name = "Asia"
	}
	for _, index := range europe {
		cells[index-1].Name = "Europe"
	}
	for _, index := range oceania {
		cells[index-1].Name = "Oceania"
	}
	return cells
}

func FormatMap(data string) []MapLocation {
	cities := map[string]int{
		"argentina":            1615,
		"australia":            1480,
		"austria":              746,
		"belarus":              640,
		"belgium":              632,
		"brazil":               1450,
		"canada":               559,
		"chile":                1666,
		"china":                816,
		"colombia":             1281,
		"costa_rica":           1113,
		"czechia":              857,
		"denmark":              634,
		"finland":              527,
		"france":               742,
		"french_polynesia":     1615,
		"germany":              688,
		"greece":               911,
		"hungary":              745,
		"india":                1030,
		"indonesia":            1255,
		"ireland":              519,
		"italy":                854,
		"japan":                876,
		"mexico":               1002,
		"netherlands":          632,
		"netherlands_antilles": 1615,
		"new_caledonia":        1615,
		"new_zealand":          1650,
		"norway":               523,
		"peru":                 1390,
		"philippines":          1615,
		"poland":               691,
		"portugal":             795,
		"qatar":                972,
		"romania":              805,
		"saudi_arabia":         1025,
		"slovakia":             801,
		"south_korea":          819,
		"spain":                851,
		"sweden":               525,
		"switzerland":          744,
		"taiwan":               928,
		"thailand":             1034,
		"uk":                   520,
		"united_arab_emirates": 969,
		"usa":                  836,
	}
	cells := CreateMap()
	for name, index := range cities {
		cells[index-1].Name = strings.Title(strings.ReplaceAll(name, "_", " "))
		cells[index-1].IsCity = true
		cells[index-1].Number = strings.Count(data, "-"+name)
	}
	return cells
}
