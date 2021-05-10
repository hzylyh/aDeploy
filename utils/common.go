/*
 * @Description:
 * @Author: John Holl
 * @Github: https://github.com/hzylyh
 * @Date: 2021-05-10 17:28:09
 * @LastEditors: John Holl
 * @LastEditTime: 2021-05-10 17:28:18
 */
package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func GetJsonBody(c *gin.Context) map[string]interface{} {
	var (
		reqInfo map[string]interface{}
		body    []byte
		err     error
	)
	if body, err = ioutil.ReadAll(c.Request.Body); err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(body, &reqInfo)
	return reqInfo
}
