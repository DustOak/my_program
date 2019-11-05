package model

type Category struct {
	ID           int `gorm :"AUTO_INCREMENT"`
	CategoryName string
	Url          string
}

func (Category) TableName() string {
	return "category"
}

func (Category) GetClass() interface{} {
	return &Category{}
}

func (Category) GetSliceClass() interface{} {
	slice := make([]Category, 0)
	return &slice
}
