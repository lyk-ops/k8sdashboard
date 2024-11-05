package pod

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"kubeimook/model/base"
	pod_req "kubeimook/model/pod/request"
	pod_res "kubeimook/model/pod/response"
	"strings"
)

const volume_type_emptydir = "emptyDir"

type K8s2ReqConvert struct {
	volumeMap map[string]string
}

func (*K8s2ReqConvert) PodK8s2ItemRes(pod corev1.Pod) pod_res.PodListItem {
	var totalC, redayC, restartC int32
	for _, containerStatus := range pod.Status.ContainerStatuses {
		if containerStatus.Ready {
			redayC++
		}
		restartC += containerStatus.RestartCount
		totalC++
	}
	var podStatus string
	if pod.Status.Phase != "Running" {
		podStatus = "Error"
	} else {
		podStatus = "Running"
	}
	return pod_res.PodListItem{
		Name:     pod.Name,
		Ready:    fmt.Sprintf("%d/%d", redayC, totalC),
		Status:   podStatus,
		Restarts: restartC,
		Age:      pod.CreationTimestamp.Unix(),
		IP:       pod.Status.PodIP, //运行起来时才有状态信息，所以从status获取
		Node:     pod.Spec.NodeName,
	}
}

func (this *K8s2ReqConvert) PodK8s2Req(podK8s corev1.Pod) pod_req.Pod {

	return pod_req.Pod{
		Base:           getReqBase(podK8s),
		Networking:     getReqNetworking(podK8s),
		Volumes:        this.getReqVolumes(podK8s.Spec.Volumes),
		Containers:     this.getReqContainers(podK8s.Spec.Containers),
		InitContainers: this.getReqContainers(podK8s.Spec.InitContainers),
	}
}
func (this *K8s2ReqConvert) getReqContainers(containersK8s []corev1.Container) []pod_req.Container {
	podReqContainers := make([]pod_req.Container, 0)
	for _, container := range containersK8s {
		// container转换
		reqContainer := this.getReqContanier(container)
		podReqContainers = append(podReqContainers, reqContainer)
	}
	return podReqContainers
}
func (this *K8s2ReqConvert) getReqContanier(container corev1.Container) pod_req.Container {
	return pod_req.Container{
		Name:            container.Name,
		Image:           container.Image,
		ImagePullPolicy: string(container.ImagePullPolicy),
		Tty:             container.TTY,
		WorkingDir:      container.WorkingDir,
		Command:         container.Command,

		//
		Ports:          getReqContainerPorts(container.Ports),
		Args:           container.Args,
		Env:            getReqContainersEnvs(container.Env),
		Privileged:     getReqContainerPrivileged(container.SecurityContext),
		Resources:      getReqContainerResources(container.Resources),
		VolumeMounts:   this.getReqContainerVolumeMount(container.VolumeMounts),
		StartupProbe:   getReqContainerProbe(container.StartupProbe),
		LivenessProbe:  getReqContainerProbe(container.LivenessProbe),
		ReadinessProbe: getReqContainerProbe(container.ReadinessProbe),
	}
}
func getReqContainerResources(requirements corev1.ResourceRequirements) pod_req.Resources {
	reqResources := pod_req.Resources{
		Enable: false,
	}
	requests := requirements.Requests
	limits := requirements.Limits
	if requests != nil {
		reqResources.Enable = true
		reqResources.CpuRequest = int32(requests.Cpu().MilliValue()) // m
		//MiB
		reqResources.MemRequest = int32(requests.Memory().Value() / (1024 * 1024)) //Bytes
	}
	if limits != nil {
		reqResources.Enable = true
		reqResources.CpuLimit = int32(limits.Cpu().MilliValue())
		reqResources.MemLimit = int32(limits.Memory().Value() / (1024 * 1024))
	}
	return reqResources
}
func getReqContainerPrivileged(ctx *corev1.SecurityContext) (privileged bool) {
	if ctx != nil {
		privileged = *ctx.Privileged
	}
	return
}
func getReqContainerProbe(probeK8s *corev1.Probe) pod_req.ContainerProbe {
	containerProbe := pod_req.ContainerProbe{
		Enable: false,
	}
	//先判断探针是否为空
	if probeK8s != nil {
		//再判断探针是什么类型
		containerProbe.Enable = true
		if probeK8s.Exec != nil {
			containerProbe.Type = probe_exec
			containerProbe.Exec.Command = probeK8s.Exec.Command
		} else if probeK8s.HTTPGet != nil {
			containerProbe.Type = probe_http
			httpGet := probeK8s.HTTPGet
			headersReq := make([]base.ListMapItem, 0)
			for _, headerK8s := range httpGet.HTTPHeaders {
				headersReq = append(headersReq, base.ListMapItem{
					Key:   headerK8s.Name,
					Value: headerK8s.Value,
				})
			}
			containerProbe.HttpGet = pod_req.ProbeHttpGet{
				Host:        httpGet.Host,
				Port:        httpGet.Port.IntVal,
				Scheme:      string(httpGet.Scheme),
				Path:        httpGet.Path,
				HttpHeaders: headersReq,
			}
		} else if probeK8s.TCPSocket != nil {
			containerProbe.Type = probe_tcp
			containerProbe.TcpSocket = pod_req.ProbeTcpSocket{
				Host: probeK8s.TCPSocket.Host,
				Port: probeK8s.TCPSocket.Port.IntVal,
			}
		} else {
			containerProbe.Type = probe_http
			return containerProbe
		}
		containerProbe.InitialDelaySeconds = probeK8s.InitialDelaySeconds
		containerProbe.PeriodSeconds = probeK8s.PeriodSeconds
		containerProbe.TimeoutSeconds = probeK8s.TimeoutSeconds
		containerProbe.FailureThreshold = probeK8s.FailureThreshold
		containerProbe.SuccessThreshold = probeK8s.SuccessThreshold
	}
	return containerProbe
}
func (this *K8s2ReqConvert) getReqContainerVolumeMount(volumeMountK8s []corev1.VolumeMount) []pod_req.VolumeMounts {
	volumesReq := make([]pod_req.VolumeMounts, 0)
	for _, volumeMount := range volumeMountK8s {
		//非emptydir的过滤掉
		_, ok := this.volumeMap[volumeMount.Name]
		if ok {
			volumesReq = append(volumesReq, pod_req.VolumeMounts{
				MountName: volumeMount.Name,
				MountPath: volumeMount.MountPath,
				ReadOnly:  volumeMount.ReadOnly,
			})
		}

	}
	return volumesReq
}
func getReqContainersEnvs(envsK8s []corev1.EnvVar) []base.ListMapItem {
	envsReq := make([]base.ListMapItem, 0)
	for _, env := range envsK8s {
		envsReq = append(envsReq, base.ListMapItem{
			Key:   env.Name,
			Value: env.Value,
		})
	}
	return envsReq

}
func getReqContainerPorts(portsK8s []corev1.ContainerPort) []pod_req.ContainerPort {
	portsReq := make([]pod_req.ContainerPort, 0)
	for _, port := range portsK8s {
		portsReq = append(portsReq, pod_req.ContainerPort{
			Name:          port.Name,
			HostPort:      port.HostPort,
			ContainerPort: port.ContainerPort,
		})
	}
	return portsReq
}
func (this *K8s2ReqConvert) getReqVolumes(volumes []corev1.Volume) []pod_req.Volume {
	volumeReq := make([]pod_req.Volume, 0)
	for _, volume := range volumes {
		if volume.EmptyDir == nil {
			continue
		}
		if this.volumeMap == nil {
			this.volumeMap = make(map[string]string)
		}
		this.volumeMap[volume.Name] = ""
		volumeReq = append(volumeReq, pod_req.Volume{
			Type: volume_type_emptydir,
			Name: volume.Name,
		})
	}
	return volumeReq
}
func getReqHostAlises(hostAlias []corev1.HostAlias) []base.ListMapItem {
	hostAliasReq := make([]base.ListMapItem, 0)
	for _, alias := range hostAlias {
		hostAliasReq = append(hostAliasReq, base.ListMapItem{
			Key:   alias.IP,
			Value: strings.Join(alias.Hostnames, ","),
		})
	}
	return hostAliasReq
}
func getReqDnsConfig(dnsConfigK8s *corev1.PodDNSConfig) pod_req.DnsConfig {
	var dnsConfigReq pod_req.DnsConfig
	if dnsConfigK8s == nil {
		return dnsConfigReq
	} else {
		dnsConfigReq.NameServers = dnsConfigK8s.Nameservers
		return dnsConfigReq
	}
}
func getReqNetworking(pod corev1.Pod) pod_req.Networking {

	return pod_req.Networking{
		HostNetwork: pod.Spec.HostNetwork,
		HostName:    pod.Spec.Hostname,
		DnsPolicy:   string(pod.Spec.DNSPolicy),
		DnsConfig:   getReqDnsConfig(pod.Spec.DNSConfig),
		HostAliases: getReqHostAlises(pod.Spec.HostAliases),
	}

}

func getReqLabels(data map[string]string) []base.ListMapItem {
	labels := make([]base.ListMapItem, 0)
	for k, v := range data {
		labels = append(labels, base.ListMapItem{
			Key:   k,
			Value: v,
		})
	}
	return labels

}
func getReqBase(pod corev1.Pod) pod_req.Base {
	return pod_req.Base{
		Name:          pod.GetName(),
		Namespace:     pod.GetNamespace(),
		Labels:        getReqLabels(pod.Labels),
		RestartPolicy: string(pod.Spec.RestartPolicy),
	}
}
