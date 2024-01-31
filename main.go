package main

import (
	"github.com/gin-gonic/gin"
	memo28Gin "github.com/memo28-space-golang/memo28.gin"
	"workspace/logs/api/monitoringErrors"
	"workspace/logs/api/settings"
	"workspace/logs/db"
	"workspace/logs/errors"
	settings2 "workspace/logs/repository/settings"
)

func init() {
	db.DbInit()
}

func main() {
	r := gin.Default()

	r.Use(memo28Gin.Cors(), memo28Gin.CatchGlobalError(memo28Gin.CatchGlobalErrorOptions{
		Callback: func(msg string, c *gin.Context) {
			response := memo28Gin.Response{}
			println(msg)
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
	{
		// 配置设置
		logSettingsApi := v1Version.Group("logSettings")

		logSettingsApi.GET("/getSettings", settings.GetSettingsDetails)

		logSettingsApi.POST("/updateSettings", settings.UpdateLogSettingsDetails)
	}

	logSettingsRepository := settings2.LogSettingsRepository{}
	logSettingsRepository.RegularlyPerform()

	_ = r.Run(":8089")
}
