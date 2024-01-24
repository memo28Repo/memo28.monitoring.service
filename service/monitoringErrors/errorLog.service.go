package monitoringErrors

import (
	"errors"
	"github.com/gin-gonic/gin"
	"workspace/logs/common"
	"workspace/logs/modules/monitoringErrors"
	monitoringErrorsRepository "workspace/logs/repository/monitoringErrors"
)

type ErrorLogService struct {
}

// ErrorLogServiceParams 用户的请求
type ErrorLogServiceParams struct {
	Data monitoringErrors.FrontendErrorReport `json:"data"`
	Id   uint                                 `json:"id"`
}

func (receiver ErrorLogService) GetParameters(c *gin.Context) (*ErrorLogServiceParams, error) {
	var postData ErrorLogServiceParams

	err := c.ShouldBind(&postData)

	if err != nil {
		print(err.Error())
		return nil, nil
	}
	return &postData, nil
}

func (receiver ErrorLogService) CreateLog(c *gin.Context) (error, uint) {
	parameters, err := receiver.GetParameters(c)
	if err != nil {
		print(err.Error())
		return errors.New("解析参数异常"), 0
	}
	errorLogRepository := monitoringErrorsRepository.ErrorLogRepository{}
	create, u := errorLogRepository.Create(&parameters.Data)
	if create != nil {
		return create, 0
	}
	return nil, u
}

func (receiver ErrorLogService) DeleteLog(c *gin.Context) (error, uint) {
	parameters, err := receiver.GetParameters(c)
	if err != nil {
		print(err.Error())
		return errors.New("解析参数异常"), 0
	}
	errorLogRepository := monitoringErrorsRepository.ErrorLogRepository{}
	del, u := errorLogRepository.Delete(parameters.Id)
	if del != nil {
		return del, 0
	}
	return nil, u
}

func (receiver ErrorLogService) GetErrorLog(c *gin.Context) (error, *monitoringErrors.FrontendErrorReport) {
	query, b := c.GetQuery("id")

	if !b {
		return errors.New("id参数不能为空"), nil
	}

	errorLogRepository := monitoringErrorsRepository.ErrorLogRepository{}

	getError, u := errorLogRepository.Get(query)
	if getError != nil {
		return getError, u
	}
	return nil, u
}

func (receiver ErrorLogService) GetList(c *gin.Context) (error, []monitoringErrors.FrontendErrorReport, int64) {
	pageNo, hasPageNo := c.GetQuery("pageNo")
	if !hasPageNo {
		return errors.New("pageNo参数不能为空"), nil, 0
	}

	pageSize, hasPageSize := c.GetQuery("pageSize")
	if !hasPageSize {
		return errors.New("pageSize参数不能为空"), nil, 0
	}

	option := common.PageOptions{
		PageSize: pageSize,
		PageNo:   pageNo,
	}

	errorLogRepository := monitoringErrorsRepository.ErrorLogRepository{}

	err, list := errorLogRepository.List(option)

	_, count := errorLogRepository.GetCount()

	if err != nil {
		return err, nil, count
	}

	return nil, list, count
}
