package pod

import (
	apiv1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

func (p *Pod) VolumeCreate () []apiv1.Volume{
	volume := make([]apiv1.Volume,0)
	for _, i := range p.Volumes{
		volume = append(volume,apiv1.Volume{
			Name:         i.Name,
			VolumeSource: apiv1.VolumeSource{
				PersistentVolumeClaim: &apiv1.PersistentVolumeClaimVolumeSource{
					ClaimName: i.Pvc,
					ReadOnly:  false,
				},
			},
		})
	}
	for _, i := range p.NfsVolumes {
		volume = append(volume, apiv1.Volume{
			Name:         i.Name,
			VolumeSource: apiv1.VolumeSource{
				NFS: &apiv1.NFSVolumeSource{
					Server:  i.NfsServer,
					Path:     i.Path,
					ReadOnly: false,
				},
			},
		})
	}

	for _, i:= range p.ConfigMapVolumes {
		volume = append(volume, apiv1.Volume{
			Name:         i.Name,
			VolumeSource: apiv1.VolumeSource{
				ConfigMap: &apiv1.ConfigMapVolumeSource{
					LocalObjectReference: apiv1.LocalObjectReference{
						Name:i.ConfigMap,
					},
				},
			},
		})
	}

	for _, i:= range p.SecretVolumes {
		volume = append(volume, apiv1.Volume{
			Name:         i.Name,
			VolumeSource: apiv1.VolumeSource{
				Secret: &apiv1.SecretVolumeSource{
					SecretName:  i.Secret,
				},
			},
		})
	}
	return volume
}

func (p *Pod) ContainerCreate () []apiv1.Container {
	container := make([]apiv1.Container, 0)
	for _, i:= range p.Containers {
		con := apiv1.Container{
			Name:                     i.Name,
			Image:                    i.Image,
			Command:                  nil,
			Env:                      nil,
			VolumeMounts:             nil,
			ImagePullPolicy:          apiv1.PullIfNotPresent,
		}
		command := strings.Split(strings.Trim(i.Command," "), ",")
		if command[0] != "" {
			con.Command = command
		}
		env := make([]apiv1.EnvVar,0)
		for _, e := range i.Envs {
			env = append(env,apiv1.EnvVar{
				Name:      e.Name,
				Value:     e.Value,
				ValueFrom: nil,
			})
		}
		for _, e := range i.ConfigMapEnvs {
			env = append(env, apiv1.EnvVar{
				Name:      e.Name,
				ValueFrom: &apiv1.EnvVarSource{
					ConfigMapKeyRef:  &apiv1.ConfigMapKeySelector{
						LocalObjectReference: apiv1.LocalObjectReference{
							Name: e.ConfigMapName,
						},
						Key:                  e.Key,
					},
				},
			})
		}
		volumeMount := make([]apiv1.VolumeMount,0)
		for _,v := range i.VolumeMounts {
			volumeMount = append(volumeMount, apiv1.VolumeMount{
				Name:             v.Name,
				ReadOnly:         false,
				MountPath:        v.Path,
			})
		}
		con.VolumeMounts = volumeMount
		con.Env = env
		container = append(container, con)
	}
	return container
}

func (p *Pod) InitContainerCreate () []apiv1.Container {
	container := make([]apiv1.Container, 0)
	for _,i := range p.InitContainer {
		 initCon := apiv1.Container{
			Name:                     i.Name,
			Image:                    i.Image,
			Command:                  strings.Split(i.Command,","),
			Env:                      nil,
			VolumeMounts:             nil,
			ImagePullPolicy:          apiv1.PullIfNotPresent,
		}
		env := make([]apiv1.EnvVar, 0)
		for _,e := range i.Envs {
			env = append(env, apiv1.EnvVar{
				Name:      e.Name,
				Value:     e.Value,
			})
		}
		initCon.Env = env
		volumeMount := make([]apiv1.VolumeMount, 0)
		for _, v:= range i.VolumeMounts {
			volumeMount = append(volumeMount, apiv1.VolumeMount{
				Name:             v.Name,
				ReadOnly:         false,
				MountPath:        v.Path,
			})
		}
		initCon.VolumeMounts = volumeMount
		container = append(container, initCon)
	}
	return container
}

func (p *Pod) CreatePod () *apiv1.Pod {
	label := make(map[string]string,0)
	for _, i := range p.Labels{
		label[i.Key] = i.Value
	}
	pod := apiv1.Pod{
		ObjectMeta: v1.ObjectMeta{Name:p.Name,Labels:label},
		Spec:       apiv1.PodSpec{
			Containers: p.ContainerCreate(),
			Volumes: p.VolumeCreate(),
			InitContainers: p.InitContainerCreate(),
		},
		Status:     apiv1.PodStatus{},
	}
	return &pod
}
