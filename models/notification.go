package models

import "time"

//easyjson:json
type Notifications []Notification

//easyjson:json
type Notification struct {
	IdU              uint      `json:"idu"`                                //кому адресовано уведомление
	NotificationType string    `json:"notification_type" gorm:"not null;"` //тип уведомления
	Date             string    `json:"date" gorm:"not null;"`              //дата адекватная
	DateToOrder      time.Time `gorm:"not null;"`                          //дата для сортировки
	IsRead           bool      `json:"is_read"`
	IdB              uint      `json:"idb"`
	IdT              uint      `json:"idt"`
	IdWh             uint      `json:"id_wh"`
	Board            Board     `json:"board" gorm:"-"`    //отправим доску, если пришло уведомление приглашения на доску
	Task             Task      `json:"task" gorm:"-"`     //отправим таску, если пришло уведомление приглашения на таску
	UserWho          User      `json:"user_who" gorm:"-"` //тот, кто отправил это уведомление
}
