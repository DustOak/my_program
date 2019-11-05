package model

type Comment struct {
	ID            int `gorm :"AUTO_INCREMENT"`
	Name          string
	Content       string
	Email         string
	Url           string
	HeadImagePath string `gorm:"column:headImagePath"`
	Date          string
	IpAddress     string `gorm:"column:ipAddress"`
}

func (Comment) TableName() string {
	return "comment"
}

func (Comment) GetClass() interface{} {
	return &Comment{}
}

func (Comment) GetSliceClass() interface{} {
	slice := make([]Comment, 0)
	return &slice
}
