/*
 * @Description: 容器service
 * @Author: John Holl
 * @Github: https://github.com/hzylyh
 * @Date: 2021-05-10 19:43:51
 * @LastEditors: John Holl
 * @LastEditTime: 2021-05-14 10:49:15
 */
package service

import (
	"aDeploy/conf"
	"aDeploy/models/vo"
	"aDeploy/utils/qjson"
	"fmt"
)

var ContainerSrv = &containerService{}

type containerService struct {
}

/**

 */
func (cs *containerService) AddContainerInfo(err error) {

}

func (cs *containerService) EditContainerInfo(err error) {

}

func (cs *containerService) DeleteContainerInfo(err error) {

}

func (cs *containerService) GetContainerInfo(err error) {

}

func (cs *containerService) GetContainerList(qj *qjson.QJson) (containerInfo []*vo.ContainerInfoVO, err error) {

	imageList := qj.GetArray("imageList")
	fmt.Println(imageList)
	if err := conf.DB.Table("t_container_infos").Where("container_id in ?", imageList).Find(&containerInfo).Error; err != nil {
		return nil, err
	}

	return containerInfo, nil
}
