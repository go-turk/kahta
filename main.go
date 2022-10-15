package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/merhaba", GetMerhaba).Methods("POST")
	router.HandleFunc("/saat-kac", SaatKac)
	server := http.Server{
		Addr:    ":1525",
		Handler: router,
	}
	fmt.Println("Sunucu 1525 Portunda ayağa kalktı")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}

func GetMerhaba(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Merhaba Yazdık Goründü mü?")
	writer.Write([]byte("Alameddin Hoş Geldin"))
}

func SaatKac(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(time.Now().String()))
}
