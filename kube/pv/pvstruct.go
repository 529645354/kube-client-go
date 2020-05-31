package pv

import (
	v1 "k8s.io/api/core/v1"
	"k8s/kube/pod"
)


type Pv struct {
	Name string `json:"name"`
	Policy v1.PersistentVolumeReclaimPolicy `json:"policy"`
	Storage int64 `json:"storage"`
	AccessMode []v1.PersistentVolumeAccessMode `json:"accessmode"`
	StorageClassName string `json:"storageClassName"`
	Nfs pod.NfsVolume `json:"nfs"`
}

type Pvc struct {
	Name string `json:"name"`
	Namespace string `json:"namespace"`
	Storage int64 `json:"storage"`
	AccessMode []v1.PersistentVolumeAccessMode `json:"accessmode"`
	StorageClassName string `json:"storageClassName"`
} 