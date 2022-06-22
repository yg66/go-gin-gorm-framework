package service

import (
	"github.com/yg66/go-gin-gorm-framework/common/errors"
	"github.com/yg66/go-gin-gorm-framework/db"
	"github.com/yg66/go-gin-gorm-framework/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

type Service struct {
	D *db.MyDb
}

type CacheServiceApi interface {
	TestDbTransaction() error
}

func (s *Service) TestDbTransaction() error {
	// open transaction
	tx := s.D.Db.Begin()
	var (
		ok  bool
		err error
	)
	defer func() {
		// 捕获到异常，执行回滚
		if err != nil {
			zap.S().Error(err)
			tx.Rollback()
		}
		if !ok {
			tx.Rollback()
		}
	}()
	// database operations
	//cache, _ := s.D.FindCacheByKey("authorization_640c9b6eb58c11eca67a000c2933e60c")
	cache, err := s.D.FindCacheByKey("authorization_640c9b6eb58c11eca67a000c2933e6")
	if err != nil {
		return err
	}
	if cache == nil {
		return errors.New(errors.DataNotFound)
	}
	cache.CacheValue = "更新字段222"
	ok, err = s.D.UpdateCache(tx, cache)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New(errors.ServerError)
	}

	c := &model.Cache{
		Model: gorm.Model{
			ID: 144,
		},
		CacheKey:   "测试键",
		CacheValue: "测试值",
		Expired:    time.Time{}.UTC().Add(24 * time.Hour),
	}
	ok, err = s.D.AddCache(tx, c)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New(errors.ServerError)
	}

	// commit transaction
	return tx.Commit().Error
}
