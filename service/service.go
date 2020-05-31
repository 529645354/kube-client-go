package service

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	"net/http"
	//appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s/status"
)

var (
	clientset *kubernetes.Clientset
	kubeconfig *string
)

func Init() {
	kubeconfig = flag.String("kubeconfig","./config","not foud")
	flag.Parse()
	config , err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err!=nil{
		panic(err)
	}
	clientset, err = kubernetes.NewForConfig(config)
	if err!=nil{
		panic(err)
	}
}


func GetNamespaces (c *gin.Context) {
	list, err := clientset.CoreV1().Namespaces().List(v1.ListOptions{})
	if err!=nil{
		c.JSON(200,gin.H{
			"status":status.InternalError,
		})
		return
	}
	var n []string
	for _,i := range list.Items{
			n = append(n,i.Name)
	}
	c.JSON(200,gin.H{
		"namespace": n,
		"status": status.Ok,
	})
}

func DeploymentList(c *gin.Context,) {
	type DeploymentList struct {
		Name string `json:"name"`
 		Replicas int32 `json:"replicas"`
		Label string `json:"label"`
		Ready string `json:"ready"`
		Selector string `json:"selector"`
	}
	namespace := c.Query("namespace")
	fmt.Println(namespace)
	list, err := clientset.AppsV1().Deployments(namespace).List(v1.ListOptions{})
	if err!=nil{
		c.JSON(200, gin.H{
			"status": status.InternalError,
		})
		return
	}
	var deploy DeploymentList
	var d []DeploymentList
	for _,i := range list.Items {
		deploy.Name = i.Name
		marshal, _ := json.Marshal(i.Labels)
		deploy.Label = string(marshal)
		ready := make(map[int32]int32)
		ready[i.Status.ReadyReplicas] = i.Status.Replicas
		r, _ := json.Marshal(ready)
		deploy.Ready = string(r)
		selector, _ := json.Marshal(i.Spec.Selector.MatchLabels)
		deploy.Selector = string(selector)
		if err!=nil{
			c.JSON(200, gin.H{
				"status": status.InternalError,
			})
		}
		d = append(d,deploy)
	}
	c.JSON(200,gin.H{
		"status": status.Ok,
		"content": d,
	})
}

func PodList(c *gin.Context) {
	type PodList struct {
		Name string `json:"name"`
		HostIP string `json:"host_ip"`
		Label string `json:"label"`
		NodeName string `json:"node_name"`
		PodIP string `json:"pod_ip"`
		PodStatus string `json:"pod_status"`
	}
	namespace := c.Query("namespace")
	list, err := clientset.CoreV1().Pods(namespace).List(v1.ListOptions{})
	if err!=nil{
		c.JSON(200,gin.H{
			"status": status.InternalError,
		})
		return
	}
	var pl []PodList
	for _,i := range list.Items {
		var p PodList
		p.HostIP = i.Status.HostIP
		p.Name = i.Name
		p.NodeName = i.Spec.NodeName
		marshal, _ := json.Marshal(i.Labels)
		p.Label = string(marshal)
		p.PodIP = i.Status.PodIP
		p.PodStatus = string(i.Status.Phase)
		pl = append(pl, p)
	}
	c.JSON(http.StatusOK,gin.H{
		"status": status.Ok,
		"content": pl,
	})
}

func GetPod(c *gin.Context) {
	podname := c.Query("pod")
	namespace := c.Query("namespace")
	get, err := clientset.CoreV1().Pods(namespace).Get(podname, v1.GetOptions{})
	if err!=nil{
		fmt.Println(err)
		c.JSON(200,gin.H{
			"status": status.InternalError,
		})
		return
	}
	c.JSON(200, gin.H{
		"status": status.Ok,
		"content": get.Spec,
	})
}

