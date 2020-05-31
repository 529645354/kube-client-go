package pv

import (
	"fmt"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (p *Pv) CreatePv() *apiv1.PersistentVolume {
	fmt.Println(p.Nfs)
	pv := apiv1.PersistentVolume{
		ObjectMeta: v1.ObjectMeta{
			Name: p.Name,
		},
		Spec:       apiv1.PersistentVolumeSpec{
			AccessModes: p.AccessMode,
			PersistentVolumeSource: apiv1.PersistentVolumeSource{
				NFS: &apiv1.NFSVolumeSource{
					Server: p.Nfs.NfsServer,
					Path: p.Nfs.Path,
				},
			},
			PersistentVolumeReclaimPolicy: p.Policy,
			Capacity: apiv1.ResourceList{

			},
		},
	}
	r := resource.Quantity{}
	r.Format = resource.BinarySI
	r.Set(p.Storage * 1024 * 1024 * 1024)
	pv.Spec.Capacity[apiv1.ResourceStorage] = r

	fmt.Println(pv.Spec.Capacity.Storage().String())
	return &pv
}

func (pvc *Pvc) CreatePvc() *apiv1.PersistentVolumeClaim {
	p := apiv1.PersistentVolumeClaim{
		ObjectMeta: v1.ObjectMeta{
			Name: pvc.Name,
		},
		Spec:       apiv1.PersistentVolumeClaimSpec{
			AccessModes: pvc.AccessMode,
			Resources: apiv1.ResourceRequirements{
				Requests: apiv1.ResourceList{

				},
			},
		},
	}
	r := resource.Quantity{}
	r.Format = resource.BinarySI
	r.Set(pvc.Storage * 1024 * 1024 * 1024)
	p.Spec.Resources.Requests[apiv1.ResourceStorage] = r
	return &p
}