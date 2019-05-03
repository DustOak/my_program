package model

import (
	"blog/logger"
	"github.com/jinzhu/gorm-master"
	"sync"

	_ "github.com/jinzhu/gorm-master/dialects/mysql"
)

type DB struct {
	db   *gorm.DB
	lock sync.Mutex
}

var db DB

func init() {
	logger.InfoLog.Println("Connect Database ....")
	var err error
	db.db, err = gorm.Open("mysql", "root:cuiwenbin1397.+@/blog?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		logger.ErrLog.Println(err)
		panic(err)
	}
	db.db.DB().SetMaxOpenConns(2000)
	db.db.DB().SetMaxIdleConns(1500)
	db.db.DB().Ping()
}

func SelectAll(object BaseObject, preload string) interface{} {
	db.lock.Lock()
	defer db.lock.Unlock()
	value := object.GetSliceClass()
	if preload == "" {
		db.db.Find(value)
	} else {
		db.db.Preload(preload).Order("id desc ").Find(value)
	}

	return value
}

func Select(object BaseObject, id int, preload ...string) interface{} {
	db.lock.Lock()
	defer db.lock.Unlock()
	if len(preload) == 0 {
		value := object.GetClass()
		db.db.First(value, id)
		return value
	} else {
		value := object.GetClass()
		db.db.Preload(preload[0]).First(value, id)
		return value
	}
}

func Save(object BaseObject) error {
	db.lock.Lock()
	defer db.lock.Unlock()
	tx := db.db.Begin()
	if err := tx.Create(object).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func SelectPreload(object BaseObject, preload, field, condition string, value ...interface{}) interface{} {
	db.lock.Lock()
	defer db.lock.Unlock()
	values := object.GetSliceClass()
	db.db.Preload(preload).Select(field).Where(condition, value...).Find(values)
	return values
}

func SelectWhere(object BaseObject, condition, field string, value ...interface{}) interface{} {
	db.lock.Lock()
	defer db.lock.Unlock()
	values := object.GetSliceClass()
	db.db.Select(field).Where(condition, value...).Find(values)
	return values
}

func Close() {
	db.db.Close()
}

func Update(object BaseObject) error {
	db.lock.Lock()
	defer db.lock.Unlock()
	tx := db.db.Begin()
	if err := tx.Model(object).Updates(object).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func Delete(object BaseObject) error {
	db.lock.Lock()
	defer db.lock.Unlock()
	tx := db.db.Begin()
	if err := tx.Delete(object).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
