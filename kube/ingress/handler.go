package ingress

import (
	"k8s.io/api/networking/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func (i *Ingress) HostPath(httpls []Rule) []v1beta1.HTTPIngressPath  {
	ingressPaths := make([]v1beta1.HTTPIngressPath, 0)
	for _, i := range httpls {
		p := v1beta1.HTTPIngressPath{
			Path:     i.Path,
			Backend:  v1beta1.IngressBackend{
				ServiceName: i.BackendService,
				ServicePort: intstr.IntOrString{
					IntVal: i.BackendPort,
				},
			},
		}
		ingressPaths = append(ingressPaths, p)
	}
	return ingressPaths
}

func (i *Ingress) CreateRule() []v1beta1.IngressRule {
	ruls := make([]v1beta1.IngressRule,0)
	for _,h := range i.Hosts {
		r := v1beta1.IngressRule{
			Host:             h.Hostname,
			IngressRuleValue: v1beta1.IngressRuleValue{
				HTTP: &v1beta1.HTTPIngressRuleValue{
					Paths: i.HostPath(h.HttpLs),
				},
			},
		}
		ruls = append(ruls, r)
	}
	return ruls
}

func (i *Ingress) CreateIngress() *v1beta1.Ingress  {
	ingress := v1beta1.Ingress{
		ObjectMeta: v1.ObjectMeta{
			Name: i.Name,
		},
		Spec:       v1beta1.IngressSpec{
			Rules: i.CreateRule(),
		},
	}
	return &ingress
}
