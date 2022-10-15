package main

import (
	"fmt"
	"github.com/go-turk/kahta/otobus"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/otobusler", otobus.GetOtobusler)
	apiRouter.HandleFunc("/otobusler/{id:[0-9]}", otobus.GetOtobus)
	server := http.Server{
		Addr:    ":1525",
		Handler: router,
	}
	fmt.Println("Sunucu 1525 Portunda ayağa kalktı")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
