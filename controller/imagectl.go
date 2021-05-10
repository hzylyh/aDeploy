/*
 * @Description:镜像controller
 * @Author: John Holl
 * @Github: https://github.com/hzylyh
 * @Date: 2021-05-05 21:44:01
 * @LastEditors: John Holl
 * @LastEditTime: 2021-05-10 17:42:25
 */
package controller

import (
	"aDeploy/conf"
	"aDeploy/service"
	"aDeploy/utils"
	"aDeploy/utils/qjson"

	"github.com/gin-gonic/gin"
)

func AddImageInfo(c *gin.Context) {
	utils.ResponseOk(c, "ddd")
}

func EditImageInfo(c *gin.Context) {

}

func DeleteImageInfo(c *gin.Context) {

}

func GetImageInfo(c *gin.Context) {

}

func GetImageList(c *gin.Context) {
	reqInfo := qjson.QJson{
		ReqInfo: utils.GetJsonBody(c),
	}
	if imageInfoList, err := service.ImageSrv.GetImageList(&reqInfo); err != nil {
		utils.ResponseFailOther(c, conf.InvalidParam, err.Error())
	} else {
		utils.ResponseOk(c, imageInfoList)
	}

}
