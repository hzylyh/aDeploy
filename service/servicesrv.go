/*
 * @Description:
 * @Author: John Holl
 * @Github: https://github.com/hzylyh
 * @Date: 2021-05-04 09:46:17
 * @LastEditors: John Holl
 * @LastEditTime: 2021-05-04 10:34:51
 */
package service

import (
	"aDeploy/conf"
	"context"

	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var Service = &k8sService{}

type k8sService struct {
}

func (srv *k8sService) CreateService(c *gin.Context) (service *v1.Service, err error) {

	err = c.ShouldBind(&service)
	// conf.Clientset.CoreV1().Services("default").Create(context.TODO(), &v1.Service{})
	service, err = conf.Clientset.CoreV1().Services("default").Create(context.TODO(), service, metav1.CreateOptions{
		TypeMeta:     metav1.TypeMeta{},
		DryRun:       nil,
		FieldManager: "",
	})
	return
}

//func (srv *demoService) DeleteDeployment(c *gin.Context) (deployment *v1.Deployment, err error) {
//}
