package client

import (
	"flag"
	"k8s-go-client/logs"
	"k8s-go-client/vo"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var clientset *kubernetes.Clientset

func K8sInit() {

	k8sconfig := flag.String("k8sconfig", "./config", "kubernetes config file path")
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *k8sconfig)
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		logs.GetLogger().Error(err)
		return
	} else {
		logs.GetLogger().Info("Client initiated successfully!")
	}
}

func getK8sNodeList() (*apiv1.NodeList, error) {

	nodes, err := clientset.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return nodes, nil
}

func getK8sPodList() (*apiv1.PodList, error) {

	pods, err := clientset.CoreV1().Pods("default").List(metav1.ListOptions{})
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return pods, nil
}

func createK8sDeployment(config vo.DeploymentConfig) (*appsv1.Deployment, error) {

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: config.Name,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &config.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: config.Labels,
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: config.Labels,
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  config.ContainerName,
							Image: config.ImageName,
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: config.ContainerPort,
								},
							},
						},
					},
				},
			},
		},
	}

	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
	result, err := deploymentsClient.Create(deployment)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return result, nil
}

func deleteK8sDeployment(deploymentName string) error {

	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
	deletePolicy := metav1.DeletePropagationForeground
	if err := deploymentsClient.Delete(deploymentName, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}

func createK8sService(config vo.ServiceConfig) (int32, error) {

	service := &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: config.Name,
		},
		Spec: apiv1.ServiceSpec{
			Type:     apiv1.ServiceTypeNodePort,
			Selector: config.Labels,
			Ports: []apiv1.ServicePort{
				{
					Port: config.Port,
				},
			},
		},
		Status: apiv1.ServiceStatus{},
	}

	servicesClient := clientset.CoreV1().Services(apiv1.NamespaceDefault)

	result, err := servicesClient.Create(service)
	if err != nil {
		logs.GetLogger().Error(err)
		return 0, err
	}

	return result.Spec.Ports[0].NodePort, nil
}

func deleteK8sService(serviceName string) error {

	servicesClient := clientset.CoreV1().Services(apiv1.NamespaceDefault)
	deletePolicy := metav1.DeletePropagationForeground
	if err := servicesClient.Delete(serviceName, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}
