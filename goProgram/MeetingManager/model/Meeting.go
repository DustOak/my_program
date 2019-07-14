package model

import (
	"time"
)

type Meeting struct {
	ID                 int         `gorm:"AUTO_INCREMENT;PRIMARY_KEY;column:ID"`
	Name               string      `gorm:"type:varchar(100);not null;column:NAME"`
	MeetingRoom        MeetingRoom `gorm:"ForeignKey:RID"`
	RID                int         `gorm:"not null;column:ROOM_ID"`
	StartTime          time.Time   `gorm:"type:date;not null;column:START_TIME"`
	EndTime            time.Time   `gorm:"type:date;not null;column:END_TIME"`
	BookTime           time.Time   `gorm:"type:date;not null;column:BOOK_TIME"`
	BookNumber         int         `gorm:"not null;column:BOOK_NUMBER"`
	MeetingDescription string      `gorm:"column:MEETING_DESCRIPTION"`
	MeetingStatus      int         `gorm:"type:int;not null;default:1;column:MEETING_STATUS"`
	BP                 Personnel   `gorm:"ForeignKey:BookPersonnel"`
	BookPersonnel      int         `gorm:"column:BOOK_PERSONNEL"`
}

func (Meeting) TableName() string {
	return "MEETING"
}
func (d *Meeting) GetSlice() interface{} {
	slice := make([]Meeting, 0)
	return &slice
}
