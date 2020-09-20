// Package strava contains an strava connection
package strava

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/adirak/app-385-core/data"
	"github.com/adirak/app-385-core/myhttp"
)

// Max Retry
var maxRetry = 1

// LoadStravaActivity is function to load strava activity from strava web
func LoadStravaActivity(token data.StravaUserToken, app data.StravaApp) (result map[string]interface{}, rToken data.RefreshTokenResult, err error) {

	fmt.Println("LoadStravaActivity Start")
	defer fmt.Println("LoadStravaActivity End")

	return doLoadStravaActivity(token, app, 0, false)
}

// doLoadStravaActivity is function to load strava activity from strava web
func doLoadStravaActivity(token data.StravaUserToken, app data.StravaApp, reTry int, fourceRefresh bool) (result map[string]interface{}, rToken data.RefreshTokenResult, err error) {

	// Make Refresh Token Data
	rToken = data.RefreshTokenResult{}

	// Check expired time
	curTime := time.Now()
	oldTime := token.ExpiredAt

	if (curTime.After(oldTime) || fourceRefresh) && reTry <= maxRetry {
		rToken, err = RefreshStravaToken(token, app)
	}

	log.Println("** accessToken : ", rToken.AccessToken, ", tokenExpired : ", rToken.TokenExpired)

	// Do get activity if token is not expired
	if err == nil && rToken.TokenExpired == false {

		// URL config
		stravaActURL := app.LoadActivityURL

		// Make request
		req, err2 := myhttp.NewGetRequest(stravaActURL)
		if err2 == nil {

			// Set Header
			myhttp.SetAcceptJSONHeader(req)
			myhttp.SetTokenHeader(req, rToken.AccessToken)

			// Create client
			client := myhttp.GetMyClient()

			// Do call request
			resp, err3 := client.Do(req)

			// Check request error
			if err3 != nil {
				return nil, rToken, err3
			}
			defer resp.Body.Close()

			// Get response body
			body, err4 := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, rToken, err4
			}

			// Check the array of response body
			isArr := myhttp.IsJSONArray(body)
			if isArr {

				activities := make([]map[string]interface{}, 0, 5)
				err = json.Unmarshal(body, &activities)
				result = make(map[string]interface{})
				result["activities"] = activities

			} else {

				respMap := make(map[string]interface{})
				err = json.Unmarshal(body, &respMap)
				result = respMap

			}

			// This token has been expired
			if err == nil && result["errors"] != nil {

				// do retry
				reTry++
				result, rToken, err = doLoadStravaActivity(token, app, reTry, true)
			}

			return result, rToken, err
		}

		// err2 is not nit it is error
		if err2 != nil {
			err = err2
		}

	} else {
		log.Println("Refresh token fail, tokenExpired=", fmt.Sprint(rToken.TokenExpired), " user_id=", token.UserID, " --> ", err.Error())
	}

	return result, rToken, err
}
