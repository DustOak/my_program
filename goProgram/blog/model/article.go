package model

type Article struct {
	ID         int `gorm :"AUTO_INCREMENT"`
	Title      string
	Date       string
	Content    string
	Category   Category
	CategoryId int
}

func (Article) TableName() string {
	return "article"
}

func (Article) GetClass() interface{} {
	return &Article{}
}

func (Article) GetSliceClass() interface{} {
	slice := make([]Article, 0)
	return &slice
}
