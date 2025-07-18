/*
this package is used for checking the latency of the sites
*/

package pingsites

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func GetLatency(url string) (time.Duration, string, error) {
	fmt.Println("sending get request.......")
	startTime := time.Now()
	response, err := http.Get(url)
	totalTime := time.Since(startTime)
	if err != nil {
		return totalTime, "", err
	}
	readableString, err := io.ReadAll(response.Body)
	if err != nil {
		return totalTime, "", err
	}
	fmt.Println("time taken: ", totalTime)
	defer response.Body.Close()

	//	return string(readableString), nil
	return totalTime, string(readableString), nil
}
