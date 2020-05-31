package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	apiv1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s/kube/deployment"
	"k8s/kube/ingress"
	"k8s/kube/pod"
	"k8s/kube/pv"
	"k8s/kube/secretandconfigmap"
	"k8s/kube/service"
	"k8s/status"
)

func CreatePod(c *gin.Context) {
	var data pod.Pod
	if err := c.ShouldBindJSON(&data); err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"status": status.DataError,
		})
		return
	}
	createPod := data.CreatePod()
	_, err := clientset.CoreV1().Pods(data.Namespace).Create(createPod)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"status": status.InternalError,
		})
		return
	}
}

func CreateDeployment(c *gin.Context) {
	var data deployment.Deploy
	if err := c.ShouldBindJSON(&data); err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"status": status.InternalError,
		})
		return
	}
	deploy := data.CreateDeploy()
	_, err := clientset.AppsV1().Deployments(data.Namespace).Create(deploy)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"status": status.DataError,
		})
		return
	}
	fmt.Println(data)
	c.JSON(200, gin.H{
		"status": status.Ok,
	})
}

func CreateSvc(c *gin.Context) {
	var data service.Service
	if err := c.ShouldBindJSON(&data); err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"status": status.DataError,
		})
	}
	fmt.Println(data)
	createService := data.CreateService()
	_, err := clientset.CoreV1().Services(data.Namespace).Create(createService)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"status": status.InternalError,
		})
		return
	}
	c.JSON(200, gin.H{
		"status": status.Ok,
	})
}

func CreateIngress(c *gin.Context) {
	var data ingress.Ingress
	if err := c.ShouldBindJSON(&data); err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"status": status.DataError,
		})
		return
	}
	createIngress := data.CreateIngress()
	_, err := clientset.NetworkingV1beta1().Ingresses(data.Namespace).Create(createIngress)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"status": status.InternalError,
		})
		return
	}
	c.JSON(200, gin.H{
		"status": status.Ok,
	})
}

func CreatePv(c *gin.Context) {
	var data pv.Pv
	if err := c.ShouldBindJSON(&data); err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"status": status.DataError,
		})
		return
	}
	createPv := data.CreatePv()
	_, err := clientset.CoreV1().PersistentVolumes().Create(createPv)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"status": status.DataError,
		})
		return
	}
	c.JSON(200, gin.H{
		"status": status.Ok,
	})
}

func CreatePvc(c *gin.Context) {
	var data pv.Pvc
	if err := c.ShouldBindJSON(&data); err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"status": status.DataError,
		})
		return
	}
	pvc := data.CreatePvc()
	_, err := clientset.CoreV1().PersistentVolumeClaims(data.Namespace).Create(pvc)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"status": status.DataError,
		})
		return
	}
	c.JSON(
		200, gin.H{
			"status": status.Ok,
		})
}

func CreateConfigMap(c *gin.Context) {
	var data secretandconfigmap.ConfigMap
	if err := c.ShouldBindJSON(&data); err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"status": status.DataError,
		})
		return
	}
	configMap := data.CreateConfigMap()
	_, err := clientset.CoreV1().ConfigMaps(data.Namespace).Create(configMap)
	if err!=nil{
		fmt.Println(err)
		c.JSON(200, gin.H{
			"status": status.DataError,
		})
		return
	}
	c.JSON(
		200, gin.H{
			"status": status.Ok,
		})
}

func CreateSecret(c *gin.Context) {
	var data secretandconfigmap.Secret
	if err := c.ShouldBindJSON(&data); err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"status": status.DataError,
		})
		return
	}
	secret, err := data.CreateSecret()
	if err!=nil {
		fmt.Println(err)
		c.JSON(200,gin.H{
			"status": status.DataError,
		})
		return
	}
	_, err = clientset.CoreV1().Secrets(data.Namespace).Create(secret)
	if err!=nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"status": status.DataError,
		})
		return
	}
	c.JSON(200, gin.H{
		"status": status.Ok,
	})
}

func CreateNamespace(c *gin.Context) {
	type NameSpace struct {
		Name string `json:"name"`
	}
	var n NameSpace
	if err := c.ShouldBindJSON(&n); err!=nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"status":status.DataError,
		})
		return
	}
	namespace := apiv1.Namespace{
		ObjectMeta: v1.ObjectMeta{
			Name: n.Name,
		},
	}
	_, err := clientset.CoreV1().Namespaces().Create(&namespace)
	if err!=nil{
		fmt.Println(err)
		c.JSON(200, gin.H{
			"status":status.DataError,
		})
		return
	}
	c.JSON(200, gin.H{
		"status": status.Ok,
	})
}