package model

type MeetingPersonnel struct {
	Meeting     Meeting   `gorm:"ForeignKey:MeetingID"`
	MeetingID   int       `gorm:"column:MEETING_ID"`
	Personnel   Personnel `gorm:"ForeignKey:PersonnelID"`
	PersonnelID int       `gorm:"column:PERSONNEL_ID"`
}

func (MeetingPersonnel) TableName() string {
	return "MEETING_PERSONNEL"
}
func (d *MeetingPersonnel) GetSlice() interface{} {
	slice := make([]MeetingPersonnel, 0)
	return &slice
}
