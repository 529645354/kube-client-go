package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	apiv1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s/status"
)

func DeleteDeployment(c *gin.Context){
	namespace := c.Query("namespace")
	deployment := c.Query("deployment")
	err := clientset.AppsV1().Deployments(namespace).Delete(deployment, &v1.DeleteOptions{})
	if err!=nil{
		fmt.Println(err)
		c.JSON(200, gin.H{
			"status": status.InternalError,
		})
	}else {
		c.JSON(200,gin.H{
			"status": status.Ok,
		})
	}
}

func DeletePod(c *gin.Context){
	namespace := c.Query("namespace")
	pod := c.Query("pod")
	err := clientset.CoreV1().Pods(namespace).Delete(pod, &v1.DeleteOptions{})
	if err!=nil{
		fmt.Println(err)
		c.JSON(200, gin.H{
			"status": status.InternalError,
		})
	}else {
		c.JSON(200,gin.H{
			"status": status.Ok,
		})
	}
}

func DeleteService (c *gin.Context) {
	namespace := c.Query("namespace")
	service := c.Query("service")
	err := clientset.CoreV1().Services(namespace).Delete(service, &v1.DeleteOptions{})
	if err!=nil{
		fmt.Println(err)
		c.JSON(200, gin.H{
			"status": status.InternalError,
		})
	}else {
		c.JSON(200,gin.H{
			"status": status.Ok,
		})
	}
}

func DeleteIngress( c *gin.Context) {
	namespace := c.Query("namespace")
	ingress := c.Query("ingress")
	err := clientset.NetworkingV1beta1().Ingresses(namespace).Delete(ingress, &v1.DeleteOptions{})
	if err!=nil{
		fmt.Println(err)
		c.JSON(200, gin.H{
			"status": status.InternalError,
		})
	}else {
		c.JSON(200,gin.H{
			"status": status.Ok,
		})
	}
}

func DeletePv(c *gin.Context) {
	pv := c.Query("pv")
	err := clientset.CoreV1().PersistentVolumes().Delete(pv, &v1.DeleteOptions{})
	if err!=nil{
		fmt.Println(err)
		c.JSON(200, gin.H{
			"status": status.InternalError,
		})
	}else {
		c.JSON(200,gin.H{
			"status": status.Ok,
		})
	}
}

func DeletePvc(c *gin.Context) {
	namespace := c.Query("namespace")
	pvc := c.Query("pvc")
	err := clientset.CoreV1().PersistentVolumeClaims(namespace).Delete(pvc, &v1.DeleteOptions{})
	if err!=nil{
		fmt.Println(err)
		c.JSON(200, gin.H{
			"status": status.InternalError,
		})
	}else {
		c.JSON(200,gin.H{
			"status": status.Ok,
		})
	}
}

func DeleteConfigMap (c *gin.Context) {
	namespace := c.Query("namespace")
	ConfigMap := c.Query("configmap")
	err := clientset.CoreV1().ConfigMaps(namespace).Delete(ConfigMap, &v1.DeleteOptions{})
	if err!=nil{
		fmt.Println(err)
		c.JSON(200, gin.H{
			"status": status.InternalError,
		})
	}else {
		c.JSON(200,gin.H{
			"status": status.Ok,
		})
	}
}

func DeleteSecret(c *gin.Context){
	namespace := c.Query("namespace")
	Secret := c.Query("secret")
	err := clientset.CoreV1().Secrets(namespace).Delete(Secret, &v1.DeleteOptions{})
	if err!=nil{
		fmt.Println(err)
		c.JSON(200, gin.H{
			"status": status.InternalError,
		})
	}else {
		c.JSON(200,gin.H{
			"status": status.Ok,
		})
	}
}

func DeleteNamespace (c *gin.Context) {
	namespace := c.Query("name")
	get, err := clientset.CoreV1().Namespaces().Get(namespace, v1.GetOptions{})
	if err!=nil{
		fmt.Println(err)
		c.JSON(200, gin.H{
			"status": status.DataError,
		})
	}
	if err!=nil{
		fmt.Println(err)
		c.JSON(200, gin.H{
			"status": status.DataError,
		})
	}
	err = clientset.CoreV1().Namespaces().Delete(namespace, &v1.DeleteOptions{})
	get.Spec = apiv1.NamespaceSpec{nil}
	_, err = clientset.CoreV1().Namespaces().Update(get)
	if err!=nil{
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