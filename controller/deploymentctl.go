/*
 * @Description: deployment操作的controller层
 * @Author: John Holl
 * @Github: https://github.com/hzylyh
 * @Date: 2021-05-04 09:43:49
 * @LastEditors: John Holl
 * @LastEditTime: 2021-05-17 20:14:20
 */
package controller

import (
	"aDeploy/conf"
	"aDeploy/service"
	"aDeploy/utils"
	"aDeploy/utils/qjson"

	"github.com/gin-gonic/gin"
)

/**
 * @name:
 * @description:
 * @param {type} {*}
 * @return {type} {*}
 * @param {*gin.Context} ctx
 */
func CreateDeployment(ctx *gin.Context) {
	if depinfo, err := service.DeploymentSrv.CreateDeployment(ctx); err != nil {
		utils.ResponseFailOther(ctx, conf.InvalidParam, err.Error())
	} else {
		utils.ResponseOkWithMsg(ctx, "ok", depinfo)
	}

}

/**
 * @name:
 * @description:
 * @param {type} {*}
 * @return {type} {*}
 * @param {*gin.Context} ctx
 */
func DeleteDeployment(ctx *gin.Context) {
	reqInfo := qjson.QJson{
		ReqInfo: utils.GetJsonBody(ctx),
	}
	if err := service.DeploymentSrv.DeleteDeployment(&reqInfo); err != nil {
		utils.ResponseFailOther(ctx, conf.InvalidParam, err.Error())
	} else {
		utils.ResponseOk(ctx, nil)
	}

}

/**
 * @name:
 * @description:
 * @param {type} {*}
 * @return {type} {*}
 * @param {*gin.Context} ctx
 */
func GetDeploymentInfo(ctx *gin.Context) {
	if depStatus, err := service.DeploymentSrv.GetDeployment(ctx); err != nil {
		utils.ResponseFailOther(ctx, conf.InvalidParam, err.Error())
	} else {
		utils.ResponseOkWithMsg(ctx, "ok", depStatus)
	}
}
