package models

type Host struct {
	Id       int    `json:"id"`
	Name     string `json:"name" form:"name"`
	Addr     string `json:"addr" form:"addr"`
	User     string `json:"user" form:"user"`
	Password string `json:"password" form:"password"`
	Key      string `json:"key" form:"key"`
	Bastion  int    `json:"bastion" form:"bastion"`
}
