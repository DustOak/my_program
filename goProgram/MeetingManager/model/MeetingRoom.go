package model

type MeetingRoom struct {
	ID           int    `gorm:"PRIMARY_KEY;AUTO_INCREMENT;column:ID"`
	RoomID       string `gorm:"type:varchar(20);not null;column:ROOM_ID"`
	RoomName     string `gorm:"type:varchar(50);not null;column:ROOM_NAME"`
	RoomCapacity int    `gorm:"not null;column:ROOM_CAPACITY"`
	RoomStatus   int    `gorm:"type:int;not null;default:1;column:ROOM_STATUS"`
	RoomRemarks  string `gorm:"column:ROOM_REMARKS"`
}

func (MeetingRoom) TableName() string {
	return "MEETINGROOM"
}
func (d *MeetingRoom) GetSlice() interface{} {
	slice := make([]MeetingRoom, 0)
	return &slice
}
