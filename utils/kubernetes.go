package utils

import (
	"context"
	"flag"
	"fmt"
	v1 "k8s.io/api/apps/v1"
	v12 "k8s.io/api/core/v1"
	//v1beta1 "k8s.io/api/apps/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	//"k8s.io/client-go/util/homedir"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
)

func main() {
	var kubeconfig *string
	//if home := homedir.HomeDir(); home != "" {
	//	kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	//} else {
	//	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	//}
	kubeconfig = flag.String("kubeconfig", "./config", "absolute path to the kubeconfig file")
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	//for {
	//	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	//	if err != nil {
	//		panic(err.Error())
	//	}
	//	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	//
	//	// Examples for error handling:
	//	// - Use helper functions like e.g. errors.IsNotFound()
	//	// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
	//	namespace := "default"
	//	pod := "example-xxxxx"
	//	_, err = clientset.CoreV1().Pods(namespace).Get(context.TODO(), pod, metav1.GetOptions{})
	//	if errors.IsNotFound(err) {
	//		fmt.Printf("Pod %s in namespace %s not found\n", pod, namespace)
	//	} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
	//		fmt.Printf("Error getting pod %s in namespace %s: %v\n",
	//			pod, namespace, statusError.ErrStatus.Message)
	//	} else if err != nil {
	//		panic(err.Error())
	//	} else {
	//		fmt.Printf("Found pod %s in namespace %s\n", pod, namespace)
	//	}
	//
	//	time.Sleep(10 * time.Second)
	//}
	createDeployment(clientset)
	//deleteDeployment(clientset)

}

func deleteDeployment(clientset *kubernetes.Clientset)  {
	if err := clientset.AppsV1().Deployments("default").Delete(context.TODO(), "testdeploy", metav1.DeleteOptions{}); err != nil {
		fmt.Println(err)
	}
}

func createDeployment(clientset *kubernetes.Clientset)  {
	var replicas int32 = 1
	var deploy = &v1.Deployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:                       "testdeploy",
			GenerateName:               "",
			Namespace:                  "",
			SelfLink:                   "",
			UID:                        "",
			ResourceVersion:            "",
			Generation:                 0,
			CreationTimestamp:          metav1.Time{},
			DeletionTimestamp:          nil,
			DeletionGracePeriodSeconds: nil,
			Labels: map[string]string{"app": "ngx"},
			Annotations:                nil,
			OwnerReferences:            nil,
			Finalizers:                 nil,
			ClusterName:                "",
			ManagedFields:              nil,
		},
		Spec: v1.DeploymentSpec{
			Replicas:                &replicas,
			Selector:                &metav1.LabelSelector{
				MatchLabels: map[string]string{"app": "ngx"},
				MatchExpressions: nil,
			},
			Template:                v12.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name:                       "",
					GenerateName:               "",
					Namespace:                  "",
					SelfLink:                   "",
					UID:                        "",
					ResourceVersion:            "",
					Generation:                 0,
					CreationTimestamp:          metav1.Time{},
					DeletionTimestamp:          nil,
					DeletionGracePeriodSeconds: nil,
					Labels: map[string]string{"app": "ngx"},
					Annotations:                nil,
					OwnerReferences:            nil,
					Finalizers:                 nil,
					ClusterName:                "",
					ManagedFields:              nil,
				},
				Spec:       v12.PodSpec{
					Volumes:                       nil,
					InitContainers:                nil,
					Containers:                    []v12.Container{{
						Name:                     "nginx",
						Image:                    "nginx:1.8",
						Command:                  nil,
						Args:                     nil,
						WorkingDir:               "",
						Ports:                    nil,
						EnvFrom:                  nil,
						Env:                      nil,
						Resources:                v12.ResourceRequirements{},
						VolumeMounts:             nil,
						VolumeDevices:            nil,
						LivenessProbe:            nil,
						ReadinessProbe:           nil,
						StartupProbe:             nil,
						Lifecycle:                nil,
						TerminationMessagePath:   "",
						TerminationMessagePolicy: "",
						ImagePullPolicy:          "",
						SecurityContext:          nil,
						Stdin:                    false,
						StdinOnce:                false,
						TTY:                      false,
					}},
					EphemeralContainers:           nil,
					RestartPolicy:                 "",
					TerminationGracePeriodSeconds: nil,
					ActiveDeadlineSeconds:         nil,
					DNSPolicy:                     "",
					NodeSelector:                  nil,
					ServiceAccountName:            "",
					DeprecatedServiceAccount:      "",
					AutomountServiceAccountToken:  nil,
					NodeName:                      "",
					HostNetwork:                   false,
					HostPID:                       false,
					HostIPC:                       false,
					ShareProcessNamespace:         nil,
					SecurityContext:               nil,
					ImagePullSecrets:              nil,
					Hostname:                      "",
					Subdomain:                     "",
					Affinity:                      nil,
					SchedulerName:                 "",
					Tolerations:                   nil,
					HostAliases:                   nil,
					PriorityClassName:             "",
					Priority:                      nil,
					DNSConfig:                     nil,
					ReadinessGates:                nil,
					RuntimeClassName:              nil,
					EnableServiceLinks:            nil,
					PreemptionPolicy:              nil,
					Overhead:                      nil,
					TopologySpreadConstraints:     nil,
					SetHostnameAsFQDN:             nil,
				},
			},
			Strategy:                v1.DeploymentStrategy{},
			MinReadySeconds:         0,
			RevisionHistoryLimit:    nil,
			Paused:                  false,
			ProgressDeadlineSeconds: nil,
		},
		Status: v1.DeploymentStatus{
			ObservedGeneration:  0,
			Replicas:            0,
			UpdatedReplicas:     0,
			ReadyReplicas:       0,
			AvailableReplicas:   0,
			UnavailableReplicas: 0,
			Conditions:          nil,
			CollisionCount:      nil,
		},
	}
	if _, err := clientset.AppsV1().Deployments("default").Create(context.TODO(), deploy, metav1.CreateOptions{
		TypeMeta:     metav1.TypeMeta{},
		DryRun:       nil,
		FieldManager: "",
	}); err != nil {
		fmt.Println(err)
	}
}