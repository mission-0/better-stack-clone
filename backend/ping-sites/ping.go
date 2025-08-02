/*
this package is used for checking the latency of the sites
*/

package pingsites

import (
	"fmt"
	"net/http"
	"time"
)

func GetLatency(url string) (string, string, string) {

	fmt.Println("sending get request.......")
	startTime := time.Now()
	response, err := http.Get(url)
	ping := time.Since(startTime)

	if response != nil {
		defer response.Body.Close()
		if err != nil {
			fmt.Println("error is :", err.Error())
			return ping.String(), response.Status, err.Error()
		}

		fmt.Println("time taken: ", ping)
		

		//	return string(readableString), nil
		return ping.String(), response.Status, "null"

	}

	return ping.String(), "500 " + http.StatusText(http.StatusInternalServerError), err.Error()
}
