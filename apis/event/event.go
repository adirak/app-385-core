// Package event contains an apis for Event data
package event

import (
	"log"

	"github.com/adirak/app-385-core/data"
	"github.com/adirak/app-385-core/dbexec"
	"github.com/adirak/app-385-core/util"
)

// GetEvents is api function to list events from database
func GetEvents(reqt data.ReqtData) (resp data.RespData) {

	log.Println("GetEvents Start")
	defer log.Println("GetEvents End")

	// Output data
	outData := make(map[string]interface{})

	events, err := dbexec.LoadRunningEvents()

	if err == nil {

		if events != nil {

			list, err1 := util.EventsToArray(events)
			if err1 == nil {
				outData["list"] = list
				resp.Success = true
			} else {
				err = err1
			}

		}

	}

	// Footer response
	resp.OutData = outData
	resp.Err = err
	return resp
}

// CreateEvent is api function to create running_event record
func CreateEvent(reqt data.ReqtData) (resp data.RespData) {

	log.Println("CreateEvent Start")
	defer log.Println("CreateEvent End")

	// Output data
	outData := make(map[string]interface{})

	// Input data
	inData := reqt.InData
	name := inData["name"].(string)
	active := inData["active"].(bool)
	data := inData["data"]

	err := dbexec.CreateRunningEvent(name, data, active)
	if err == nil {
		resp.Success = true
	}

	// Footer response
	resp.OutData = outData
	resp.Err = err
	return resp
}
