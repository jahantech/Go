package main

import "fmt"
import "time"

func main() {

	maxroutines := make(chan int, 5)
	for i := 0; i < 20; i++ {
		maxroutines <- i
		go func() {
			fmt.Println(i)
			time.Sleep(1000 * time.Millisecond)

		}()
		<- maxroutines
	}

	return
}
