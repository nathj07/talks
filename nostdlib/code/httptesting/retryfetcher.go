package httptesting

import (
	"net/http"
)

func retryFetcher(url string, retryCodes []int, maxRetries int) (int, error) {
	count := 0
	respCode := 0
	for {
		resp, err := http.Get(url)
		if err != nil {
			return 0, err
		}
		respCode = resp.StatusCode
		resp.Body.Close()
		if respCode%200 == 0 {
			return respCode, nil
		}
		for _, rCode := range retryCodes {
			if respCode == rCode {
				count++
			}
		}
		if count == maxRetries {
			return respCode, nil // return the current code
		}
	}
	// return the final respCode ad no error
	return respCode, nil
}
