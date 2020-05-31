package deployment

import "k8s/kube/pod"

type Deploy struct {
	Name string `json:"name"`
	DeployLabel []pod.Label `json:"deploylabel"`
	Replicas int32 `json:"replicas"`
	pod.Pod
}
