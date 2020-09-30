// Package regis contains an apis for Register data
package regis

import (
	"encoding/json"
	"errors"

	"github.com/adirak/app-385-core/data"
	"github.com/adirak/app-385-core/dbexec"
)

// SaveRegisterLog is function to set data in register_log table
func SaveRegisterLog(reqt data.ReqtData) (resp data.RespData) {

	// Output data
	outData := make(map[string]interface{})

	// Input data
	inData := reqt.InData
	var err error

	if inData != nil {

		userID := inData["userId"]
		eventID := inData["eventId"]
		userInfo := inData["userInfo"]
		address := inData["address"]
		item := inData["item"]
		step := inData["step"]

		if userID != nil && eventID != nil {

			// Convert data
			uID := userID.(string)
			eID := int64(eventID.(float64))

			// Load Register Log data
			rLog, err2 := dbexec.LoadRegisterLog(uID, eID)
			if err2 == nil {

				if rLog.RegisterLogID == 0 {

					// Insert data to table because it is never create record
					err = dbexec.InsertRegisterLog(uID, eID)
					if err == nil {

						// Load Register Log data again
						rLog, err = dbexec.LoadRegisterLog(uID, eID)

					}
				}

				// Data ready to update
				if err == nil && rLog.RegisterLogID > 0 {

					// Prepare Data
					if userInfo != nil {
						bData, err3 := json.Marshal(userInfo)
						if err3 == nil {
							rLog.UserInfo = string(bData)
						}
					}
					if address != nil {
						bData, err3 := json.Marshal(address)
						if err3 == nil {
							rLog.Address = string(bData)
						}
					}
					if item != nil {
						bData, err3 := json.Marshal(item)
						if err3 == nil {
							rLog.Item = string(bData)
						}
					}
					if step != nil {
						fStep := step.(float64)
						rLog.Step = int64(fStep)
					}

					// Update Register Log
					err = dbexec.UpdateRegisterLog(rLog)
					if err == nil {
						resp.Success = true
					}

				} else {
					if err == nil {
						err = errors.New("Cannot insert data to database")
					}
				}

			} else {
				err = err2
			}

		} else {

			resp.Msg = "Bad Request, userId or eventId is nil"
			resp.Code = 400
		}

	} else {
		resp.Msg = "Bad Request"
		resp.Code = 400
	}

	// Footer response
	resp.OutData = outData
	resp.Err = err
	return resp
}

// SetStepOfRegisterLog is function to set step column in register_log table
func SetStepOfRegisterLog(reqt data.ReqtData) (resp data.RespData) {

	// Output data
	outData := make(map[string]interface{})

	// Input data
	inData := reqt.InData
	var err error

	if inData != nil {

		userID := inData["userId"]
		eventID := inData["eventId"]
		step := inData["step"]

		if userID != nil && eventID != nil && step != nil {

			// Convert data
			uID := userID.(string)
			eID := int64(eventID.(float64))

			// Load Register Log data
			rLog, err2 := dbexec.LoadRegisterLog(uID, eID)
			if err2 == nil {

				if rLog.RegisterLogID > 0 {
					fStep := step.(float64)
					rLog.Step = int64(fStep)

					// Update step into database
					err = dbexec.SetStepRegisterLog(rLog)

				} else {
					if err == nil {
						err = errors.New("No register data in DB")
					}
				}

			} else {
				err = err2
			}

		} else {

			resp.Msg = "Bad Request, userId or eventId or step is nil"
			resp.Code = 400
		}

	} else {
		resp.Msg = "Bad Request"
		resp.Code = 400
	}

	// Footer response
	resp.OutData = outData
	resp.Err = err
	return resp

}
