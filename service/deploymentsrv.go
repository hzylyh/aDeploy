/*
 * @Description:deployment相关操作
 * @Author: John Holl
 * @Github: https://github.com/hzylyh
 * @Date: 2021-05-04 09:41:14
 * @LastEditors: John Holl
 * @LastEditTime: 2021-05-13 13:11:16
 */
package service

import (
	"aDeploy/conf"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	appsV1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DeploymentSrv = &deploymentService{}

type deploymentService struct {
}

func (srv *deploymentService) CreateDeployment(c *gin.Context) (deployment *appsV1.Deployment, err error) {
	err = c.ShouldBind(&deployment)
	fmt.Println(deployment.ObjectMeta)
	fmt.Println(deployment.Status)
	deployment, err = conf.Clientset.AppsV1().Deployments("default").Create(context.TODO(), deployment, metav1.CreateOptions{
		TypeMeta:     metav1.TypeMeta{},
		DryRun:       nil,
		FieldManager: "",
	})
	return
}

//func (srv *demoService) DeleteDeployment(c *gin.Context) (deployment *v1.Deployment, err error) {
//}

func (srv *deploymentService) GetDeploymentStatus(c *gin.Context) (depStatus *appsV1.Deployment, err error) {
	// conf.Clientset.AppsV1().Deployments("default").Get()
	return
}
