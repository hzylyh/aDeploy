/*
 * @Description:
 * @Author: John Holl
 * @Github: https://github.com/hzylyh
 * @Date: 2021-05-10 17:08:03
 * @LastEditors: John Holl
 * @LastEditTime: 2021-05-14 10:55:06
 */

package vo

type ImageInfoVO struct {
	Id           string `json:"imageId"`
	ImageName    string `json:"imageName"`
	ImageVersion string `json:"imageVersion"`
	ImageDesc    string `json:"imageDesc"`
}
