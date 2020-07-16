package model

type EventType int8

const (
	EventNormal = iota
	EventBirth
	EventDeath
)

type Event struct {
	ID 			int64
	Class	 	EventType
	Year		string
	Date 	 	string
	Detail	 	string
	Links 	 	string
	ImgLinks 	string
}