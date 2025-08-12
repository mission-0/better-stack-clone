package pingsites

import (
	"fmt"
	"net/http"
	"time"
)

func GetLatency(url string) (string, string, string, bool) {

	fmt.Println("sending get request.......")
	startTime := time.Now()
	response, err := http.Get(url)
	ping := time.Since(startTime)

	if response != nil {
		defer response.Body.Close()

		if err != nil {
			fmt.Println("error is :", err.Error())
			return ping.String(), response.Status, err.Error(), false
		}

		fmt.Println("time taken: ", ping)

		//	return string(readableString), nil
		// return ping.String(), response.Status, "null", true
		isSuccess := response.StatusCode >= 200 && response.StatusCode < 300
        
        return ping.String(), response.Status, "null", isSuccess

	}

	return ping.String(), "500 " + http.StatusText(http.StatusInternalServerError), err.Error(), false
}
