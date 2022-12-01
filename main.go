package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yg66/go-gin-gorm-framework/common/logger"
	"github.com/yg66/go-gin-gorm-framework/config"
	"github.com/yg66/go-gin-gorm-framework/db"
	"github.com/yg66/go-gin-gorm-framework/handler"
	"github.com/yg66/go-gin-gorm-framework/routers"
)

func main() {
	// ---- Init ----
	iConfig, err := config.InitConfigByYml()
	if err != nil {
		log.Fatal(err)
	}

	if iConfig.BasicsConfig.IsDev {
		go func() {
			err := http.ListenAndServe(":6666", nil)
			if err != nil {
				log.Fatal(err)
			}
		}()
	}

	err = logger.Init(iConfig.LogConfig)
	if err != nil {
		log.Fatal(err)
	}

	idb, e := db.NewDB(iConfig.DbConfig)
	if e != nil {
		log.Fatal(e)
	}

	// ---- rpc ----
	ginEngine := gin.Default()
	if iConfig.BasicsConfig.IsDev {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	ginEngine.NoRoute(handler.HandleNotFound)
	ginEngine.NoMethod(handler.HandleNotFound)
	ginEngine.Use(handler.GinLogger(), handler.GinRecovery(iConfig.BasicsConfig.StackTrace), handler.Cors())

	// load routers
	if router, err := routers.NewRouter(idb, *iConfig); err != nil {
		log.Fatal(err)
	} else {
		router.LoadRouters(ginEngine)
	}

	// run
	e = ginEngine.Run(fmt.Sprintf(":%d", iConfig.BasicsConfig.Port))
	if e != nil {
		log.Fatal(e)
	}
}
