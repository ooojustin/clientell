package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetPageSize(c *gin.Context) int {
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	return pageSize
}

func GetPage(c *gin.Context) int {
	page, _ := strconv.Atoi(c.Query("page"))
	if page == 0 {
		page = 1
	}
	return page
}

func Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := GetPage(c)
		pageSize := GetPageSize(c)
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
