package main

import "fmt"
import "time"

func main() {

	maxroutines := make(chan int, 5)
	for i := 0; i < 20; i++ {
		maxroutines <- i
		go func(j int) {
			fmt.Println(j)
			time.Sleep(1000 * time.Millisecond)

		}(i)
		<- maxroutines
	}

	return
}
