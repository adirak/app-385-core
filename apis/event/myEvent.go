// Package event contains an apis for Event data
package event

import (
	"log"

	"github.com/adirak/app-385-core/data"
	"github.com/adirak/app-385-core/dbexec"
	"github.com/adirak/app-385-core/util"
)

// GetMyEvent is api function to list my events from database
func GetMyEvent(reqt data.ReqtData) (resp data.RespData) {

	// Output data
	outData := make(map[string]interface{})

	// Input data
	inData := reqt.InData
	userID := inData["userId"].(string)
	log.Println("**userID : ", userID)

	// Get Data from eventRegister
	myEvents, err := dbexec.LoadMyEvent(userID)

	if err == nil {
		list := make([]map[string]interface{}, 0, 5)

		if myEvents != nil {
			for _, event := range myEvents {
				runEvent, err1 := dbexec.LoadRunningEvent(event)
				if err1 == nil {
					eventMap, err2 := util.EventToMap(runEvent)
					if err2 == nil {
						list = append(list, eventMap)
					} else {
						err = err2
					}
				} else {
					err = err1
				}
			}
		}

		if err == nil {
			outData["list"] = list
			resp.Success = true
		}
	}

	// Footer response
	resp.OutData = outData
	resp.Err = err
	return resp
}
