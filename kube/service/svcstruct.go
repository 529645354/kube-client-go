package service

type Selectors struct {
	Key string `json:"key"`
	Value string `json:"value"`
}

type Ports struct {
	Port int32 `json:"port"`
	Protocol string `json:"protocol"`
	TargetPort int32 `json:"targetPort"`
}

type ClusterIP struct {

	Port []Ports `json:"ports"`
	Headless bool `json:"headless"`
}

type NodePort struct {
	Port []Ports `json:"ports"`
}

type ExternalName struct {
	ExternalNameDNS string `json:"externalname"`
}

type Service struct {
	Name string `json:"name"`
	Namespace string `json:"namespace"`
	Selector []Selectors `json:"selector"`
	Type string `json:"type"`
	ClusterIP `json:"clusterip"`
	NodePort `json:"nodeport"`
	ExternalName
}