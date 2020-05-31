package main

import (
	"github.com/gin-gonic/gin"
	"k8s/service"
)


func main() {
	service.Init()
	r := gin.Default()
	group := r.Group("/k8s")
	{
		group.GET("/namespace", service.GetNamespaces)
		group.GET("/deployment", service.DeploymentList)
		group.GET("/podlist", service.PodList)
		group.GET("/getpod", service.GetPod)
		group.GET("/servicelist", service.ServiceList)
		group.GET("/ingresslist", service.IngressList)
		group.GET("/pvlist", service.PvList)
		group.GET("/pvclist", service.GetPVC)
		group.GET("/configmaplist", service.ConfigMapList)
		group.GET("/secretlist", service.SecretList)
		group.GET("/nodelist", service.NodeList)
		group.GET("/eventlist", service.EventList)
	}
	{
		group.DELETE("/deletedeploy", service.DeleteDeployment)
		group.DELETE("/deletepod", service.DeletePod)
		group.DELETE("/deleteservice", service.DeleteService)
		group.DELETE("/deleteingress", service.DeleteIngress)
		group.DELETE("/deletepv", service.DeletePv)
		group.DELETE("/detelepvc", service.DeletePvc)
		group.DELETE("/deleteconfigmap", service.DeleteConfigMap)
		group.DELETE("/deletesecret", service.DeleteSecret)
		group.DELETE("/deletenamespace", service.DeleteNamespace)
	}
	{
		group.POST("/createpod", service.CreatePod)
		group.POST("/createdeploy", service.CreateDeployment)
		group.POST("/createsvc", service.CreateSvc)
		group.POST("/createingress", service.CreateIngress)
		group.POST("/createpv", service.CreatePv)
		group.POST("/createpvc", service.CreatePvc)
		group.POST("/createconfigmap",service.CreateConfigMap)
		group.POST("/createsecret", service.CreateSecret)
		group.POST("/createnamespace", service.CreateNamespace)
	}
	r.Run(":9001")
}
