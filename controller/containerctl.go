/*
 * @Description: container操作的controller层
 * @Author: John Holl
 * @Github: https://github.com/hzylyh
 * @Date: 2021-05-10 19:41:40
 * @LastEditors: John Holl
 * @LastEditTime: 2021-05-10 19:55:12
 */

package controller

import (
	"aDeploy/conf"
	"aDeploy/service"
	"aDeploy/utils"
	"aDeploy/utils/qjson"

	"github.com/gin-gonic/gin"
)

func AddContainerInfo(c *gin.Context) {
	utils.ResponseOk(c, "ddd")
}

func EditContainerInfo(c *gin.Context) {

}

func DeleteContainerInfo(c *gin.Context) {

}

func GetContainernfo(c *gin.Context) {

}

func GetContainerList(c *gin.Context) {
	reqInfo := qjson.QJson{
		ReqInfo: utils.GetJsonBody(c),
	}
	if containerInfoList, err := service.ContainerSrv.GetContainerList(&reqInfo); err != nil {
		utils.ResponseFailOther(c, conf.InvalidParam, err.Error())
	} else {
		utils.ResponseOk(c, containerInfoList)
	}
}
