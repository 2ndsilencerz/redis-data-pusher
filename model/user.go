package model

type User struct {
	Nama      string `json:"nama"`
	Umur      int    `json:"umur"`
	RedisPush bool
}
