/*
 * @Description:镜像service
 * @Author: John Holl
 * @Github: https://github.com/hzylyh
 * @Date: 2021-05-05 21:49:54
 * @LastEditors: John Holl
 * @LastEditTime: 2021-05-10 18:18:42
 */

package service

import (
	"aDeploy/conf"
	"aDeploy/models/vo"
	"aDeploy/utils"
	"aDeploy/utils/qjson"
)

var ImageSrv = &imageService{}

type imageService struct {
}

/**

 */
func (is *imageService) AddImageInfo(err error) {

}

func (is *imageService) EditImageInfo(err error) {

}

func (is *imageService) DeleteImageInfo(err error) {

}

func (is *imageService) GetImageInfo(err error) {

}

func (is *imageService) GetImageList(qj *qjson.QJson) (imageInfo []*vo.ImageInfoVO, err error) {
	var (
	// ret      []*vo.CaseStepInfoVO
	// pageNum  = qj.GetNum("pageNum")
	// pageSize = qj.GetNum("pageSize")
	// caseId   = qj.GetString("caseId")
	)
	if err := conf.DB.Table("t_image_infos").Scopes(utils.Paginate(qj)).Find(&imageInfo).Error; err != nil {
		return nil, err
	}

	return imageInfo, nil
}
