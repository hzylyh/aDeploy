/*
 * @Description:
 * @Author: John Holl
 * @Github: https://github.com/hzylyh
 * @Date: 2021-05-04 09:46:17
 * @LastEditors: John Holl
 * @LastEditTime: 2021-05-17 20:51:46
 */
package service

import (
	"aDeploy/conf"
	"aDeploy/utils/qjson"
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

func (srv *k8sService) DeleteService(qj *qjson.QJson) (err error) {
	name := qj.GetString("name")
	// conf.Clientset.CoreV1().Services("default").Create(context.TODO(), &v1.Service{})
	if err = conf.Clientset.CoreV1().Services("default").Delete(context.TODO(), name, metav1.DeleteOptions{}); err != nil {
		return
	}
	return
}

//func (srv *demoService) DeleteDeployment(c *gin.Context) (deployment *v1.Deployment, err error) {
//}

func (srv *k8sService) GetServiceInfo(c *gin.Context) (services *v1.ServiceList, err error) {
	if services, err = conf.Clientset.CoreV1().Services("default").List(context.TODO(), metav1.ListOptions{}); err != nil {
		return
	}
	return
}
