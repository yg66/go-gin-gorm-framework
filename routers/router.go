package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/yg66/go-gin-gorm-framework/common/errors"
	"github.com/yg66/go-gin-gorm-framework/common/res"
	"github.com/yg66/go-gin-gorm-framework/config"
	"github.com/yg66/go-gin-gorm-framework/db"
	"github.com/yg66/go-gin-gorm-framework/service"
	"go.uber.org/zap"
)

type Router struct {
	iConfig       config.Config
	iCacheService service.ICacheService
}

func NewRouter(idb *db.Db, iConfig config.Config) (*Router, error) {
	router := Router{
		iConfig:       iConfig,
		iCacheService: &service.CacheService{IDB: idb, TX: idb.IDB},
	}
	return &router, nil
}

func (myRouter *Router) LoadRouters(engine *gin.Engine) {
	engine.POST("/test", test(myRouter))
}

func test(r *Router) gin.HandlerFunc {
	return func(context *gin.Context) {
		err := r.iCacheService.TestDbTransaction()
		if err != nil {
			zap.S().Errorf("err: %v", err)
			panic(errors.New(errors.ServerError))
		} else {
			success := res.Success(nil)
			context.JSON(success.Code, success)
		}
	}
}
