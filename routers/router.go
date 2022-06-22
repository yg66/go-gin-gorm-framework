package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/yg66/go-gin-gorm-framework/common/errors"
	"github.com/yg66/go-gin-gorm-framework/common/res"
	"github.com/yg66/go-gin-gorm-framework/db"
	"github.com/yg66/go-gin-gorm-framework/model"
	"github.com/yg66/go-gin-gorm-framework/service"
	"go.uber.org/zap"
)

type Router struct {
	MyConfig     *model.Config
	CacheService *service.Service
}

func NewRouter(myDb *db.MyDb, myConfig *model.Config) (*Router, error) {
	router := Router{
		CacheService: &service.Service{D: myDb},
	}
	return &router, nil
}

func (myRouter *Router) LoadRouters(engine *gin.Engine) {
	engine.POST("/test", test(myRouter))
}

func test(myRouter *Router) gin.HandlerFunc {
	return func(context *gin.Context) {
		err := myRouter.CacheService.TestDbTransaction()
		if err != nil {
			zap.S().Errorf("err: %v", err)
			panic(errors.New(errors.ServerError))
		} else {
			success := res.Success(nil)
			context.JSON(success.Code, success)
		}
	}
}