func ServiceList(c *gin.Context) {
	type Service struct {
		Name string `json:"name"`
		Selector string `json:"selector"`
		Type string `json:"type"`
		ClusterIP string `json:"cluster_ip"`
		Port map[int32]map[string]int32 `json:"port"`
		ExternalName string `json:"external_name"`
		Endpoints []string `json:"endpoints"`
	}
	namespace := c.Query("namespace")
	list, err := clientset.CoreV1().Services(namespace).List(v1.ListOptions{})
	if err!=nil{
		fmt.Println(err.Error())
		c.JSON(200,gin.H{
			"status":status.InternalError,
		})
		return
	}
	endpointsList, err := clientset.CoreV1().Endpoints(namespace).List(v1.ListOptions{})
	if err!=nil{
		fmt.Println(err.Error())
		c.JSON(200,gin.H{
			"status":status.InternalError,
		})
	}
	var service []Service
	for _,i := range list.Items {
		marshal, _ := json.Marshal(i.Spec.Selector)
		s := Service{
			Name:      i.Name,
			Selector:  string(marshal),
			Type:      string(i.Spec.Type),
			ClusterIP: i.Spec.ClusterIP,
			Port:      make(map[int32]map[string]int32),
			ExternalName: i.Spec.ExternalName,
			Endpoints: make([]string,0),
		}
		if endpointsList != nil {
			for _,e := range endpointsList.Items {
				if e.Name == i.Name {
					for _,addr := range e.Subsets {
						for _,a := range addr.Addresses {
							s.Endpoints = append(s.Endpoints, a.IP)
						}
					}
					break
				}
			}
		}
		for _,p := range i.Spec.Ports{
			m := make(map[string]int32, 0)
			m["主机节点端口"] = p.NodePort
			m["目标容器内暴露端口"] = p.TargetPort.IntVal
			s.Port[p.Port] = m
		}
		service = append(service, s)
	}
	c.JSON(200, gin.H{
		"status": status.Ok,
		"content": service,
	})
}

func IngressList(c *gin.Context){
	type Backend struct {
		Path string `json:"path"`
		ServiceName string `json:"service_name"`
		ServicePort int32 `json:"service_port"`
	}
	type Rule struct {
		Host string `json:"host"`
		Backends []Backend `json:"backends"`
	}
	type Ingress struct {
		Name string `json:"name"`
		Rules []Rule `json:"rules"`
	}
	namespace := c.Query("namespace")
	ingresslist, err := clientset.NetworkingV1beta1().Ingresses(namespace).List(v1.ListOptions{})
	if err!=nil{
		fmt.Println(err)
		c.JSON(200,gin.H{"status": status.InternalError})
		return
	}
	var ingressAll []Ingress
	for _,i := range ingresslist.Items {
		ingress := Ingress{
			Name:    i.Name,
			Rules: make([]Rule,0),
		}
		for _,r := range i.Spec.Rules{
			rules := Rule{
				Host:        r.Host,
				Backends:        make([]Backend,0),
			}
			for _,rule := range r.HTTP.Paths {
				backend := Backend{
					Path: rule.Path,
					ServiceName: rule.Backend.ServiceName,
					ServicePort: rule.Backend.ServicePort.IntVal,
				}
				rules.Backends = append(rules.Backends,backend)
			}
			ingress.Rules = append(ingress.Rules,rules)
		}
		ingressAll = append(ingressAll,ingress)
	}
	c.JSON(200,gin.H{
		"status": status.Ok,
		"content": ingressAll,
	})
}

func PvList(c *gin.Context){
	type Pv struct {
		Name string `json:"name"`
		Status string `json:"status"`
		Capacity string `json:"capacity"`
		Policy string `json:"policy"`
		StorageClass string `json:"storage_class"`
		AccessMode []corev1.PersistentVolumeAccessMode `json:"access_mode"`
	}
	list, err := clientset.CoreV1().PersistentVolumes().List(v1.ListOptions{})
	if err!=nil{
		fmt.Println(err)
		c.JSON(200,gin.H{
			"status": status.InternalError,
		})
		return
	}
	var pv []Pv
	for _,i := range list.Items {
		p := Pv{
			Name:         i.Name,
			Status:       string(i.Status.Phase),
			Capacity:     i.Spec.Capacity.Storage().String(),
			Policy:       string(i.Spec.PersistentVolumeReclaimPolicy),
			StorageClass: i.Spec.StorageClassName,
			AccessMode:   i.Spec.AccessModes,
		}
		pv = append(pv, p)
	}
	c.JSON(200, gin.H{
		"status": status.Ok,
		"content": pv,
	})
}

