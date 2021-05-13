/*
 * @Description:
 * @Author: John Holl
 * @Github: https://github.com/hzylyh
 * @Date: 2021-05-05 17:05:21
 * @LastEditors: John Holl
 * @LastEditTime: 2021-05-10 19:21:44
 */
package models

import "aDeploy/conf"

type TContainerInfo struct {
	conf.Model
	ContainerId         string `json:"containerId"`
	ContainerName       string `json:"containerName"`
	ContainerDeployInfo string `json:"containerDeployInfo"`
	ContainerDynamicCol string `json:"containerDynamicCol"`
	ContainerDesc       string `json:"containerDesc"`
}
