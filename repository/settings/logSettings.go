package settings

import (
	"errors"
	memo_slice "github.com/memo28-space-golang/memo28.slice"
	"github.com/rs/zerolog/log"
	"time"
	"workspace/logs/db"
	"workspace/logs/modules/monitoringErrors"
	"workspace/logs/modules/settings"
)

type LogSettingsRepository struct {
}

func (receiver LogSettingsRepository) GetSettings() (error, *settings.LogSettings) {

	var LogSettings settings.LogSettings

	err := db.DB.Model(&settings.LogSettings{}).Last(&LogSettings).Error

	if err != nil {
		log.Error().Msg(err.Error())
		return err, nil
	}

	if LogSettings.Number == 0 && LogSettings.ExpirationTime == 0 {
		log.Error().Msg("logSettings配置不可用")
		return errors.New("logSettings配置不可用"), nil
	}

	return nil, &LogSettings
}

func (receiver LogSettingsRepository) UpdateSettings(set settings.LogSettings) (error, uint) {

	createError := db.DB.Model(&settings.LogSettings{}).Create(&set).Error
	if createError != nil {
		log.Error().Msg(createError.Error())
		return createError, 0
	}

	go receiver.DeleteLogsDontComplyWithConfiguration()
	return nil, set.ID
}

func (receiver LogSettingsRepository) DeleteLogsDontComplyWithConfiguration() {
	err, set := receiver.GetSettings()
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Info().Msg("start clear...")
	go receiver.DeleteLogsDontComplyWithConfigurationByNumber(set)
	go receiver.DeleteLogsDontComplyWithConfigurationByExpirationTime(set)
}

func (receiver LogSettingsRepository) DeleteLogsDontComplyWithConfigurationByExpirationTime(set *settings.LogSettings) {

	t := int64(set.ExpirationTime) * int64(24*time.Hour)

	var list memo_slice.Slice[monitoringErrors.FrontendErrorReport]

	queryList := db.DB.Model(&monitoringErrors.FrontendErrorReport{}).Select("ID,created_at").Find(&list).Error
	if queryList != nil {
		log.Error().Msg(queryList.Error())
	}

	filterListByCreateAt := list.Filter(func(report monitoringErrors.FrontendErrorReport) bool {
		add := report.CreatedAt.Add(time.Duration(t))
		result := time.Now().After(add)
		return result
	}).Map(func(report monitoringErrors.FrontendErrorReport) interface{} {
		return report.ID
	})

	if len(filterListByCreateAt) != 0 {
		go receiver.DeleteLogsDontComplyWithConfigurationByIds(filterListByCreateAt, func() {
			log.Info().Msg("clear expirationTime done...")
		})
	}
}

func (receiver LogSettingsRepository) DeleteLogsDontComplyWithConfigurationByIds(ids []interface{}, callback func()) {
	deleteError := db.DB.Delete(&monitoringErrors.FrontendErrorReport{}, ids).Error
	if deleteError != nil {
		log.Error().Msg(deleteError.Error())
	} else {
		callback()
	}
}

func (receiver LogSettingsRepository) DeleteLogsDontComplyWithConfigurationByNumber(settings *settings.LogSettings) {
	var monitoringErrorsCount int64

	countError := db.DB.Model(&monitoringErrors.FrontendErrorReport{}).Count(&monitoringErrorsCount).Error

	if countError != nil {
		log.Error().Msg(countError.Error())
	} else {
		if settings.Number < monitoringErrorsCount {
			// 删除数量大于 Number 的log记录
			maxNumber := monitoringErrorsCount - settings.Number
			var ids []interface{}
			idsError := db.DB.Model(&monitoringErrors.FrontendErrorReport{}).Order("ID DESC").Limit(int(maxNumber)).Select("ID").Find(&ids).Error
			if idsError != nil {
				log.Error().Msg(idsError.Error())
				return
			} else {
				go receiver.DeleteLogsDontComplyWithConfigurationByIds(ids, func() {
					log.Info().Msg("clear with number done...")
				})
			}
		}
	}
}

// RegularlyPerform 定时任务执行清除无用的errorLog数据
func (receiver LogSettingsRepository) RegularlyPerform() {
	// 创建一个定时器，每隔30秒触发一次
	ticker := time.NewTicker(10 * time.Second)
	//defer ticker.Stop()

	// 使用goroutine执行定时任务
	go func() {
		for {
			select {
			case <-ticker.C:
				// 定时任务的具体逻辑
				go receiver.DeleteLogsDontComplyWithConfiguration()
			}
		}
	}()

}