func GetPVC(c *gin.Context) {
	type Pvc struct {
		Name string `json:"name"`
		Status string `json:"status"`
		VolumeName string `json:"volume_name"`
		Capacity string `json:"capacity"`
		AccessMode []corev1.PersistentVolumeAccessMode `json:"access_mode"`
		StorageClass *string `json:"storage_class"`
	} 
	namespace := c.Query("namespace")
	list, err := clientset.CoreV1().PersistentVolumeClaims(namespace).List(v1.ListOptions{})
	if err!=nil{
		fmt.Println(err)
		c.JSON(200,gin.H{"status": status.Ok})
		return
	}
	var pvc []Pvc
	for _,i := range list.Items {
		p := Pvc{
			Name:         i.Name,
			Status:       string(i.Status.Phase),
			VolumeName:       i.Spec.VolumeName,
			Capacity:     i.Status.Capacity.Storage().String(),
			AccessMode:   i.Spec.AccessModes,
			StorageClass: i.Spec.StorageClassName,
		}
		pvc = append(pvc,p)
	}
	c.JSON(200,gin.H{
		"status": status.Ok,
		"content": pvc,
	})
}

func ConfigMapList(c *gin.Context){
	type ConfigMap struct {
		Name string `json:"name"`
		Data map[string]string `json:"data"`
	}
	namespace := c.Query("namespace")
	list, err := clientset.CoreV1().ConfigMaps(namespace).List(v1.ListOptions{})
	if err!=nil{
		fmt.Println(err)
		c.JSON(200,gin.H{"status": status.InternalError})
		return
	}
	var configmap []ConfigMap
	for _,i := range list.Items {
		c := ConfigMap{
			Name: i.Name,
			Data: i.Data,
		}
		configmap = append(configmap, c)
	}
	c.JSON(200, gin.H{
		"status": status.Ok,
		"content": configmap,
	})
}

func SecretList (c *gin.Context){
	type Secret struct {
		Name string `json:"name"`
		Data map[string][]byte `json:"data"`
		Type string `json:"type"`
	}
	namepsace := c.Query("namespace")
	secretList, err := clientset.CoreV1().Secrets(namepsace).List(v1.ListOptions{})
	if err!=nil{
		fmt.Println(err)
		c.JSON(200,gin.H{"status": status.InternalError})
		return
	}
	var s []Secret
	for _,i := range secretList.Items {
		secret := Secret{
			Name: i.Name,
			Data: i.Data,
			Type: string(i.Type),
		}
	s = append(s, secret)
	}
	c.JSON(200,gin.H{
		"status": status.Ok,
		"content": s,
	})
}

func NodeList (c *gin.Context){
	type Addre struct {
		Type string `json:"type"`
		Ip string `json:"ip"`
	}
	type Node struct {
		Name string `json:"name"`
		Cpu int64 `json:"cpu"`
		PodCIDR string `json:"pod_cidr"`
		Sched bool `json:"sched"`
		Net []Addre `json:"net"`
		TotalMemory int64 `json:"total_memory"`
	}
	nodeList, err := clientset.CoreV1().Nodes().List(v1.ListOptions{})
	if err!=nil{
		fmt.Println(err)
		c.JSON(200, gin.H{
			"status":status.InternalError,
		})
		return
	}
	var n []Node
	for _,i := range nodeList.Items {

		node := Node{
			Name:     i.Name,
			Cpu: i.Status.Allocatable.Cpu().Value(),
			PodCIDR: i.Spec.PodCIDR,
			Sched: i.Spec.Unschedulable,
			Net: make([]Addre,0),
			TotalMemory: i.Status.Capacity.Memory().Value()/1024/1024,
		}

		for _,a := range i.Status.Addresses{
			addr := Addre{
				Type: string(a.Type),
				Ip:   a.Address,
			}
			node.Net = append(node.Net, addr)
		}
		n = append(n ,node)
	}
	c.JSON(200,gin.H{
		"status": status.Ok,
		"content": n,
	})
}

func EventList(c *gin.Context){
	type Event struct {
		Name string `json:"name"`
		Type string `json:"type"`
		Reason string `json:"reason"`
		Message string `json:"message"`
		EventType string `json:"event_type"`
		EventRegardingName string `json:"event_regarding_name"`
	}
	namespace := c.Query("namespace")
	list, err := clientset.EventsV1beta1().Events(namespace).List(v1.ListOptions{})
	if err!=nil{
		fmt.Println(err)
		c.JSON(200,gin.H{
			"status": status.InternalError,
		})
		return
	}
	var e []Event
	for _,i := range list.Items {
		event := Event{
			Name:               i.Name,
			Type:               i.Type,
			Reason:             i.Reason,
			Message:            i.Note,
			EventType:          i.Regarding.Kind,
			EventRegardingName: i.Regarding.Name,
		}
		e = append(e, event)
	}
	c.JSON(200,gin.H{
		"status": status.Ok,
		"content": e,
	})
}