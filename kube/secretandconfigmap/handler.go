package secretandconfigmap

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	apiv1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *ConfigMap) CreateConfigMap() *apiv1.ConfigMap {
	configmap := apiv1.ConfigMap{
		ObjectMeta: v1.ObjectMeta{
			Name: c.Name,
		},
		Data:       make(map[string]string, 0),
	}
	for _, i:= range c.Data {
		configmap.Data[i.Key] = i.Value
	}
	return &configmap
}

func (s *Secret) CreateSecret() (*apiv1.Secret,error) {
	secret := apiv1.Secret{
		ObjectMeta: v1.ObjectMeta{
			Name: s.Name,
		},
		StringData: make(map[string]string, 0),
		Type:       "",
	}
	if s.Stype == "generic" {
		for _, i:= range s.Generic {
			secret.StringData[i.Key] = i.Value
		}
		return &secret, nil
	} else if s.Stype == "tls" {
		secret.Type = apiv1.SecretTypeTLS
		secret.StringData["tls.crt"] = s.Tls.CertPath
		secret.StringData["tls.key"] = s.Tls.KeyPath
		return &secret, nil
	} else {
		secret.Type = apiv1.SecretTypeDockerConfigJson

		type Auths struct {
			Auth map[string]map[string]string `json:"auths"`
		}
		ah := make(map[string]string, 0)
		ah["username"] = s.DockerUsername
		ah["password"] = s.DockerPassword
		ah["email"] = s.DockerEmail
		ah["auth"] = base64.StdEncoding.EncodeToString([]byte(s.DockerUsername + ":" + s.DockerPassword))
		reg := make(map[string]map[string]string,0)
		reg[s.DockerServer] = ah
		a := Auths{Auth:reg}
		marshal, err := json.Marshal(a)
		if err!=nil{
			fmt.Println(err)
			return nil, err
			}
			secret.StringData[".dockerconfigjson"] = string(marshal)
			return &secret, nil
	}
}