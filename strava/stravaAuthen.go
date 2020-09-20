// Package strava contains an strava connection
package strava

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"time"

	"github.com/adirak/app-385-core/data"
	"github.com/adirak/app-385-core/myhttp"
)

// RefreshStravaToken is function to request token from strava
func RefreshStravaToken(token data.StravaUserToken, app data.StravaApp) (result data.RefreshTokenResult, err error) {

	// Result
	result = data.RefreshTokenResult{}
	result.UserToken = token

	// URL config
	refreshURL := app.RefreshURL

	// From Data
	formData := url.Values{}
	formData.Set("client_id", fmt.Sprint(app.ClientID))
	formData.Set("client_secret", app.ClientSecret)
	formData.Set("refresh_token", token.RefreshToken)
	formData.Set("grant_type", "refresh_token")

	log.Println("**Try to refreshToken", "\nclient_id=", fmt.Sprint(app.ClientID), "\nclient_secret=", app.ClientSecret, "\nrefresh_token=", token.RefreshToken, "\nrefreshURL=", refreshURL)

	// Make request
	req, err := myhttp.NewPostRequest(refreshURL, formData)
	if err == nil {

		// Set Header
		myhttp.SetAcceptJSONHeader(req)

		// Create client
		client := myhttp.GetMyClient()

		// Do request
		resp, err1 := client.Do(req)

		// Check Request error
		if err1 != nil {
			return result, err1
		}
		defer resp.Body.Close()

		// Read body data
		body, err2 := ioutil.ReadAll(resp.Body)
		if err2 != nil {
			return result, err2
		}

		// Convert body to map
		respData := make(map[string]interface{})
		err = json.Unmarshal(body, &respData)

		// This token has been expired
		if err == nil {

			result.Success = true

			if respData["errors"] != nil {
				result.TokenExpired = true
			} else {

				accessToken := respData["access_token"].(string)
				refreshToken := respData["refresh_token"].(string)
				expiredAt := respData["expires_at"].(float64)
				expiredIn := respData["expires_in"].(float64)
				expiredAtTime := time.Unix(int64(expiredAt), 0)

				result.AccessToken = accessToken
				result.RefreshToken = refreshToken
				result.ExpiredAt = expiredAtTime
				result.ExpiredIn = int64(expiredIn)
				result.TokenExpired = false
			}
		}

	}

	return result, err
}
