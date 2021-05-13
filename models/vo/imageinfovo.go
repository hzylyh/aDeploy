/*
 * @Description:
 * @Author: John Holl
 * @Github: https://github.com/hzylyh
 * @Date: 2021-05-10 17:08:03
 * @LastEditors: John Holl
 * @LastEditTime: 2021-05-10 18:33:00
 */

package vo

type ImageInfoVO struct {
	ImageId      string `json:"imageId"`
	ImageName    string `json:"imageName"`
	ImageVersion string `json:"imageVersion"`
	ImageDesc    string `json:"imageDesc"`
}
