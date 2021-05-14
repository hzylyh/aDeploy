/*
 * @Description:
 * @Author: John Holl
 * @Github: https://github.com/hzylyh
 * @Date: 2021-05-05 17:05:21
 * @LastEditors: John Holl
 * @LastEditTime: 2021-05-14 10:52:32
 */
package models

import "aDeploy/conf"

type TContainerInfo struct {
	conf.Model
	ImageId             string `json:"imageId"`
	ContainerName       string `json:"containerName"`
	ContainerDeployInfo string `json:"containerDeployInfo"`
	ContainerDynamicCol string `json:"containerDynamicCol"`
	ContainerDesc       string `json:"containerDesc"`
}
