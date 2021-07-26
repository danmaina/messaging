package models

type Email struct {
	Pass string `json:"pass"`
	From string `json:"from"`
	Host string `json:"host"`
	Port string `json:"port"`
}
