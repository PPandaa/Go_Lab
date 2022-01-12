package desk

import (
	"GoLab/auth"
	"GoLab/database/mongodb"
	"GoLab/dependency"
	"GoLab/guard"
	"GoLab/server"

	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/bitly/go-simplejson"
	"gopkg.in/mgo.v2/bson"
)

type GroupStruct struct {
	UnderId     string    `json:"_id" bson:"_id"`
	GroupID     string    `json:"id" bson:"GroupID"`
	GroupName   string    `json:"name" bson:"GroupName"`
	ParentID    string    `json:"parentId" bson:"ParentID"`
	TimeZone    string    `json:"timeZone" bson:"TimeZone"`
	Description string    `json:"description" bson:"Description"`
	CreatedAt   time.Time `json:"createdAt" bson:"CreatedAt"`
	UpdatedAt   time.Time `json:"updatedAt" bson:"UpdatedAt"`
}

type EquipmentStruct struct {
	UnderId       string         `json:"_id" bson:"_id"`
	GroupID       string         `bson:"GroupID"`
	EquipmentType string         `bson:"EquipmentType"`
	EquipmentID   string         `json:"id" bson:"EquipmentID"`
	EquipmentName string         `json:"name" bson:"EquipmentName"`
	ImageURL      string         `json:"imageUrl" bson:"ImageURL"`
	Capacity      CapacityStruct `bson:"Capacity"`
	CreatedAt     time.Time      `json:"createdAt" bson:"CreatedAt"`
	UpdatedAt     time.Time      `json:"updatedAt" bson:"UpdatedAt"`
}

type CapacityStruct struct {
	Value    float32 `bson:"Value"`
	Unit     string  `bson:"Unit"`
	Duration string  `bson:"Duration"`
}

var (
	LastWaconnTime time.Time
)

func RegisterOutbound() {

	content := map[string]interface{}{"name": server.AppNameC, "sourceId": "scada_" + server.AppNameL, "url": server.DAEMON_DATABROKER_API_URL.String(), "active": true}
	variable := map[string]interface{}{"input": content}
	httpRequestBody, _ := json.Marshal(map[string]interface{}{
		"query":     "mutation ($input: AddOutboundInput!) { addOutbound(input: $input) { outbound { id name url sourceId allowUnauthorized active connected } } }",
		"variables": variable,
	})
	request, _ := http.NewRequest("POST", dependency.IFP_DESK_API_URL.String(), bytes.NewBuffer(httpRequestBody))
	if server.Location == server.Cloud {
		request.Header.Set("X-Ifp-App-Secret", auth.IFPToken)
	} else {
		request.Header.Set("cookie", auth.IFPToken)
	}
	request.Header.Set("Content-Type", "application/json")
	response, _ := server.HttpClient.Do(request)
	m, _ := simplejson.NewFromReader(response.Body)
	if len(m.Get("errors").MustArray()) == 0 {
		guard.Logger.Info("register outbound " + server.AppNameC + " success")
	} else {
		guard.Logger.Info("outbound " + server.AppNameC + " is already exist")
	}

}

func GetGroup(mode string) {

	mongodb.ConnectCheck()

	httpRequestBody, _ := json.Marshal(map[string]interface{}{
		"query": "query groups { groups { _id id name createdAt updatedAt parentId timeZone description } }",
	})
	request, _ := http.NewRequest("POST", dependency.IFP_DESK_API_URL.String(), bytes.NewBuffer(httpRequestBody))
	if server.Location == server.Cloud {
		request.Header.Set("X-Ifp-App-Secret", auth.IFPToken)
	} else {
		request.Header.Set("cookie", auth.IFPToken)
	}
	request.Header.Set("Content-Type", "application/json")
	response, _ := server.HttpClient.Do(request)
	m, _ := simplejson.NewFromReader(response.Body)
	if len(m.Get("errors").MustArray()) == 0 {
		groups := []GroupStruct{}
		temp, _ := m.Get("data").Get("groups").MarshalJSON()
		json.Unmarshal(temp, &groups)
		for _, group := range groups {
			SetGroup(group.GroupID, group)
		}
	} else {
		guard.Logger.Fatal("GetGroup =>  " + m.Get("errors").GetIndex(0).Get("message").MustString())
	}

}

