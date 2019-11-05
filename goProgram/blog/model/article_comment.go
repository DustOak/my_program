package model

type ArticleComment struct {
	ID            int `gorm :"AUTO_INCREMENT"`
	ArticleId     int `gorm:"article_id"`
	Name          string
	Content       string
	Email         string
	Url           string
	Date          string
	HeadImagePath string `gorm:"column:headImagePath"`
	IpAddress     string `gorm:"column:ipAddress"`
}

func (ArticleComment) TableName() string {
	return "articleComment"
}

func (ArticleComment) GetClass() interface{} {
	return &ArticleComment{}
}

func (ArticleComment) GetSliceClass() interface{} {
	slice := make([]ArticleComment, 0)
	return &slice
}
