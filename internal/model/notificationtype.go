package model

import (
	"github.com/jinzhu/gorm"
	"errors"
	"fmt"
)

type NotificationType struct {
	Id          uint   `gorm:"primary_key;column:id" json:"id,omitempty"`
	SystemKey   string `gorm:"size:25";column:"system_key" json:"systemKey,omitempty"`
	Version     string `gorm:"size:10";column:"version" json:"version,omitempty"`
	MaxGrouping uint   `gorm:"size:25";column:"max_grouping"  json:"maxGrouping,omitempty"`
}

func CreateNotificationType(sk string, v string) NotificationType {
	return NotificationType{
		SystemKey: sk,
		Version: v,
		MaxGrouping: 10,
	}
}

func (nt *NotificationType) Save(db *gorm.DB) {
	db.Save(nt)
}


func (nt *NotificationType) GetById(db *gorm.DB, id int) error {
	if db.First(nt, id).RecordNotFound() {
		return errors.New(fmt.Sprint("Notification Type not found with id: %d", id))
	}

	return nil
}