func SetGroup(groupID string, group interface{}) {

	mongodb.ConnectCheck()

	collection := mongodb.DB.C(mongodb.Test)
	existGroup := map[string]interface{}{}
	collection.Pipe([]bson.M{{"$match": bson.M{"GroupID": groupID}}}).One(&existGroup)
	if len(existGroup) == 0 {
		collection.Insert(group)
	} else {
		collection.Update(bson.M{"_id": existGroup["_id"]}, bson.M{"$set": group})
	}

}

func DeleteGroup(groupID string) {

	mongodb.ConnectCheck()

	collection := mongodb.DB.C(mongodb.Test)
	filter := bson.M{"GroupID": groupID}
	_, err := collection.RemoveAll(filter)
	if err != nil {
		guard.Logger.Error(err.Error())
	}

}

func GetMachine(mode string, groupUnderID ...string) {

	mongodb.ConnectCheck()

	var groupIDs []string
	gCollection := mongodb.DB.C(mongodb.Test)
	eCollection := mongodb.DB.C(mongodb.Test)
	if mode == "init" {
		var groupTopoResults []map[string]interface{}
		gCollection.Find(bson.M{}).All(&groupTopoResults)
		if len(groupTopoResults) != 0 {
			for _, groupTopoResult := range groupTopoResults {
				groupIDs = append(groupIDs, groupTopoResult["GroupID"].(string))
			}
		}
	} else if mode == "outbound" {
		var groupTopoResults map[string]interface{}
		gCollection.Find(bson.M{"_id": groupUnderID[0]}).One(&groupTopoResults)
		groupIDs = append(groupIDs, groupTopoResults["GroupID"].(string))
	} else if mode == "redis" {
		var groupTopoResults map[string]interface{}
		gCollection.Find(bson.M{"GroupID": groupUnderID[0]}).One(&groupTopoResults)
		groupIDs = append(groupIDs, groupTopoResults["GroupID"].(string))
	}

	if len(groupIDs) != 0 {
		httpRequestBody, _ := json.Marshal(map[string]interface{}{
			"query":     "query ($groupId: [ID!]!) { groupsByIds(ids: $groupId) { _id id name timeZone machines { _id id name createdAt updatedAt imageUrl } } }",
			"variables": map[string][]string{"groupId": groupIDs},
		})
		request, _ := http.NewRequest("POST", dependency.IFP_DESK_API_URL.String(), bytes.NewBuffer(httpRequestBody))
		if server.Location == server.Cloud {
			request.Header.Set("X-Ifp-App-Secret", auth.IFPToken)
		} else {
			request.Header.Set("cookie", auth.IFPToken)
		}
		request.Header.Set("Content-Type", "application/json")
		response, _ := server.HttpClient.Do(request)
		m, _ := simplejson.NewFromReader(response.Body)

		if len(m.Get("errors").MustArray()) == 0 {
			groupsLayer := m.Get("data").Get("groupsByIds")
			for indexOfGroups := 0; indexOfGroups < len(groupsLayer.MustArray()); indexOfGroups++ {
				groupID := groupsLayer.GetIndex(indexOfGroups).Get("id").MustString()
				machinesLayer := groupsLayer.GetIndex(indexOfGroups).Get("machines")
				for indexOfMachines := 0; indexOfMachines < len(machinesLayer.MustArray()); indexOfMachines++ {
					machineJSON, _ := machinesLayer.GetIndex(indexOfMachines).MarshalJSON()
					equipment := EquipmentStruct{GroupID: groupID, EquipmentType: "Machine"}
					json.Unmarshal(machineJSON, &equipment)
					dbEquipment := map[string]interface{}{}
					eCollection.Pipe([]bson.M{{"$match": bson.M{"EquipmentID": equipment.EquipmentID}}}).One(&dbEquipment)
					if len(dbEquipment) == 0 {
						eCollection.Insert(equipment)
					} else {
						eCollection.Update(bson.M{"_id": dbEquipment["_id"]}, bson.M{"$set": bson.M{"GroupID": equipment.GroupID, "ImageURL": equipment.ImageURL, "EquipmentType": "Machine", "EquipmentID": equipment.EquipmentID, "EquipmentName": equipment.EquipmentName}})
					}
				}
			}
		} else {
			guard.Logger.Fatal("GetMachine =>  " + m.Get("errors").GetIndex(0).Get("message").MustString())
		}
	}

}

