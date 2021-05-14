/*
 * @Description:
 * @Author: John Holl
 * @Github: https://github.com/hzylyh
 * @Date: 2021-05-10 19:45:28
 * @LastEditors: John Holl
 * @LastEditTime: 2021-05-14 10:54:02
 */

package vo

type ContainerInfoVO struct {
	Id                  string `json:"containerId"`
	ImageId             string `json:"imageId"`
	ContainerName       string `json:"imageName"`
	ContainerDeployInfo string `json:"containerDeployInfo"`
	ContainerDynamicCol string `json:"containerDynamicCol"`
	ContainerDesc       string `json:"imageDesc"`
}
