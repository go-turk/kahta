package otobus

import (
	"encoding/json"
	"net/http"
)

func GetOtobusler(w http.ResponseWriter, r *http.Request) {
	// aktif port : 8080
	// Get Otobüsler için 5 tane otobüsü bir array'e tanımlayıp aynı şekilde marshallayıp
	// ekrana otobusler endpointi çalıştığında göstermemiz
}

func GetOtobus(w http.ResponseWriter, r *http.Request) {
	ilkOtobus := NewOtobus(1, "34 1234", 40)
	JsonOtobus, err := json.Marshal(ilkOtobus)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Otobüs Tanımlaması yapılamadı"))
		return
	}
	w.Write(JsonOtobus)
}
