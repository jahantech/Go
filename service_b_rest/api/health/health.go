package health

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func Health(w http.ResponseWriter, r *http.Request) {

	maxDelay := os.Getenv("MAX_DELAY")
	_, err := strconv.Atoi(maxDelay)
	if err == nil {
		fmt.Fprintf(w, "ok")
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "UnHealthy")

}
