package secretandconfigmap

type Datas struct {
	Key string `json:"key"`
	Value string `json:"value"`
}

type ConfigMap struct {
	Name string `json:"name"`
	Namespace string `json:"namespace"`
	Data []Datas `json:"data"`
}