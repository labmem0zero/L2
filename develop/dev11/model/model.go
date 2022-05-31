package model

import "time"

type LoggedRequest struct {
	Request string
	Client string
	Result string
}

type Event struct{
	Id int `json:"id,omitempty"`
	Date string	`json:"date"`
	Time string	`json:"time,omitempty"`
	Name string	`json:"name"`
}

type Day struct {
	Date time.Time
	Events []Event
}