func DeleteMachine(machineID string) {

	mongodb.ConnectCheck()

	collection := mongodb.DB.C(mongodb.Test)
	filter := bson.M{"EquipmentID": machineID}
	_, err := collection.RemoveAll(filter)
	if err != nil {
		guard.Logger.Error(err.Error())
	}

}

func GetStation(mode string, groupUnderID ...string) {

	mongodb.ConnectCheck()

	var groupIDs []string
	gCollection := mongodb.DB.C(mongodb.Test)
	eCollection := mongodb.DB.C(mongodb.Test)
	if mode == "init" {
		var groupTopoResults []map[string]interface{}
		gCollection.Find(bson.M{}).All(&groupTopoResults)
		if len(groupTopoResults) != 0 {
			for _, groupTopoResult := range groupTopoResults {
				groupIDs = append(groupIDs, groupTopoResult["GroupID"].(string))
			}
		}
	} else if mode == "outbound" {
		var groupTopoResults map[string]interface{}
		gCollection.Find(bson.M{"_id": groupUnderID[0]}).One(&groupTopoResults)
		groupIDs = append(groupIDs, groupTopoResults["GroupID"].(string))
	} else if mode == "redis" {
		var groupTopoResults map[string]interface{}
		gCollection.Find(bson.M{"GroupID": groupUnderID[0]}).One(&groupTopoResults)
		groupIDs = append(groupIDs, groupTopoResults["GroupID"].(string))
	}

	if len(groupIDs) != 0 {
		httpRequestBody, _ := json.Marshal(map[string]interface{}{
			"query":     "query ($groupId: [ID!]!) { groupsByIds(ids: $groupId) { _id id name createdAt updatedAt machines(isStation: true) { _id id name createdAt updatedAt } } }",
			"variables": map[string][]string{"groupId": groupIDs},
		})
		request, _ := http.NewRequest("POST", dependency.IFP_DESK_API_URL.String(), bytes.NewBuffer(httpRequestBody))
		if server.Location == server.Cloud {
			request.Header.Set("X-Ifp-App-Secret", auth.IFPToken)
		} else {
			request.Header.Set("cookie", auth.IFPToken)
		}
		request.Header.Set("Content-Type", "application/json")
		response, _ := server.HttpClient.Do(request)
		m, _ := simplejson.NewFromReader(response.Body)

		if len(m.Get("errors").MustArray()) == 0 {
			groupsLayer := m.Get("data").Get("groupsByIds")
			for indexOfGroups := 0; indexOfGroups < len(groupsLayer.MustArray()); indexOfGroups++ {
				groupID := groupsLayer.GetIndex(indexOfGroups).Get("id").MustString()
				machinesLayer := groupsLayer.GetIndex(indexOfGroups).Get("machines")
				for indexOfMachines := 0; indexOfMachines < len(machinesLayer.MustArray()); indexOfMachines++ {
					machineJSON, _ := machinesLayer.GetIndex(indexOfMachines).MarshalJSON()
					equipment := EquipmentStruct{GroupID: groupID, EquipmentType: "Machine"}
					json.Unmarshal(machineJSON, &equipment)
					dbEquipment := map[string]interface{}{}
					eCollection.Pipe([]bson.M{{"$match": bson.M{"EquipmentID": equipment.EquipmentID}}}).One(&dbEquipment)
					if len(dbEquipment) == 0 {
						eCollection.Insert(equipment)
					} else {
						eCollection.Update(bson.M{"_id": dbEquipment["_id"]}, bson.M{"$set": bson.M{"GroupID": equipment.GroupID, "ImageURL": equipment.ImageURL, "EquipmentType": "Machine", "EquipmentID": equipment.EquipmentID, "EquipmentName": equipment.EquipmentName}})
					}
				}
			}
		} else {
			guard.Logger.Fatal("GetStation =>  " + m.Get("errors").GetIndex(0).Get("message").MustString())
		}
	}

}

func DeleteStation(stationID string) {

	mongodb.ConnectCheck()

	collection := mongodb.DB.C(mongodb.Test)
	filter := bson.M{"EquipmentID": stationID}
	_, err := collection.RemoveAll(filter)
	if err != nil {
		guard.Logger.Error(err.Error())
	}

}
