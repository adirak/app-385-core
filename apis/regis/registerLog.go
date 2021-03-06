// Package regis contains an apis for Register data
package regis

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/adirak/app-385-core/data"
	"github.com/adirak/app-385-core/dbexec"
)

// SaveRegisterLog is function to set data in register_log table
func SaveRegisterLog(reqt data.ReqtData) (resp data.RespData) {

	log.Println("SaveRegisterLog Start")
	defer log.Println("SaveRegisterLog End")

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

			//log.Println("Step 1")

			// Convert data
			uID := userID.(string)
			eID := int64(eventID.(float64))

			//log.Println("Step 2")

			// Load Register Log data
			rLog, err2 := dbexec.LoadRegisterLog(uID, eID)
			if err2 == nil {

				//log.Println("Step 3")

				if rLog.RegisterLogID == 0 {

					// Insert data to table because it is never create record
					err = dbexec.InsertRegisterLog(uID, eID)

					//log.Println("Step 4")
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

					//log.Println("Step 5")

					// Update Register Log
					err = dbexec.UpdateRegisterLog(rLog)
					if err == nil {
						resp.Success = true
					}

					//log.Println("Step 6")

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

	log.Println("SetStepOfRegisterLog Start")
	defer log.Println("SetStepOfRegisterLog End")

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

					// Return success
					if err == nil {
						resp.Success = true
					}

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

// LoadRegisterLog is function to load register log from db
func LoadRegisterLog(reqt data.ReqtData) (resp data.RespData) {

	log.Println("LoadRegisterLog Start")
	defer log.Println("LoadRegisterLog End")

	// Output data
	outData := make(map[string]interface{})

	// Input data
	inData := reqt.InData
	var err error

	if inData != nil {

		userID := inData["userId"]
		eventID := inData["eventId"]

		// Step description
		// 1.userInfo 2.address 3.item 4.confirm 5.confirm role 6.payment
		stepDesc := make(map[string]interface{})
		stepDesc["1"] = "User Info Page"
		stepDesc["2"] = "Address Page"
		stepDesc["3"] = "Item Size Page"
		stepDesc["4"] = "Confirm all steps and Role Page"
		stepDesc["5"] = "Confirm Consent Page"
		stepDesc["6"] = "Payment Page"
		outData["stepDesc"] = stepDesc

		if userID != nil && eventID != nil {

			// Convert data
			uID := userID.(string)
			eID := int64(eventID.(float64))

			// Load Register Log data
			rLog, err2 := dbexec.LoadRegisterLog(uID, eID)
			if err2 == nil {

				step := rLog.Step
				outData["userId"] = rLog.UserID
				outData["eventId"] = rLog.EventID
				outData["step"] = step

				// User Info Page
				if step == 1 || step == 4 {

					// User Info
					mapData := make(map[string]interface{})
					err = json.Unmarshal([]byte(rLog.UserInfo), &mapData)
					if err == nil {
						outData["userInfo"] = mapData
					}

				}

				// Address Page
				if step == 2 || step == 4 {

					// Address
					mapData := make(map[string]interface{})
					err = json.Unmarshal([]byte(rLog.Address), &mapData)
					if err == nil {
						outData["address"] = mapData
					}

				}

				// Item Size Page
				if step == 3 || step == 4 {

					// Item
					mapData := make(map[string]interface{})
					err = json.Unmarshal([]byte(rLog.Item), &mapData)
					if err == nil {
						outData["item"] = mapData
					}

				}

				// Confirm all steps and Role Page
				if step == 4 {

					runEvent, err2 := dbexec.LoadRunningEventByID(eID)
					if err2 == nil {
						ruleData, err3 := dbexec.LoadRuleFunction(runEvent.RuleID)
						if err3 == nil {
							outData["rule"] = ruleData
						} else {
							err = err3
						}
					} else {
						err = err2
					}

				}

				// Confirm Consent Page
				if step == 5 {

					runEvent, err2 := dbexec.LoadRunningEventByID(eID)
					if err2 == nil {

						mapData := make(map[string]interface{})
						err = json.Unmarshal([]byte(runEvent.Data), &mapData)
						if err == nil {
							consent := mapData["consent"]
							outData["consent"] = consent
						}

					} else {
						err = err2
					}

				}

				// Payment Page
				if step == 6 {

					// Confirm
					// TODO
					mapData := make(map[string]interface{})
					mapData["desc"] = "No implement yet"
					outData["payment"] = mapData

				}

				// Unsupport Page
				if step < 1 || step > 6 {
					err = errors.New("step is not corrected")
				}

				// Return Success
				if err == nil {
					resp.Success = true
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
