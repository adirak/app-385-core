// Package webhook contains an apis for Webhook
package webhook

import (
	"encoding/json"
	"log"

	"github.com/adirak/app-385-core/trigger"

	"github.com/adirak/app-385-core/dbexec"

	"github.com/adirak/app-385-core/data"
)

// StravaCallBack is api function for webhook of strava to call back
func StravaCallBack(reqt data.ReqtData) (resp data.RespData) {

	log.Println("StravaCallBack Start")
	defer log.Println("StravaCallBack End")

	// Output data
	outData := make(map[string]interface{})

	// Input data
	inData := reqt.InData

	// Check verifying token from strava
	isVerified, challengeCode := isVerifiedToken(inData)
	if isVerified {

		// Make data and force response to this pattern
		outData["hub.challenge"] = challengeCode
		resp.ForceResponseData = outData
		log.Println("WebhookCallBack Subscribe hub.challenge: ", challengeCode)
		return resp

	}

	// Convert map to json string
	bData, err := json.Marshal(inData)
	if err == nil {

		// Prepare log
		actionType := inData["object_type"].(string)
		jsonStr := string(bData)

		// Request data is Json
		// Insert data to log
		err = dbexec.InsertWebhookLog(actionType, jsonStr)

		if err == nil {

			log.Println("WebhookCallBack Insert Log Success")

			stravaIDFloat := inData["owner_id"].(float64)
			stravaID := int64(stravaIDFloat)

			// Insert data to Activity Queue
			err = dbexec.InsertActivityQueue(stravaID)

			if err == nil {
				log.Println("WebhookCallBack Insert Activity Queue Success")

				// Call activity-queue-trigger
				err = trigger.CallActivityQueueTrigger()
				if err != nil {
					log.Println("WebhookCallBack Call trigger fail : ", err.Error())
				}

				resp.Success = true
				return resp
			}

		}

	}

	// Footer response
	resp.OutData = outData
	resp.Err = err
	return resp
}

// isVirifiedToken is checking virifying token from strava
func isVerifiedToken(inData map[string]interface{}) (bool, string) {

	verifyTokenKey := "hub.verify_token"
	challengeKey := "hub.challenge"

	verifyToken := inData[verifyTokenKey]
	if verifyToken == "STRAVA" {
		challenge := inData[challengeKey].(string)
		return true, challenge
	}

	return false, ""
}
