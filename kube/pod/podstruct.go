package pod

type Label struct {
	Key string `json:"key"`
	Value string `json:"value"`
}
type Volume struct {
	Name string `json:"name"`
	Pvc string `json:"pvc"`
}

type NfsVolume struct {
	NfsServer string `json:"nfsserver"`
	Path string `json:"path"`
	Name string `json:"name"`
}

type SecretVolume struct {
	Name string `json:"name"`
	Secret string `json:"secret"`
}
type ConfigMapVolume struct {
	Name string `json:"name"`
	ConfigMap string `json:"secretandconfigmap"`
}

type Env struct {
	Name string `json:"name"`
	Value string `json:"value"`
}

type ConfigMap struct {
	Name string `json:"name"`
	Key string `json:"key"`
	ConfigMapName string `json:"configname"`
}

type VolumeMount struct {
	Name string `json:"name"`
	Path string `json:"path"`
}
type Container struct {
	Command string `json:"Command"`
	Image string `json:"Image"`
	Name string `json:"Name"`
	Envs []Env `json:"Env"`
	ConfigMapEnvs []ConfigMap `json:"ConfigMap"`
	VolumeMounts []VolumeMount `json:"VolumeMount"`
}


type Pod struct {
	Labels []Label `json:"label"`
	Namespace string `json:"namespace"`
	Name string `json:"name"`
	Volumes []Volume `json:"volume"`
	Containers []Container `json:"container"`
	InitContainer []Container `json:"initContainer"`
	NfsVolumes []NfsVolume `json:"nfsvolume"`
	SecretVolumes []SecretVolume `json:"secretvolume"`
	ConfigMapVolumes []ConfigMapVolume `json:"configmapvolume"`
}