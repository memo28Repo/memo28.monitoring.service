package common

import (
	"gorm.io/gorm"
	"strconv"
)

type PageOptions struct {
	PageNo   string
	PageSize string
	Model    interface{}
}

func Paginate(db *gorm.DB, PageOptions PageOptions) *gorm.DB {
	page, _ := strconv.Atoi(PageOptions.PageNo)
	if page <= 0 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(PageOptions.PageSize)
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	tx := db.Model(PageOptions.Model).Offset(offset)

	return tx.Limit(pageSize)
}
