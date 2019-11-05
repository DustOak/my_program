package model

type Personnel struct {
	ID                int        `gorm:"AUTO_INCREMENT;PRIMARY_KEY;column:ID"`
	PersonnelName     string     `gorm:"type:varchar(50);not null;column:PERSONNEL_NAME"`
	PersonnelAccount  string     `gorm:"type:varchar(50);not null;column:PERSONNEL_ACCOUNT"`
	PersonnelPassword string     `gorm:"type:varchar(50);not null;column:PERSONNEL_PASSWORD"`
	PersonnelEmail    string     `gorm:"type:varchar(50);not null;column:PERSONNEL_EMAIL"`
	Department        Department `gorm:"ForeignKey:DepartmentID;PRELOAD:false"`
	DepartmentID      int        `gorm:"column:DEPARTMENT_ID"`
	PersonnelStatus   int        `gorm:"type:int;not null;default:1;column:PERSONNEL_STATUS"`
	PersonnelPhone    string     `gorm:"type:varchar(20);not null;column:PERSONNEL_PHONE"`
}

func (Personnel) TableName() string {
	return "PERSONNEL"
}
func (d *Personnel) GetSlice() interface{} {
	slice := make([]Personnel, 0)
	return &slice
}
