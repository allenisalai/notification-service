package model

import "time"

type Notification struct {
	Id                 int
	NotificationTypeId int
	FirstAdded         time.Time
	LastAdded          time.Time
	Parameters         []interface{}
	Count              uint32
}

