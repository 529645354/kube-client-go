package deployment

import (
	appsv1 "k8s.io/api/apps/v1"
	v12 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (d *Deploy) Seletor() map[string]string {
	label := make(map[string]string)
	for _, i := range d.Pod.Labels {
		label[i.Key] = i.Value
	}
	return label
}

func (d *Deploy) CreateDeploy() *appsv1.Deployment {
	deployment := appsv1.Deployment{
		ObjectMeta: v1.ObjectMeta{
			Name:      d.Name,
			Namespace: d.Pod.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &d.Replicas,
			Selector: &v1.LabelSelector{
				MatchLabels: d.Seletor(),
			},
			Template: v12.PodTemplateSpec{
				ObjectMeta: v1.ObjectMeta{
					Labels: d.Seletor(),
				},
				Spec: v12.PodSpec{
					Containers:     d.ContainerCreate(),
					Volumes:        d.VolumeCreate(),
					InitContainers: d.InitContainerCreate(),
				},
			},
		},
	}
	label := make(map[string]string,0)
	for _,i := range d.DeployLabel {
		label[i.Key] = i.Value
	}
	deployment.Labels = label
	return &deployment
}
