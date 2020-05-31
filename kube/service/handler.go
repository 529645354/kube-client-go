package service

import (
	"fmt"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func (s *Service) Selectors() map[string]string {
	selector := make(map[string]string, 0)
	for _, i := range s.Selector {
		selector[i.Key] = i.Value
	}
	return selector
}

func (s *Service) PortList() []apiv1.ServicePort {
	sp := make([]apiv1.ServicePort, 0)
	if s.Type == "nodeport" {
		for _, i := range s.NodePort.Port {
			p := apiv1.ServicePort{
				Protocol: apiv1.ProtocolTCP,
				Port:     i.Port,
				Name:     fmt.Sprintf("%dto%d", i.Port, i.TargetPort),
				TargetPort: intstr.IntOrString{
					IntVal: i.TargetPort,
				},
			}
			if i.Protocol == "TCP" {
				p.Protocol = apiv1.ProtocolTCP
			} else if i.Protocol == "SCTP" {
				p.Protocol = apiv1.ProtocolSCTP
			} else {
				p.Protocol = apiv1.ProtocolUDP
			}
			sp = append(sp, p)
		}
	} else {
		for _, i := range s.ClusterIP.Port {
			p := apiv1.ServicePort{
				Protocol: apiv1.ProtocolTCP,
				Port:     i.Port,
				TargetPort: intstr.IntOrString{
					IntVal: i.TargetPort,
				},
			}
			if i.Protocol == "TCP" {
				p.Protocol = apiv1.ProtocolTCP
			} else if i.Protocol == "SCTP" {
				p.Protocol = apiv1.ProtocolSCTP
			} else {
				p.Protocol = apiv1.ProtocolUDP
			}
			sp = append(sp, p)
		}
	}

	return sp
}

func (s *Service) CreateService() *apiv1.Service {
	svc := apiv1.Service{
		ObjectMeta: v1.ObjectMeta{
			Name: s.Name,
		},
		Spec: apiv1.ServiceSpec{},
	}
	if s.Type == "externalname" {
		svc.Spec = apiv1.ServiceSpec{
			Type:         apiv1.ServiceTypeExternalName,
			ExternalName: s.ExternalNameDNS,
		}
		return &svc
	} else if s.Type == "nodeport" {
		svc.Spec = apiv1.ServiceSpec{
			Ports:    s.PortList(),
			Selector: s.Selectors(),
			Type:     apiv1.ServiceTypeNodePort,
		}
		return &svc
	} else {
		sp := apiv1.ServiceSpec{
			Ports:    s.PortList(),
			Selector: s.Selectors(),
			Type:     apiv1.ServiceTypeClusterIP,
		}
		if s.ClusterIP.Headless == true {
			sp.ClusterIP = "None"
		}
		svc.Spec = sp
		return &svc
	}
}
