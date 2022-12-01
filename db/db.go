package db

import (
	"errors"
	"fmt"
	"strings"

	"github.com/yg66/go-gin-gorm-framework/config"
	"github.com/yg66/go-gin-gorm-framework/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Db struct {
	IDB *gorm.DB
}

func NewDB(iDBConfig *config.DbConfig) (idb *Db, err error) {
	var logLevel logger.LogLevel
	switch strings.ToLower(iDBConfig.LogLevel) {
	case "error":
		logLevel = logger.Error
	case "warn":
		logLevel = logger.Warn
	case "info":
		logLevel = logger.Info
	default:
		err = fmt.Errorf("%w", errors.New("db log level is incorrect"))
		return
	}

	gormDb, err := gorm.Open(mysql.Open(iDBConfig.MysqlDns), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return
	}
	idb = &Db{IDB: gormDb}
	return
}

type IDB interface {
	FindCacheByKey(cacheKey string) (*model.Cache, error)
	AddCache(tx *gorm.DB, cache *model.Cache) (bool, error)
	UpdateCache(tx *gorm.DB, cache *model.Cache) (bool, error)
}

func (m *Db) FindCacheByKey(cacheKey string) (cache *model.Cache, err error) {
	tx := m.IDB.Where(&model.Cache{CacheKey: cacheKey}).First(&cache)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		err = tx.Error
	}
	return
}

func (m *Db) AddCache(tx *gorm.DB, cache *model.Cache) (ok bool, err error) {
	tx = tx.Create(&cache)
	if tx.Error != nil {
		err = tx.Error
		return
	}
	if tx.RowsAffected < 1 {
		return
	}
	return true, nil
}

func (m *Db) UpdateCache(tx *gorm.DB, cache *model.Cache) (bool, error) {
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
