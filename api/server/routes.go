package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/StefanoPernat/GE/api/reddit"
	"github.com/julienschmidt/httprouter"
)

// TodayTop is the main server route that serve the top game deals for today
func TodayTop(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	deals, err := reddit.GetTodayHotDeals()

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		handleError(w, err)
		return
	}

	data, err := json.Marshal(deals)
	if err != nil {
		handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func handleError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "{\"error\": \"%s\"}", err.Error())
}
