/*
 * @Description:
 * @Author: John Holl
 * @Github: https://github.com/hzylyh
 * @Date: 2021-05-10 19:45:28
 * @LastEditors: John Holl
 * @LastEditTime: 2021-05-10 19:46:29
 */

package vo

type ContainerInfoVO struct {
	ContainerId         string `json:"imageId"`
	ContainerName       string `json:"imageName"`
	ContainerDeployInfo string `json:"containerDeployInfo"`
	ContainerDynamicCol string `json:"containerDynamicCol"`
	ContainerDesc       string `json:"imageDesc"`
}
