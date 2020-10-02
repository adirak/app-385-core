// Package activity contains an apis for activity data
package activity

import (
	"fmt"
	"log"
	"strings"

	"github.com/adirak/app-385-core/data"
	"github.com/adirak/app-385-core/dbexec"
)

// calulateEventProgress is function to calulate EventProgross by activity
func calulateEventProgress(actData map[string]interface{}, history data.StravaActivityHistory, token data.StravaUserToken) (updateCount int, err error) {

	log.Println("**calulateEventProgress Start")
	defer log.Println("**calulateEventProgress End")

	activities, ok := actData["activities"].([]map[string]interface{})
	updateCount = 0
	if ok {

		hText := history.History

		// Looping
		for _, act := range activities {

			ID := act["id"].(float64)
			intID := int(ID)

			// Checking new ID with history because user cannot add old ID of activity
			if !strings.Contains(hText, fmt.Sprint(intID)) {

				distance := act["distance"].(float64)
				elapsedTime := act["elapsed_time"].(float64)
				//manual := act["manual"].(bool)

				validDistance := getValidDistance(distance)
				validElapsedTime := getValidElapsedTime(elapsedTime)

				// Load Event Progress from event_progress table
				listEventProg, err := dbexec.LoadEventProgress(token.UserID)
				var errUpdate error
				if err == nil && listEventProg != nil {

					for _, eventProg := range listEventProg {

						// Ignore if event progress is not active
						if eventProg.Active {

							// Increase data in event_progress
							eventProg.Distance = eventProg.Distance + distance
							eventProg.ElapsedTime = eventProg.ElapsedTime + elapsedTime
							eventProg.ValidDistance = eventProg.ValidDistance + validDistance
							eventProg.ValidElapsedTime = eventProg.ValidElapsedTime + validElapsedTime

							// Update eventProgress to database
							errUpdate = dbexec.UpdateEventProgress(eventProg)
							if errUpdate != nil {
								log.Println("Cannot update event_progress user_id=", eventProg.UserID, " event_id=", eventProg.EventID)
								break
							}
						}

					}

					if errUpdate == nil {
						hText = fmt.Sprint(intID) + "," + hText
					} else {
						updateCount++
					}
				}

			} else {
				updateCount++
			}
		}

		// Length limit
		if len(hText) > 300 {
			hText = hText[:150]
		}

		// Update Strava Activity History
		history.History = hText
		if history.HasHistory {
			err := dbexec.UpdateStravaActivityHistory(history)
			if err != nil {
				log.Println("updateStravaActivityHistory fail at user_id=", history.UserID)
			}
		} else {
			err := dbexec.InsertStravaActivityHistory(history)
			if err != nil {
				log.Println("insertStravaActivityHistory fail at user_id=", history.UserID)
			}
		}

	} else {
		log.Println("Cannot cast activities to array : ", actData)
	}

	return
}

// getValidDistance to get valid distance by rule
func getValidDistance(distance float64) float64 {
	return distance
}

// getValidElapsedTime to get elapsed time by rule
func getValidElapsedTime(elapsedTime float64) float64 {
	return elapsedTime
}
