/*
this package is used for checking the latency of the sites
*/

package pingsites

import (
	"fmt"
	"net/http"
	"time"
)

func GetLatency(url string) (string, string, error) {

	fmt.Println("sending get request.......")
	startTime := time.Now()
	response, err := http.Get(url)
	ping := time.Since(startTime)
	if err != nil {
		return ping.String(), response.Status, err
	}

	fmt.Println("time taken: ", ping)
	defer response.Body.Close()

	//	return string(readableString), nil
	return ping.String(), response.Status, nil
}
