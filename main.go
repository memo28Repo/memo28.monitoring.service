package main

import (
	"github.com/gin-gonic/gin"
	memo28Gin "github.com/memo28-space-golang/memo28.gin/pkg"
	"workspace/logs/api/monitoringErrors"
	"workspace/logs/db"
	"workspace/logs/errors"
)

func init() {
	db.DbInit()
}

func main() {
	r := gin.Default()
	r.Use(memo28Gin.Cors(), memo28Gin.CatchGlobalError(memo28Gin.CatchGlobalErrorOptions{
		Callback: func(msg string, c *gin.Context) {
			response := memo28Gin.Response{}
			response.Error(c, errors.ASystemException, nil, "系统异常")
		},
	}))

	apiPrefix := r.Group("api")

	v1Version := apiPrefix.Group("v1")

	{
		// 错误监控
		monitoringErrorsApi := v1Version.Group("monitoringErrors")

		monitoringErrorsApi.POST("/createErrorLog", monitoringErrors.CreateErrorLog)

		monitoringErrorsApi.DELETE("/deleteErrorLog", monitoringErrors.DeleteErrorLog)

		monitoringErrorsApi.GET("/getErrorLog", monitoringErrors.GetErrorLog)

		monitoringErrorsApi.GET("/getErrorLogList", monitoringErrors.GetErrorLogList)
	}

	_ = r.Run(":8089")
}
