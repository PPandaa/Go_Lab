package ehs

import (
	"bytes"
	"encoding/json"
	"net/http"

	"GoLab/auth"
	"GoLab/dependency"
	"GoLab/guard"
	"GoLab/server"

	"github.com/bitly/go-simplejson"
)

var (
	httpClient = &http.Client{}
)

type KpiStruct struct {
	ID              string             `json:"id"`
	Name            string             `json:"name"`
	GenericKpiLists []GenericKpiStruct `json:"genericKpiList"`
}

type GenericKpiStruct struct {
	RelationKey string  `json:"relationKey"`
	Kind        string  `json:"kind"`
	Year        int     `json:"year"`
	Hour        int     `json:"hour"`
	KPI         float64 `json:"kpi"`
}

type SetDMYGenericKpiStruct struct {
	RelationKey string  `json:"relationKey"`
	Kind        string  `json:"kind"`
	Year        int     `json:"year"`
	KPI         float64 `json:"kpi"`
}

type SetHourGenericKpiStruct struct {
	RelationKey string         `json:"relationKey"`
	Kind        string         `json:"kind"`
	Year        int            `json:"year"`
	ByHours     []ByHourStruct `json:"byHours"`
}

type ByHourStruct struct {
	Hour int     `json:"hour"`
	KPI  float64 `json:"kpi"`
}

func GQL_Query_genericKpiListByGroupIDs(groupIDs []string, relationKeys []string, years []int, kinds []string) []KpiStruct {

	var result []KpiStruct

	variable := map[string]interface{}{
		"groupIds":     groupIDs,
		"relationKeys": relationKeys,
		"years":        years,
		"kinds":        kinds,
	}
	httpRequestBody, _ := json.Marshal(map[string]interface{}{
		"query": `query ($groupIds: [ID!]!, $relationKeys: [String!]!, $years: [NonNegativeInt!]!, $kinds: [GenericKpiKind!]!) { 
			groupsByIds(ids: $groupIds) {
				id
				name
				genericKpiList(relationKeys: $relationKeys, years: $years, kinds: $kinds) {
					relationKey
					kind
					year
					hour
					kpi
				}
			}
		}`,
		"variables": variable,
	})

	request, _ := http.NewRequest("POST", dependency.IFP_DESK_API_URL.String(), bytes.NewBuffer(httpRequestBody))
	request.Header.Set("Content-Type", "application/json")
	if server.Location == server.Cloud {
		request.Header.Set("X-Ifp-App-Secret", auth.IFPToken)
	} else {
		request.Header.Set("cookie", auth.IFPToken)
	}

	response, _ := httpClient.Do(request)
	m, _ := simplejson.NewFromReader(response.Body)
	if len(m.Get("errors").MustArray()) != 0 {
		guard.Logger.Error("GQL_Query_genericKpiListByGroupIDs GraphQL Error")
	} else {
		temp, err := json.Marshal(m.Get("data").Get("groupsByIds"))
		if err != nil {
			guard.Logger.Error(err.Error())
		} else {
			json.Unmarshal(temp, &result)
		}
	}

	return result

}

func GQL_Query_genericKpiListByMachineIDs(machineIDs []string, relationKeys []string, years []int, kinds []string) []KpiStruct {

	var result []KpiStruct

	variable := map[string]interface{}{
		"machineIds":   machineIDs,
		"relationKeys": relationKeys,
		"years":        years,
		"kinds":        kinds,
	}
	httpRequestBody, _ := json.Marshal(map[string]interface{}{
		"query": `query ($machineIds: [ID!]!, $relationKeys: [String!]!, $years: [NonNegativeInt!]!, $kinds: [GenericKpiKind!]!) { 
			machinesByIds(ids: $machineIds) {
				id
				name
				genericKpiList(relationKeys: $relationKeys, years: $years, kinds: $kinds) {
					relationKey
					kind
					year
					hour
					kpi
				}
			}
		}`,
		"variables": variable,
	})

	request, _ := http.NewRequest("POST", dependency.IFP_DESK_API_URL.String(), bytes.NewBuffer(httpRequestBody))
	request.Header.Set("Content-Type", "application/json")
	if server.Location == server.Cloud {
		request.Header.Set("X-Ifp-App-Secret", auth.IFPToken)
	} else {
		request.Header.Set("cookie", auth.IFPToken)
	}

	response, _ := httpClient.Do(request)
	m, _ := simplejson.NewFromReader(response.Body)
	if len(m.Get("errors").MustArray()) != 0 {
		guard.Logger.Error("GQL_Query_genericKpiListByMachineIDs GraphQL Error")
	} else {
		temp, err := json.Marshal(m.Get("data").Get("machinesByIds"))
		if err != nil {
			guard.Logger.Error(err.Error())
		} else {
			json.Unmarshal(temp, &result)
		}
	}

	return result

}

