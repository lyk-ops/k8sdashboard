package pod

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"kubeimook/model/base"
	pod_req "kubeimook/model/pod/request"
	"strconv"
	"strings"
)

const (
	probe_http              = "http"
	probe_tcp               = "tcp"
	probe_exec              = "exec"
	volume_emptyDir         = "emptyDir"
	volume_configMap        = "configMap"
	volume_secret           = "secret"
	volume_hostPath         = "hostPath"
	volume_downword         = "downwardAPI"
	volume_pvc              = "pvc"
	scheduling_nodename     = "=nodeName"
	scheduling_nodeselector = "=nodeSelector"
	scheduling_nodeaffinity = "=nodeAffinity"
	scheduling_nodeany      = "nodeAny"
	ref_type_configMap      = "configMap"
	ref_type_secret         = "secret"
)

type Req2K8sConvert struct {
}

func getNodeK8sScheduling(podReq pod_req.Pod) (affinity *corev1.Affinity, nodeSelector map[string]string, nodeName string) {
	nodeScheduling := podReq.NodeScheduling
	switch nodeScheduling.Type {
	case scheduling_nodename:
		nodeName = nodeScheduling.NodeName
		return
	case scheduling_nodeselector:
		nodeSelectorMap := make(map[string]string)
		for _, item := range nodeScheduling.NodeSelector {
			nodeSelector[item.Key] = item.Value
		}
		nodeSelector = nodeSelectorMap
		return
	case scheduling_nodeaffinity:
		nodeSelectorTermExpressions := nodeScheduling.NodeAffinity
		matchExpressions := make([]corev1.NodeSelectorRequirement, 0)
		for _, expression := range nodeSelectorTermExpressions {
			matchExpressions = append(matchExpressions, corev1.NodeSelectorRequirement{
				Key:      expression.Key,
				Operator: expression.Operator,
				Values:   strings.Split(expression.Values, ","),
			})

		}
		affinity = &corev1.Affinity{
			NodeAffinity: &corev1.NodeAffinity{
				RequiredDuringSchedulingIgnoredDuringExecution: &corev1.NodeSelector{
					NodeSelectorTerms: []corev1.NodeSelectorTerm{
						{
							MatchExpressions: matchExpressions,
						},
					},
				},
			},
		}
	case scheduling_nodeany:
		// do nothing
	default:
		// do nothing

	}
	return

}

