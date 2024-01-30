package monitoringErrors

import (
	"errors"
	"github.com/rs/zerolog/log"
	"workspace/logs/common"
	"workspace/logs/db"
	"workspace/logs/modules/monitoringErrors"
)

type ErrorLogRepository struct {
}

func (receiver ErrorLogRepository) Create(reportData *monitoringErrors.FrontendErrorReport) (error, uint) {
	err := db.DB.Table(monitoringErrors.FrontendErrorReport{}.TableName()).Create(reportData).Error

	if err != nil {
		log.Error().Msg(err.Error())
		return errors.New("创建errorLog失败,系统异常"), 0
	}
	return nil, reportData.ID
}

func (receiver ErrorLogRepository) Delete(id uint) (error, uint) {
	err := db.DB.Table(monitoringErrors.FrontendErrorReport{}.TableName()).Where("id = ?", id).Delete(id).Error

	if err != nil {
		log.Error().Msg(err.Error())
		return errors.New("删除errorLog失败,系统异常"), 0
	}
	return nil, id
}

type ListQuery struct {
	PageOptions common.PageOptions
	// Types 筛选错误类型 window.error | user-defined ...
	Types string
}

func (receiver ErrorLogRepository) List(page ListQuery) (error, []monitoringErrors.FrontendErrorReport) {

	var list []monitoringErrors.FrontendErrorReport

	handler := common.Paginate(db.DB, page.PageOptions).Model(monitoringErrors.FrontendErrorReport{}).Order("ID DESC")

	if page.Types != "" {
		handler = handler.Where("type = ?", page.Types)
	}

	err := handler.Find(&list).Error

	if err != nil {
		log.Error().Msg(err.Error())
		return err, nil
	}
	return nil, list
}

func (receiver ErrorLogRepository) GetCount() (error, int64) {

	var count int64

	err := db.DB.Model(monitoringErrors.FrontendErrorReport{}).Count(&count).Error

	if err != nil {
		log.Error().Msg(err.Error())
		return err, 0
	}

	return nil, count

}

func (receiver ErrorLogRepository) Get(id string) (error, *monitoringErrors.FrontendErrorReport) {

	result := &monitoringErrors.FrontendErrorReport{}

	err := db.DB.Table(monitoringErrors.FrontendErrorReport{}.TableName()).Where("id = ?", id).Find(result).Error

	if err != nil {
		log.Error().Msg(err.Error())
		return errors.New("查询errorLog失败,系统异常"), nil
	}
	return nil, result
}
