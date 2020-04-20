package delayed

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

func DelayedResponse(w http.ResponseWriter, r *http.Request) {

	maxDelay := os.Getenv("MAX_DELAY")
	maxDelayInt, err := strconv.Atoi(maxDelay)
	if err != nil {
		panic("Invalid Max delay value")
	}

	selectedDelay := rand.Intn(maxDelayInt)
	fmt.Println(selectedDelay)
	time.Sleep(time.Second * time.Duration(selectedDelay))
	fmt.Fprintf(w, "Done")
}
