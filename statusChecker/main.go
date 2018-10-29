package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"
)

/*
--url : URL for the status end point
--status : Status code for the above URL
--checks : Number of checks performed
--interval : Wait between each check
--port : If URL is not given, a connection to tcp should be used for status
*/
var result bool
var validCheck int

func usage() {
	fmt.Printf("--url : URL for the status end point\n" +
		"--host   : Host for tcp connection test\n" +
		"--status : Status code for the above URL\n" +
		"--checks : Number of checks performed\n" +
		"--interval : Wait between each check\n" +
		"--port : If URL is not given, a connection to tcp should be used for status\n" +
		"--healthy : Number of expected healthy checks\n")
}

func debugLogging(debugEnable bool, logString string) {
	if debugEnable == true {
		fmt.Println(logString)
	}
}
func main() {

	checkURL := flag.String("url", "http://localhost", "URL string")
	checkHost := flag.String("host", "localhost", "Value for host")
	checkStatus := flag.Int("status", -1, "Integer value for http return codes")
	checkChecks := flag.Int("checks", -1, "Number of checks")
	checkInterval := flag.Int("interval", -1, "Seconds to wait for each check")
	checkPort := flag.Int("port", -1, "Port number for tcp connection test")
	checkHealthy := flag.Int("healthy", -1, "Number of valid health checks")
	debugFlag := flag.Bool("debug", false, "Enable debug log")
	if len(*checkURL) == 0 && *checkPort == -1 {
		fmt.Println("You must specify URL or Port!")
		usage()
		return
	} else if *checkPort != -1 {
		debugLogging(*debugFlag, "Running TCP port check")
		time.Sleep(time.Duration(*checkInterval) * time.Second)
		_, err := net.Dial("tcp", *checkHost)
		validCheck := 0
		for i := 0; i < *checkChecks; i++ {

			if err != nil {
				fmt.Println("Unable to connect")
			} else {
				validCheck = validCheck + 1
			}

			if validCheck >= *checkHealthy {
				result = true
				break
			}

		}
	} else if checkURL != nil {
		debugLogging(*debugFlag, "Running URL check")
		validCheck = 0
		var netClient = &http.Client{
			Timeout: time.Second * 5,
		}
		for i := 0; i < *checkChecks; i++ {

			response, err := netClient.Get(*checkURL)
			if err != nil {

			}
			if response.StatusCode == *checkStatus {
				validCheck = validCheck + 1
			}
		}
	}

	if !result {
		fmt.Println("Connection: Failed")
		os.Exit(1)
	} else {
		fmt.Println("Connection: Success")
		os.Exit(0)
	}

}
