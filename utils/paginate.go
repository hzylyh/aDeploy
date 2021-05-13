/*
 * @Description: gorm分页方法
 * @Author: John Holl
 * @Github: https://github.com/hzylyh
 * @Date: 2021-05-10 17:17:30
 * @LastEditors: John Holl
 * @LastEditTime: 2021-05-10 18:20:02
 */
package utils

import (
	"aDeploy/utils/qjson"

	"gorm.io/gorm"
)

func Paginate(qj *qjson.QJson) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// page, _ := strconv.Atoi(r.Query("page"))
		pageNum := qj.GetNum("pageNum")
		if pageNum == 0 {
			pageNum = 1
		}

		// pageSize, _ := strconv.Atoi(r.Query("page_size"))
		pageSize := qj.GetNum("pageSize")
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (pageNum - 1) * pageSize
		return db.Offset(int(offset)).Limit(int(pageSize))
	}
}
