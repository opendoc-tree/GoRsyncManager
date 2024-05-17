package models

type Index struct {
	Id   int    `json:"id"`
	Path string `json:"path"`
	Page int    `json:"page"`
}
