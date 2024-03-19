package models

type Node struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Addr     string `json:"addr"`
	Password string `json:"password"`
	Key      string `json:"key"`
	Jump1    int    `json:"jump1_id"`
	Jump2    int    `json:"jump2_id"`
}