// 将pod的 请求格式 的数据转换为k8s结构的数据
func (pc *Req2K8sConvert) PodReq2K8s(podReq pod_req.Pod) *corev1.Pod {
	affinity, selector, name := getNodeK8sScheduling(podReq)
	return &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      podReq.Base.Name,
			Labels:    pc.getK8sLabels(podReq.Base.Labels),
			Namespace: podReq.Base.Namespace,
		},
		Spec: corev1.PodSpec{
			NodeName:       name,
			NodeSelector:   selector,
			Affinity:       affinity,
			Tolerations:    podReq.Tolerations,
			InitContainers: pc.GetK8sContainers(podReq.InitContainers),
			Containers:     pc.GetK8sContainers(podReq.Containers),
			Volumes:        pc.getK8sVolumes(podReq.Volumes),
			DNSConfig: &corev1.PodDNSConfig{
				Nameservers: podReq.Networking.DnsConfig.NameServers,
			},
			DNSPolicy:     corev1.DNSPolicy(podReq.Networking.DnsPolicy),
			HostAliases:   pc.getK8sHostAliases(podReq.Networking.HostAliases),
			Hostname:      podReq.Networking.HostName,
			RestartPolicy: corev1.RestartPolicy(podReq.Base.RestartPolicy),
		},
	}
}
func (pc *Req2K8sConvert) getK8sHostAliases(podReqHostAliases []base.ListMapItem) []corev1.HostAlias {
	podK8sHostAliases := make([]corev1.HostAlias, 0)
	for _, item := range podReqHostAliases {
		podK8sHostAliases = append(podK8sHostAliases, corev1.HostAlias{
			IP:        item.Key,
			Hostnames: strings.Split(item.Value, ","),
		})
	}
	return podK8sHostAliases
}
func (pc *Req2K8sConvert) getK8sVolumes(podReqVolumes []pod_req.Volume) []corev1.Volume {
	podK8sVolumes := make([]corev1.Volume, 0)
	for _, volume := range podReqVolumes {
		//if volume.Type != volume_emptyDir {
		//	continue
		//}
		source := corev1.VolumeSource{}
		switch volume.Type {
		case volume_emptyDir:
			source = corev1.VolumeSource{
				EmptyDir: &corev1.EmptyDirVolumeSource{},
			}
		case volume_hostPath:
			source = corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: volume.HostPathVolume.Path,
					Type: &volume.HostPathVolume.Type,
				},
			}
		case volume_configMap:
			source = corev1.VolumeSource{
				ConfigMap: &corev1.ConfigMapVolumeSource{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: volume.ConfigMapRefVolume.Name,
					},
					Optional: &volume.ConfigMapRefVolume.Optional,
				},
			}
		case volume_secret:
			source = corev1.VolumeSource{
				Secret: &corev1.SecretVolumeSource{
					SecretName: volume.SecretRefVolume.Name,
				},
			}
		case volume_downword:
			items := make([]corev1.DownwardAPIVolumeFile, 0)
			for _, item := range volume.DownWardAPIVolume.Items {
				items = append(items, corev1.DownwardAPIVolumeFile{
					// 容器内的文件访问路径
					Path: item.Path,
					FieldRef: &corev1.ObjectFieldSelector{
						FieldPath: item.FieldRefPath,
					},
				})

				source = corev1.VolumeSource{
					DownwardAPI: &corev1.DownwardAPIVolumeSource{
						Items: []corev1.DownwardAPIVolumeFile{},
					},
				}
			}
		case volume_pvc:
			source = corev1.VolumeSource{
				PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
					ClaimName: volume.PvcVolume.ClaimName,
				},
			}

		default:
			continue
		}

		source = corev1.VolumeSource{
			EmptyDir: &corev1.EmptyDirVolumeSource{},
		}
		podK8sVolumes = append(podK8sVolumes, corev1.Volume{
			VolumeSource: source,
			Name:         volume.Name,
		})
	}
	return podK8sVolumes

}
func (pc *Req2K8sConvert) GetK8sContainers(podReqContainers []pod_req.Container) []corev1.Container {
	podK8sContainers := make([]corev1.Container, 0)
	for _, container := range podReqContainers {
		podK8sContainers = append(podK8sContainers, pc.GetK8sContainer(container))
	}
	return podK8sContainers
}
func (pc *Req2K8sConvert) GetK8sContainer(podReqContainer pod_req.Container) corev1.Container {
	return corev1.Container{
		Name:            podReqContainer.Name,
		Image:           podReqContainer.Image,
		ImagePullPolicy: corev1.PullPolicy(podReqContainer.ImagePullPolicy),
		TTY:             podReqContainer.Tty,
		WorkingDir:      podReqContainer.WorkingDir,
		Command:         podReqContainer.Command,
		Args:            podReqContainer.Args,
		Ports:           pc.getK8sPorts(podReqContainer.Ports),
		Env:             pc.getK8sEnv(podReqContainer.Envs),
		EnvFrom:         pc.getK8sEnvsFrom(podReqContainer.EnvsFrom),
		SecurityContext: &corev1.SecurityContext{
			Privileged: &podReqContainer.Privileged,
		},
		VolumeMounts:   pc.getK8sVolumentMount(podReqContainer.VolumeMounts),
		StartupProbe:   pc.getK8sContainerProbe(podReqContainer.StartupProbe),
		LivenessProbe:  pc.getK8sContainerProbe(podReqContainer.LivenessProbe),
		ReadinessProbe: pc.getK8sContainerProbe(podReqContainer.ReadinessProbe),
		Resources:      pc.getK8sResources(podReqContainer.Resources),
	}
}
func (pc *Req2K8sConvert) getK8sPorts(podReqPorts []pod_req.ContainerPort) []corev1.ContainerPort {
	podK8sContainerPorts := make([]corev1.ContainerPort, 0)
	for _, port := range podReqPorts {
		podK8sContainerPorts = append(podK8sContainerPorts, corev1.ContainerPort{
			Name:          port.Name,
			ContainerPort: port.ContainerPort,
			HostPort:      port.HostPort,
		})
	}
	return podK8sContainerPorts
}
func (pc *Req2K8sConvert) getK8sResources(podReqResources pod_req.Resources) corev1.ResourceRequirements {
	var k8sResources corev1.ResourceRequirements
	if !podReqResources.Enable {
		return k8sResources
	}
	k8sResources.Requests = corev1.ResourceList{
		corev1.ResourceCPU:    resource.MustParse(strconv.Itoa(int(podReqResources.CpuRequest)) + "m"),
		corev1.ResourceMemory: resource.MustParse(strconv.Itoa(int(podReqResources.MemRequest)) + "Mi"),
	}
	k8sResources.Limits = corev1.ResourceList{
		corev1.ResourceCPU:    resource.MustParse(strconv.Itoa(int(podReqResources.CpuLimit)) + "m"),
		corev1.ResourceMemory: resource.MustParse(strconv.Itoa(int(podReqResources.MemLimit)) + "Mi"),
	}
	return corev1.ResourceRequirements{}
}
func (pc *Req2K8sConvert) getK8sContainerProbe(podReqProbe pod_req.ContainerProbe) *corev1.Probe {
	if !podReqProbe.Enable {
		return nil
	}
	var k8sProbe corev1.Probe
	switch podReqProbe.Type {
	case probe_http:
		httpGet := podReqProbe.HttpGet
		k8sHttpHeaders := make([]corev1.HTTPHeader, 0)
		for _, header := range httpGet.HttpHeaders {
			k8sHttpHeaders = append(k8sHttpHeaders, corev1.HTTPHeader{
				Name:  header.Key,
				Value: header.Value,
			})
		}
		k8sProbe.HTTPGet = &corev1.HTTPGetAction{
			Scheme:      corev1.URIScheme(httpGet.Scheme),
			Host:        httpGet.Host,
			Path:        httpGet.Path,
			Port:        intstr.FromInt(int(httpGet.Port)),
			HTTPHeaders: k8sHttpHeaders,
		}
	case probe_exec:
		exec := podReqProbe.Exec
		k8sProbe.Exec = &corev1.ExecAction{
			Command: exec.Command,
		}
	case probe_tcp:
		tcpSocket := podReqProbe.TcpSocket
		k8sProbe.TCPSocket = &corev1.TCPSocketAction{
			Host: tcpSocket.Host,
			Port: intstr.FromInt(int(tcpSocket.Port)),
		}
	}
	return &k8sProbe
}
func (pc *Req2K8sConvert) getK8sVolumentMount(podReqMounts []pod_req.VolumeMounts) []corev1.VolumeMount {
	podK8sVolumeMounts := make([]corev1.VolumeMount, 0)
	for _, mount := range podReqMounts {
		podK8sVolumeMounts = append(podK8sVolumeMounts, corev1.VolumeMount{
			Name:      mount.MountName,
			MountPath: mount.MountPath,
			ReadOnly:  mount.ReadOnly,
		})
	}
	return podK8sVolumeMounts
}
func (pc *Req2K8sConvert) getK8sEnv(podReqEnvs []pod_req.EnvVar) []corev1.EnvVar {
	podK8sEnvs := make([]corev1.EnvVar, 0)
	for _, env := range podReqEnvs {
		envVar := corev1.EnvVar{
			Name: env.Name,
		}
		switch env.Type {
		case ref_type_configMap:
			envVar.ValueFrom = &corev1.EnvVarSource{
				ConfigMapKeyRef: &corev1.ConfigMapKeySelector{
					Key: env.Value,
					LocalObjectReference: corev1.LocalObjectReference{
						Name: env.RefName,
					},
				},
			}
		case ref_type_secret:
			envVar.ValueFrom = &corev1.EnvVarSource{
				SecretKeyRef: &corev1.SecretKeySelector{
					Key: env.Value,
					LocalObjectReference: corev1.LocalObjectReference{
						Name: env.RefName,
					},
				},
			}
		default:
			envVar.Value = env.Value

		}
		podK8sEnvs = append(podK8sEnvs, envVar)
	}
	return podK8sEnvs
}

