// Package trigger contains an http request trigger
package trigger

import (
	"errors"

	"github.com/adirak/app-385-core/util"

	"github.com/adirak/app-385-core/myhttp"
)

// CallActivityQueueTrigger is function to call Activity Queue Trigger
func CallActivityQueueTrigger() error {

	// URL config
	activityQueueURL := "https://asia-southeast2-app-385.cloudfunctions.net/ativity-queue-trigger"

	// Make request
	req, err := myhttp.NewGetRequest(activityQueueURL)
	if err == nil {

		client := myhttp.GetMyClient()
		resp, err2 := client.Do(req)

		if err2 == nil {
			defer resp.Body.Close()

			respData, err3 := util.GetRespDataFromResponse(resp)

			if err3 == nil {

				if respData.Code < 200 || respData.Code >= 300 {
					err = errors.New("Cron job to run activity queue cannot not run yet")
					return err
				}

				return nil

			}

			return err3

		}

		return err2
	}

	return err
}
