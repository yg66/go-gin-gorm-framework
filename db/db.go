package db

import (
	"errors"
	"github.com/yg66/go-gin-gorm-framework/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type MyDb struct {
	MyDbConfig *model.DbConfig
	Db         *gorm.DB
}

func Init(dbConfig *model.DbConfig) (*MyDb, error) {
	db := MyDb{
		MyDbConfig: dbConfig,
	}
	// ---- Init Db ----
	gormDb, err := gorm.Open(mysql.Open(dbConfig.MysqlDns), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return &db, err
	}
	db.Db = gormDb
	return &db, nil
}

type DatabaseApi interface {
	FindCacheByKey(cacheKey string) (*model.Cache, error)
	AddCache(cache *model.Cache) (bool, error)
	UpdateCache(cache *model.Cache) (bool, error)
}

func (m *MyDb) FindCacheByKey(cacheKey string) (*model.Cache, error) {
	var cache model.Cache
	tx := m.Db.Where(&model.Cache{CacheKey: cacheKey}).First(&cache)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, tx.Error
	}
	return &cache, nil
}

func (m *MyDb) AddCache(tx *gorm.DB, cache *model.Cache) (bool, error) {
	tx = tx.Create(&cache)
	if tx.Error != nil {
		return false, tx.Error
	}
	if tx.RowsAffected < 1 {
		return false, nil
	}
	return true, nil
}

func (m *MyDb) UpdateCache(tx *gorm.DB, cache *model.Cache) (bool, error) {
	u := model.Cache{
		CacheValue: cache.CacheValue,
		Expired:    cache.Expired,
	}
	tx = tx.Model(cache).Updates(u)
	if tx.Error != nil {
		return false, tx.Error
	}
	if tx.RowsAffected < 1 {
		return false, nil
	}
	return true, nil
}
