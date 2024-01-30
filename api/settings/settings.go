package settings

import (
	"github.com/gin-gonic/gin"
	memo_gin "github.com/memo28-space-golang/memo28.gin"
	"workspace/logs/errors"
	"workspace/logs/service/settings"
)

func GetSettingsDetails(c *gin.Context) {

	logSettingsService := settings.LogSettingsService{}

	response := memo_gin.Response{}

	err, s := logSettingsService.GetSettings()

	if err != nil {
		response.Error(c, errors.ASystemException, nil, err.Error())
		return
	}

	response.Suc(c, s)
}

func UpdateLogSettingsDetails(c *gin.Context) {

	logSettingsService := settings.LogSettingsService{}

	response := memo_gin.Response{}

	err, id := logSettingsService.UpdateSettings(c)
	if err != nil {
		response.Error(c, errors.ASystemException, nil, err.Error())
		return
	}
	response.Suc(c, id)
}
