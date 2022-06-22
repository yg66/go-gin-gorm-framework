package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yg66/go-gin-gorm-framework/common/logger"
	"github.com/yg66/go-gin-gorm-framework/config"
	"github.com/yg66/go-gin-gorm-framework/db"
	"github.com/yg66/go-gin-gorm-framework/handler"
	"github.com/yg66/go-gin-gorm-framework/routers"
	"log"
	"net/http"
)

func main() {
	// ---- Init ----
	myConfig, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = logger.Init(myConfig.LogConfig)
	if err != nil {
		log.Fatal(err)
	}

	if myConfig.BasicsConfig.IsDev {
		go func() {
			err := http.ListenAndServe(":6666", nil)
			if err != nil {
				log.Fatal(err)
			}
		}()
	}

	myDb, e := db.Init(myConfig.DbConfig)
	if e != nil {
		log.Fatal(e)
	}

	//// ---- start job ----
	//myJob := job.Job{
	//	JobConfig: &myConfig.JobConfig,
	//	MyChain:   myChain,
	//	MyDb:      myDb,
	//}
	//myJob.StartJob(myConfig.JobConfig.JobCycle, myJob.DelExpiredCache)

	// ---- rpc ----
	ginEngine := gin.Default()
	if myConfig.BasicsConfig.IsDev {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// gin log
	//f, _ := os.Create("/tmp/go_gin_gorm_framework.log")
	//gin.DefaultErrorWriter = io.MultiWriter(f, os.Stdout)
	ginEngine.NoRoute(handler.HandleNotFound)
	ginEngine.NoMethod(handler.HandleNotFound)
	ginEngine.Use(handler.GinLogger(), handler.GinRecovery(true), handler.Cors())

	// load routers
	if router, err := routers.NewRouter(myDb, myConfig); err != nil {
		log.Fatal(err)
	} else {
		router.LoadRouters(ginEngine)
	}

	// run
	e = ginEngine.Run(fmt.Sprintf(":%d", myConfig.BasicsConfig.Port))
	if e != nil {
		log.Fatal(e)
	}
}
