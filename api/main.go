package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/StefanoPernat/GE/api/server"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.GET("/top", server.TodayTop)
	log.Printf("server listening on 127.0.0.1:%d...\n", server.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", server.Port), router))
}
