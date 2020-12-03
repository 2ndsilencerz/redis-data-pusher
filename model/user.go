package model

// User is a dummy model to load data from database
type User struct {
	Nama      string `json:"nama"`
	Umur      int    `json:"umur"`
	RedisPush bool
}
