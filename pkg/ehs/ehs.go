package ehs

import (
	"strconv"
	"strings"
)

func GetKinds(time string) []string {

	switch time {
	case "Hour":
		return []string{
			"HourSun",
			"HourMon",
			"HourTue",
			"HourWed",
			"HourThu",
			"HourFri",
			"HourSat",
		}
	case "Day":
		return []string{
			"DayWeekday",
			"DayHoliday",
		}
	case "Month":
		return []string{
			"MonthJan",
			"MonthFeb",
			"MonthMar",
			"MonthApr",
			"MonthMay",
			"MonthJun",
			"MonthJul",
			"MonthAug",
			"MonthSep",
			"MonthOct",
			"MonthNov",
			"MonthDec",
		}
	case "Year":
		return []string{"Year"}
	case "Unlimited":
		return []string{"Unlimited"}
	default:
		return []string{}
	}

}

func CalculateAvgValue(category string, categoryKpis []KpiStruct) []map[string]interface{} {

	categoryKindAvgKpi := []map[string]interface{}{}

	for _, categoryKpi := range categoryKpis {
		// generate kpi list by kind
		kindKpis := map[string][]float64{}
		for _, genericKpi := range categoryKpi.GenericKpiLists {
			if strings.HasPrefix(genericKpi.Kind, "Hour") {
				key := genericKpi.Kind + "_" + strconv.Itoa(genericKpi.Hour)
				kindKpis[key] = append(kindKpis[key], genericKpi.KPI)
			} else {
				kindKpis[genericKpi.Kind] = append(kindKpis[genericKpi.Kind], genericKpi.KPI)
			}
		}
		// calculate average kpi
		kindAvgKpi := map[string]float64{}
		for kind, kpis := range kindKpis {
			var sumKpi float64
			for _, kpi := range kpis {
				sumKpi += kpi
			}

			var avgKpi float64
			if len(kpis) != 0 {
				avgKpi = sumKpi / float64(len(kpis))
			}

			kindAvgKpi[kind] = avgKpi
		}
		// build return value
		element := map[string]interface{}{}
		switch category {
		case "Group":
			element = map[string]interface{}{
				"GroupID":    categoryKpi.ID,
				"KindAvgKpi": kindAvgKpi,
			}
		case "Machine":
			element = map[string]interface{}{
				"MachineID":  categoryKpi.ID,
				"KindAvgKpi": kindAvgKpi,
			}
		case "Parameter":
			element = map[string]interface{}{
				"ParameterID": categoryKpi.ID,
				"KindAvgKpi":  kindAvgKpi,
			}
		}
		categoryKindAvgKpi = append(categoryKindAvgKpi, element)
	}

	return categoryKindAvgKpi

}

func BuildInput(category string, relationkey string, year int, categorysKindAvgKpi []map[string]interface{}) []map[string]interface{} {

	var categorysKpi []map[string]interface{}

	for _, categoryKindAvgKpi := range categorysKindAvgKpi {
		var kpiList []interface{}
		for kind, avgKpi := range categoryKindAvgKpi["KindAvgKpi"].(map[string]float64) {
			// format gql input value
			var input interface{}
			if strings.HasPrefix(kind, "Hour") {
				// Ex: HourSun_1
				temp := strings.Split(kind, "_")
				realKind := temp[0]
				hour, _ := strconv.Atoi(temp[1])
				input = SetHourGenericKpiStruct{
					RelationKey: relationkey,
					Kind:        realKind,
					Year:        year,
					ByHours: []ByHourStruct{
						{
							Hour: hour,
							KPI:  avgKpi,
						},
					},
				}
			} else {
				input = SetDMYGenericKpiStruct{
					RelationKey: relationkey,
					Kind:        kind,
					Year:        year,
					KPI:         avgKpi,
				}
			}
			kpiList = append(kpiList, input)
		}
		// build return value
		element := map[string]interface{}{}
		switch category {
		case "Group":
			element = map[string]interface{}{
				"GroupID": categoryKindAvgKpi["GroupID"],
				"KpiList": kpiList,
			}
		case "Machine":
			element = map[string]interface{}{
				"MachineID": categoryKindAvgKpi["MachineID"],
				"KpiList":   kpiList,
			}
		case "Parameter":
			element = map[string]interface{}{
				"ParameterID": categoryKindAvgKpi["ParameterID"],
				"KpiList":     kpiList,
			}
		}
		categorysKpi = append(categorysKpi, element)
	}

	return categorysKpi

}
