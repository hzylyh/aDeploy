/*
 * @Description: deployment操作的controller层
 * @Author: John Holl
 * @Github: https://github.com/hzylyh
 * @Date: 2021-05-04 09:43:49
 * @LastEditors: John Holl
 * @LastEditTime: 2021-05-13 13:06:24
 */
package controller

import (
	"aDeploy/conf"
	"aDeploy/service"
	"aDeploy/utils"

	"github.com/gin-gonic/gin"
)

func CreateDeployment(ctx *gin.Context) {
	if depinfo, err := service.DeploymentSrv.CreateDeployment(ctx); err != nil {
		utils.ResponseFailOther(ctx, conf.InvalidParam, err.Error())
	} else {
		utils.ResponseOkWithMsg(ctx, "ok", depinfo)
	}

}

func GetDeploymentStatus(ctx *gin.Context) {
	// if depStatus, err := service.DeploymentSrv.GetDeploymentStatus(); err != nil {
	// 	utils.ResponseFailOther(ctx, conf.InvalidParam, err.Error())
	// } else {
	// 	utils.ResponseOkWithMsg(ctx, "ok", depStatus)
	// }
}
