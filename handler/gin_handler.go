package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yg66/go-gin-gorm-framework/common/errors"
	"github.com/yg66/go-gin-gorm-framework/common/res"
	"go.uber.org/zap"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

// GinLogger Receive the default log of the gin framework
func GinLogger() gin.HandlerFunc {
	logger := zap.L()
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		logger.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// GinRecovery gin exception handling
func GinRecovery(stack bool) gin.HandlerFunc {
	logger := zap.S()
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.Errorf(
						"\n\n**************** Error ****************\n%v"+
							"\n**************** Request ****************\n%v\n",
						err,
						string(httpRequest),
					)
					// If the connection is dead, we can't write a status to it.
					failed := res.Failed(errors.New(errors.NetworkAnomaly))
					c.JSON(failed.Code, failed)
					//c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					logger.Errorf(
						"\n\n**************** Error ****************\n%v"+
							"\n**************** Request ****************\n%v"+
							"\n**************** Stack ****************\n%v\n",
						err,
						string(httpRequest),
						string(debug.Stack()),
					)
				} else {
					logger.Errorf(
						"\n\n**************** Error ****************\n%v"+
							"\n**************** Request ****************\n%v\n",
						err,
						string(httpRequest),
					)
				}

				switch err.(type) {
				case *errors.Err:
					e := err.(*errors.Err)
					failed := res.Failed(e)
					c.JSON(failed.Code, failed)
					c.Abort()
				default:
					unknownErr := res.UnknownErr(err)
					c.JSON(http.StatusOK, unknownErr)
					c.Abort()
				}
			}
		}()
		c.Next()
	}
}

// HandleNotFound 404
func HandleNotFound(c *gin.Context) {
	zap.S().Errorf("handle not found: %v", c.Request.RequestURI)
	c.JSON(http.StatusNotFound, res.Failed(errors.New(errors.UriNotFoundOrMethodNotSupport)))
}
