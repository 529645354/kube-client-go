package secretandconfigmap

type Tls struct {
	CertPath string `json:"certpath"`
	KeyPath string `json:"keypath"`
} 

type Generic struct {
	Key string `json:"key"`
	Value string `json:"value"`
}

type DockerRegistry struct {
	DockerServer string `json:"dockerserver"`
	DockerUsername string `json:"dockerusername"`
	DockerPassword string `json:"dockerpassword"`
	DockerEmail string `json:"dockeremail"`
}

type Secret struct {
	Name string `json:"name"`
	Stype string `json:"stype"`
	Namespace string `json:"namespace"`
	Tls  `json:"tls"`
	Generic []Generic `json:"generic"`
	DockerRegistry `json:"dockerregistry"`
} 