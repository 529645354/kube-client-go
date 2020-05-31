package ingress

type Rule struct {
	BackendPort int32 `json:"backendPort"`
	BackendService string `json:"backendService"`
	Path string `json:"path"`
}

type Host struct {
	Hostname string `json:"hostname"`
	HttpLs []Rule `json:"http"`
}

type Ingress struct {
	Name string `json:"name"`
	Namespace string `json:"namespace"`
	Hosts []Host `json:"host"`
}