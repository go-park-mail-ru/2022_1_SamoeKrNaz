package models

type List struct {
	IdL      uint   `json:"idl"`
	Title    string `json:"title"`
	Position uint   `json:"position"`
	IdB      uint   `json:"idb"`
	Tasks    []Task
}
