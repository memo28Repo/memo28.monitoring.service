package settings

import (
	"github.com/gin-gonic/gin"
	"workspace/logs/modules/settings"
	settings2 "workspace/logs/repository/settings"
)

type LogSettingsService struct {
}

func (receiver LogSettingsService) GetSettings() (error, *settings.LogSettings) {
	logSettingsRepository := settings2.LogSettingsRepository{}

	return logSettingsRepository.GetSettings()
}

func (receiver LogSettingsService) UpdateSettings(c *gin.Context) (error, uint) {

	var logSettings settings.LogSettings

	err := c.ShouldBind(&logSettings)
	if err != nil {
		return err, 0
	}

	logSettingsRepository := settings2.LogSettingsRepository{}

	return logSettingsRepository.UpdateSettings(logSettings)
}
