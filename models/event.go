package models

//easyjson:json
type Event struct {
	EventType string `json:"event_type"`
	IdB       uint   `json:"id_b"`
	IdT       uint   `json:"id_t"`
}
