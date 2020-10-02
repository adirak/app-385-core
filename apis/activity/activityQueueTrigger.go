// Package activity contains an apis for activity data
package activity

import (
	"log"

	"github.com/adirak/app-385-core/data"
	"github.com/adirak/app-385-core/dbexec"
	"github.com/adirak/app-385-core/strava"
)

// TriggerActivityQueue is api function to triger strava query from its api
func TriggerActivityQueue(reqt data.ReqtData) (resp data.RespData) {

	log.Println("TriggerActivityQueue Start")
	defer log.Println("TriggerActivityQueue End")

	// Output data
	outData := make(map[string]interface{})

	// Get Data from activity queue
	queue, err := dbexec.LoadActivityQueue()

	if err == nil {

		// Loop Get Activity
		loopToGetStravaActivity(queue)

		// Looping success
		resp.Success = true
	}

	// Footer response
	resp.OutData = outData
	resp.Err = err
	return resp
}

// loopToGetStravaActivity is function loop to get Strava Activity
func loopToGetStravaActivity(queue []data.ActivityQueue) {

	// Log
	log.Println("**loopToGetStravaActivity Start")
	defer log.Println("**loopToGetStravaActivity End")

	// Looping
	for _, actQ := range queue {

		if actQ.Active {

			// Query Strava User Token
			token, err := dbexec.LoadStravaUserToken(actQ.StravaUserID)
			if err == nil && token.AppID > 0 {

				// Query Strava App
				app, err := dbexec.LoadStravaApp(token.AppID)
				if err == nil {

					// Call to request activity from Strava web
					result, rToken, err := strava.LoadStravaActivity(token, app)
					if err == nil && rToken.TokenExpired == false {

						// Query Strava Activity History
						history, err := dbexec.LoadStravaActivityHistory(token.UserID)

						if err == nil {

							updateCount, err := calulateEventProgress(result, history, token)

							if err == nil && updateCount > 0 {
								dbexec.UpdateActivityQueue(actQ)
							}

						} else {
							log.Println("cannot loadStravaActivityHistory by user_id : ", token.UserID, " --> ", err.Error())
						}

					} else {

						// To update Strava User Token by active = false
						err = dbexec.UpdateTokenExpired(token)
						if err != nil {
							log.Println("cannot updateTokenExpired by user_id : ", token.UserID, " --> ", err.Error())
						}

					}

				} else {
					log.Println("cannot loadStravaApp by app_id : ", token.AppID, " --> ", err.Error())
				}

			} else {
				log.Println("cannot loadStravaUserToken by strava_user_id : ", actQ.StravaUserID, " --> ", err.Error())
			}

		} else {
			log.Println("strava_user_id : ", actQ.StravaUserID, " is not active")
		}
	}

}
