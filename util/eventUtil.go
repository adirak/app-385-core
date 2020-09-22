// Package util contains the utility function for this project
package util

import (
	"encoding/json"
	"log"

	"github.com/adirak/app-385-core/data"
)

// EventToMap is function to convert event data to map
func EventToMap(event data.RunningEvent) (eventMap map[string]interface{}, err error) {

	data := event.Data
	eventMap = make(map[string]interface{})
	err = json.Unmarshal([]byte(data), &eventMap)
	if err == nil {
		eventMap["eventId"] = event.EventID
		eventMap["name"] = event.Name
	} else {
		log.Println("err : ", err.Error())
	}

	return eventMap, err
}

// EventsToArray is function to convert event data to map
func EventsToArray(events []data.RunningEvent) (list []map[string]interface{}, err error) {

	list = make([]map[string]interface{}, 0, 5)

	for _, event := range events {
		eventMap, err1 := EventToMap(event)
		if err1 == nil {
			list = append(list, eventMap)
		} else {
			err = err1
		}
	}

	return list, err

}
