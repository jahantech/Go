package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

const CIDR_RANGE string = "192.168.0"


struct Hosts{
	IpAddress string,
	Status string,
	OpenPorts []string
}
func main() {

	var wg sync.WaitGroup

	fmt.Println("Starting CIDR Scanner...")

	wg.Add(255)
	for i := 0; i < 255; i++ {

		address := CIDR_RANGE + "." + strconv.Itoa(i)
		go func(address string) {
			defer wg.Done()
			if checkRemotePorts(address) {
				fmt.Println(address + " is online")
			}
		}(address)
	}

	fmt.Println("Waiting To Finish")
	wg.Wait()

	return
}

func checkRemotePorts(ipAddress string) bool {

	CheckOnlinePortList := []string{":80", ":443", ":22", ":389", ":636", ":31813"}
	for _, v := range CheckOnlinePortList {
		conn, err := net.DialTimeout("tcp", ipAddress+v, time.Millisecond*50)
		if err == nil {
			fmt.Println(ipAddress + " open port " + v)
			defer conn.Close()
			return true
		}
	}
	return false
}
