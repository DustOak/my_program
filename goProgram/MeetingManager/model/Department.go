package model

type Department struct {
	ID             int    `gorm:"AUTO_INCREMENT;PRIMARY_KEY;column:ID"`
	DepartmentName string `gorm:"type:varchar(50);not null;column:DEPARTMENT_NAME"`
}

func (Department) TableName() string {
	return "DEPARTMENT"
}

func (d *Department) GetSlice() interface{} {
	slice := make([]Department, 0)
	return &slice
}
