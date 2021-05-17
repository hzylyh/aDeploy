/*
 * @Description:deployment相关操作
 * @Author: John Holl
 * @Github: https://github.com/hzylyh
 * @Date: 2021-05-04 09:41:14
 * @LastEditors: John Holl
 * @LastEditTime: 2021-05-17 20:08:55
 */
package service

import (
	"aDeploy/conf"
	"aDeploy/utils/qjson"
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

func (srv *deploymentService) DeleteDeployment(qj *qjson.QJson) (err error) {
	name := qj.GetString("name")
	if err = conf.Clientset.AppsV1().Deployments("default").Delete(context.TODO(), name, metav1.DeleteOptions{}); err != nil {
		return err
	}
	return nil
}

//func (srv *demoService) DeleteDeployment(c *gin.Context) (deployment *v1.Deployment, err error) {
//}

func (srv *deploymentService) GetDeployment(c *gin.Context) (deployments *appsV1.DeploymentList, err error) {
	if deployments, err = conf.Clientset.AppsV1().Deployments("default").List(context.TODO(), metav1.ListOptions{}); err != nil {
		return nil, err
	}
	return deployments, nil
}