func GQL_Query_genericKpiListByParameterIDs(parameterIDs []string, relationKeys []string, years []int, kinds []string) []KpiStruct {

	var result []KpiStruct

	variable := map[string]interface{}{
		"parameterIds": parameterIDs,
		"relationKeys": relationKeys,
		"years":        years,
		"kinds":        kinds,
	}
	httpRequestBody, _ := json.Marshal(map[string]interface{}{
		"query": `query ($parameterIds: [ID!]!, $relationKeys: [String!]!, $years: [NonNegativeInt!]!, $kinds: [GenericKpiKind!]!) { 
			parametersByIds(ids: $parameterIds) {
				id
				name
				genericKpiList(relationKeys: $relationKeys, years: $years, kinds: $kinds) {
					relationKey
					kind
					year
					hour
					kpi
				}
			}
		}`,
		"variables": variable,
	})

	request, _ := http.NewRequest("POST", dependency.IFP_DESK_API_URL.String(), bytes.NewBuffer(httpRequestBody))
	request.Header.Set("Content-Type", "application/json")
	if server.Location == server.Cloud {
		request.Header.Set("X-Ifp-App-Secret", auth.IFPToken)
	} else {
		request.Header.Set("cookie", auth.IFPToken)
	}

	response, _ := httpClient.Do(request)
	m, _ := simplejson.NewFromReader(response.Body)
	if len(m.Get("errors").MustArray()) != 0 {
		guard.Logger.Error("GQL_Query_genericKpiListByParameterIDs GraphQL Error")
	} else {
		temp, err := json.Marshal(m.Get("data").Get("parametersByIds"))
		if err != nil {
			guard.Logger.Error(err.Error())
		} else {
			json.Unmarshal(temp, &result)
		}
	}

	return result

}

func GQL_Mutation_setGenericKpiList(category string, id string, kpiList interface{}) {

	var input map[string]interface{}
	switch category {
	case "Group":
		input = map[string]interface{}{
			"groupId": id,
			"kpiList": kpiList,
		}
	case "Machine":
		input = map[string]interface{}{
			"machineId": id,
			"kpiList":   kpiList,
		}
	case "Parameter":
		input = map[string]interface{}{
			"parameterId": id,
			"kpiList":     kpiList,
		}
	}

	variable := map[string]interface{}{
		"input": input,
	}
	httpRequestBody, _ := json.Marshal(map[string]interface{}{
		"query": `mutation setGenericKpiList($input: SetGenericKpiListInput!) {
			setGenericKpiList(input: $input) {
				kpiList {
					id
				}
			}
		}`,
		"variables": variable,
	})

	request, _ := http.NewRequest("POST", dependency.IFP_DESK_API_URL.String(), bytes.NewBuffer(httpRequestBody))
	request.Header.Set("Content-Type", "application/json")
	if server.Location == server.Cloud {
		request.Header.Set("X-Ifp-App-Secret", auth.IFPToken)
	} else {
		request.Header.Set("cookie", auth.IFPToken)
	}

	response, _ := httpClient.Do(request)
	m, _ := simplejson.NewFromReader(response.Body)
	if len(m.Get("errors").MustArray()) != 0 {
		guard.Logger.Error("GQL_Mutation_setGenericKpiList GraphQL Error")
	}

}
