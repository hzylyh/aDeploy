/*
 * @Description:
 * @Author: John Holl
 * @Github: https://github.com/hzylyh
 * @Date: 2021-05-04 09:46:00
 * @LastEditors: John Holl
 * @LastEditTime: 2021-05-17 20:52:57
 */
package controller

import (
	"aDeploy/conf"
	"aDeploy/service"
	"aDeploy/utils"
	"aDeploy/utils/qjson"

	"github.com/gin-gonic/gin"
)

func CreateService(ctx *gin.Context) {
	if depinfo, err := service.Service.CreateService(ctx); err != nil {
		utils.ResponseFailOther(ctx, conf.InvalidParam, err.Error())
	} else {
		utils.ResponseOkWithMsg(ctx, "ok", depinfo)
	}

}

func DeleteService(ctx *gin.Context) {
	reqInfo := qjson.QJson{
		ReqInfo: utils.GetJsonBody(ctx),
	}
	if err := service.Service.DeleteService(&reqInfo); err != nil {
		utils.ResponseFailOther(ctx, conf.InvalidParam, err.Error())
	} else {
		utils.ResponseOk(ctx, "ok")
	}

}

func GetServiceInfo(ctx *gin.Context) {
	if svcStatus, err := service.Service.GetServiceInfo(ctx); err != nil {
		utils.ResponseFailOther(ctx, conf.InvalidParam, err.Error())
	} else {
		utils.ResponseOkWithMsg(ctx, "ok", svcStatus)
	}
}