// 获取k8s的labels
func (*Req2K8sConvert) getK8sLabels(podReqLabels []base.ListMapItem) map[string]string {
	podK8sLabels := make(map[string]string)
	for _, label := range podReqLabels {
		podK8sLabels[label.Key] = label.Value
	}
	return podK8sLabels
}

func (pc *Req2K8sConvert) getK8sEnvsFrom(podReqEnvsFrom []pod_req.EnvVarFromResource) []corev1.EnvFromSource {
	podK8sEnvsFrom := make([]corev1.EnvFromSource, 0)
	for _, fromResource := range podReqEnvsFrom {
		//前缀通用
		envFromResource := corev1.EnvFromSource{
			Prefix: fromResource.Prefix,
		}
		switch fromResource.RefType {
		case ref_type_configMap:
			envFromResource.ConfigMapRef = &corev1.ConfigMapEnvSource{
				LocalObjectReference: corev1.LocalObjectReference{
					Name: fromResource.Name,
				},
			}
			podK8sEnvsFrom = append(podK8sEnvsFrom, envFromResource)
		case ref_type_secret:
			envFromResource.SecretRef = &corev1.SecretEnvSource{
				LocalObjectReference: corev1.LocalObjectReference{
					Name: fromResource.Name,
				},
			}
			podK8sEnvsFrom = append(podK8sEnvsFrom, envFromResource)
		}

	}
	return podK8sEnvsFrom
}
