package monitoringErrors

import (
	"github.com/gin-gonic/gin"
	"github.com/memo28-space-golang/memo28.gin/pkg"
	memo28Map "github.com/memo28-space-golang/memo28.map/pkg"
	"workspace/logs/errors"
	"workspace/logs/service/monitoringErrors"
)

func CreateErrorLog(c *gin.Context) {
	logService := monitoringErrors.ErrorLogService{}

	logError, id := logService.CreateLog(c)

	response := memo28Gin.Response{}

	if logError != nil {
		response.Error(c, errors.ASystemException, "", logError.Error())
		return
	}

	response.Suc(c, id)
}

func DeleteErrorLog(c *gin.Context) {
	logService := monitoringErrors.ErrorLogService{}

	logError, id := logService.DeleteLog(c)

	response := memo28Gin.Response{}

	if logError != nil {
		response.Error(c, errors.ASystemException, "", logError.Error())
		return
	}

	response.Suc(c, id)
}

func GetErrorLog(c *gin.Context) {
	logService := monitoringErrors.ErrorLogService{}

	logError, id := logService.GetErrorLog(c)

	response := memo28Gin.Response{}

	if logError != nil {
		response.Error(c, errors.ASystemException, "", logError.Error())
		return
	}

	response.Suc(c, id)
}

func GetErrorLogList(c *gin.Context) {
	logService := monitoringErrors.ErrorLogService{}

	logError, list, count := logService.GetList(c)

	response := memo28Gin.Response{}

	if logError != nil {
		response.Error(c, errors.ASystemException, "", logError.Error())
		return
	}

	m := memo28Map.Map[string, interface{}]{}

	m.Add("list", list)

	m.Add("total", count)

	response.Suc(c, m)
}
