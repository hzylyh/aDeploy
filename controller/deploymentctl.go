/*
 * @Description: deployment操作的controller层
 * @Author: John Holl
 * @Github: https://github.com/hzylyh
 * @Date: 2021-05-04 09:43:49
 * @LastEditors: John Holl
 * @LastEditTime: 2021-05-04 09:44:58
 */
package controller

import (
	"aDeploy/conf"
	"aDeploy/service"
	"aDeploy/utils"

	"github.com/gin-gonic/gin"
)

func CreateDeployment(ctx *gin.Context) {
	if depinfo, err := service.Deployment.CreateDeployment(ctx); err != nil {
		utils.ResponseFailOther(ctx, conf.InvalidParam, err.Error())
	} else {
		utils.ResponseOkWithMsg(ctx, "ok", depinfo)
	}

}