/*
 * @Description:
 * @Author: John Holl
 * @Github: https://github.com/hzylyh
 * @Date: 2021-05-05 17:04:45
 * @LastEditors: John Holl
 * @LastEditTime: 2021-05-14 10:55:26
 */
package models

import "aDeploy/conf"

type TImageInfo struct {
	conf.Model
	// ImageId      string `json:"imageId"`
	ImageName    string `json:"imageName"`
	ImageVersion string `json:"imageVersion"`
	ImageDesc    string `json:"imageDesc"`
}
