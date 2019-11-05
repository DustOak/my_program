package model

type Admin struct {
	ID            int `gorm :"AUTO_INCREMENT"`
	Email         string
	Password      string
	LastTime      string `gorm:"column:lastTime"`
	LastIpAddress string `gorm:"column:lastIpAddress"`
}

func (Admin) TableName() string {
	return "admin"
}

func (Admin) GetClass() interface{} {
	return &Admin{}
}

func (Admin) GetSliceClass() interface{} {
	slice := make([]Admin, 0)
	return &slice
}
