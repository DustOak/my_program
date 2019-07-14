package database

import (
	"MeetingManager/model"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:950712cwb@/MEETINGMANAGER?charset=utf8&parseTime=True&loc=Local")
	defer func() {
		if err := recover(); err != nil {
			log.Fatalln("数据库初始化异常,已被recover,错误信息:[", err.(error).Error(), "],程序退出")
		}
	}()
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

// 通过id 获取一个对象
func QueryObject(isPreload bool, object interface{}, id int) interface{} {
	db.Set("gorm:auto_preload", isPreload).First(object, id)
	return object
}

// 通过id 获取一个对象
func QueryObjectByPreload(object interface{}, id int) interface{} {
	db.Set("gorm:auto_preload", true).First(object, id)
	return object
}

// 通过条件获取对象
func QueryByCondition(object model.BaseInterface) interface{} {
	db.Where(object).Find(object)
	return object
}
func QueryPreloadCondition(object model.BaseInterface, condition string, value ...interface{}) interface{} {
	a := object.GetSlice()
	db.Set("gorm:auto_preload", true).Where(condition, value...).Find(a)
	return a
}

//查询全部
func QueryAll(object model.BaseInterface) interface{} {
	a := object.GetSlice()
	db.Where(object).Find(a)
	return a
}

// 原生sql 查询
func Query(object model.BaseInterface, sql string, value ...interface{}) interface{} {
	a := object.GetSlice()
	db.Raw(sql, value).Scan(a)
	return a
}

//原生sql执行
func Exec(sql string, value ...interface{}) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	if err := tx.Exec(sql, value).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error

}

// 删除
func Delete(object interface{}) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	if err := tx.Delete(object).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// 插入 返回true 则保存成功
func Insert(object interface{}) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	if ok := tx.NewRecord(object); ok {
		if err := tx.Create(object).Error; err != nil {
			tx.Rollback()
			return err
		}
	} else {
		return errors.New("已有相同主键")
	}
	return tx.Commit().Error
}

// 更新
func Update(object interface{}) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	if err := tx.Model(object).Updates(object).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
