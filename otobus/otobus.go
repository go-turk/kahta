package otobus

type Otobus struct {
	Id           int    `json:"id"`
	Plaka        string `json:"plaka"`
	KoltukSayisi int    `json:"koltuk_sayisi"`
}

func NewOtobus(id int, plaka string, koltukSayisi int) *Otobus {
	return &Otobus{
		Id:           id,
		Plaka:        plaka,
		KoltukSayisi: koltukSayisi,
	}
}
