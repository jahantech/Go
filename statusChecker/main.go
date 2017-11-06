package main

import "flag"
import "fmt"

/*
--url : URL for the status end point
--status : Status code for the above URL
--checks : Number of checks performed
--interval : Wait between each check
--port : If URL is not given, a connection to tcp should be used for status
*/
func main() {

	checkURL := flag.String("url", "http://localhost", "URL string")
	checkStatus := flag.Int("status", 200, "Integer value for http return codes")
	checkChecks := flag.Int("checks", 1, "Number of checks")
	checkInterval := flag.Int("interval", 1, "Seconds to wait for each check")
	checkPort := flag.Int("port", 80, "Port number for tcp connection test")
	checkHealthy := flag.Int("healthy", 1, "Number of valid health checks")

	if checkURL != nil && checkPort != nil {
		fmt.Println("You must specify URL or Port!")
		return
	} else if checkPort != nil {

	} else if checkURL != nil {

	}

	return
}
