/*
 * @Description:
 * @Author: John Holl
 * @Github: https://github.com/hzylyh
 * @Date: 2021-05-04 09:46:00
 * @LastEditors: John Holl
 * @LastEditTime: 2021-05-04 10:35:09
 */
package controller

import (
	"aDeploy/conf"
	"aDeploy/service"
	"aDeploy/utils"

	"github.com/gin-gonic/gin"
)

func CreateService(ctx *gin.Context) {
	if depinfo, err := service.Service.CreateService(ctx); err != nil {
		utils.ResponseFailOther(ctx, conf.InvalidParam, err.Error())
	} else {
		utils.ResponseOkWithMsg(ctx, "ok", depinfo)
	}

}
