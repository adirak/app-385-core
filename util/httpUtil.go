// Package util contains the utility function for this project
package util

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/adirak/app-385-core/data"
)

// GetReqtData is function to convert http.Request to ReqtData
func GetReqtData(r *http.Request) (reqtData data.ReqtData, err error) {

	// Make Input Json Data
	var inData map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&inData)
	if err != nil {
		reqBody, err1 := ioutil.ReadAll(r.Body)
		if err1 == nil {
			log.Println("Request body is not json : ", string(reqBody), "\nError : ", err.Error())
		} else {
			log.Println("Request body is not json : ", "\nError : ", err.Error())
		}

	}

	// Make parameter from url request
	if r.URL.Query() != nil {

		// Make if it is nil
		if inData == nil {
			inData = make(map[string]interface{})
		}

		// Loop to get value
		for name, values := range r.URL.Query() {
			buff := ""
			for _, value := range values {
				buff += value
			}
			inData[name] = buff
		}

	}

	// Set to ReqtData
	reqtData.InData = inData

	return reqtData, err
}

// SetJSONResponse is function to make data from http response to frontend
func SetJSONResponse(respData data.RespData, w http.ResponseWriter) {

	// Write Header
	w.Header().Set("Content-Type", "application/json")

	// Force response data if not nil
	if respData.ForceResponseData != nil {

		// Make Response JSON data
		json.NewEncoder(w).Encode(respData.ForceResponseData)

	} else {

		// Make Result data
		response := make(map[string]interface{})
		msgBody := getMessageBody(respData)

		if respData.Success {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(respData.Code)
		}

		// Set Response Data
		response["data"] = respData.OutData

		// Set Message Body
		response["message"] = msgBody

		// Make Response JSON data
		json.NewEncoder(w).Encode(response)

	}
}

// getMessageBody is function to make message body
func getMessageBody(respData data.RespData) map[string]interface{} {
	msgBody := make(map[string]interface{})

	if respData.Success {
		msgBody["code"] = 200

		if respData.Msg != "" {
			msgBody["desc"] = respData.Msg
		} else {
			msgBody["desc"] = "Successful"
		}

	} else {

		if respData.Code == 0 {
			respData.Code = 500
		}

		msgBody["code"] = respData.Code

		if respData.Msg != "" {
			msgBody["desc"] = respData.Msg
		} else {
			if respData.Err != nil {
				msgBody["desc"] = respData.Err.Error()
			} else {
				msgBody["desc"] = "Internal Server Error!"
			}
		}

	}

	return msgBody
}

// GetRespDataFromResponse is function to convert response to RespData
func GetRespDataFromResponse(resp *http.Response) (respData data.RespData, err error) {

	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		// Convert body to map
		result := make(map[string]interface{})
		err = json.Unmarshal(body, &result)

		if err == nil {

			dataX := result["data"]
			messageX := result["message"]

			if dataX != nil {
				data := dataX.(map[string]interface{})
				respData.OutData = data
			}

			if messageX != nil {
				message := messageX.(map[string]interface{})
				code := message["code"].(float64)
				desc := message["desc"].(string)
				respData.Code = int(code)
				respData.Msg = desc
				if code == 200 {
					respData.Success = true
				}

			}
		}
	}

	return respData, err
}
