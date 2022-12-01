package service

import (
	"time"

	"github.com/yg66/go-gin-gorm-framework/common/errors"
	"github.com/yg66/go-gin-gorm-framework/db"
	"github.com/yg66/go-gin-gorm-framework/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CacheService struct {
	IDB db.IDB
	TX  *gorm.DB
}

type ICacheService interface {
	TestDbTransaction() error
}

func (s *CacheService) TestDbTransaction() (err error) {
	// open transaction
	tx := s.TX.Begin()
	var (
		ok bool
	)
	defer func() {
		// recover error, transaction rollback
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
	cache, err := s.IDB.FindCacheByKey("authorization_640c9b6eb58c11eca67a000c2933e6")
	if err != nil {
		return err
	}
	if cache == nil {
		return errors.New(errors.ServerError)
	}
	cache.CacheValue = "updateCol222"
	ok, err = s.IDB.UpdateCache(tx, cache)
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
		CacheKey:   "TestKey",
		CacheValue: "TestVal",
		Expired:    time.Time{}.UTC().Add(24 * time.Hour),
	}
	ok, err = s.IDB.AddCache(tx, c)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New(errors.ServerError)
	}

	// commit transaction
	return tx.Commit().Error
}
