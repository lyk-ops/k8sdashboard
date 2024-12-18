package request

import (
	corev1 "k8s.io/api/core/v1"
	"kubeimook/model/base"
)

type Base struct {
	Name          string             `json:"name"`
	Labels        []base.ListMapItem `json:"label"`
	Namespace     string             `json:"namespace"`
	RestartPolicy string             `json:"restartPolicy"`

	//
}
type ConfigMapRefVolume struct {
	Name     string `json:"name"`
	Optional bool   `json:"optional"`
}
type SecretRefVolume struct {
	Name     string `json:"name"`
	Optional bool   `json:"optional"`
}
type HostPathVolume struct {
	Path string              `json:"path"` // 宿主机路径
	Type corev1.HostPathType `json:"type"`
}
type DownWardAPIVolumeItem struct {
	Path         string `json:"path"` // 容器路径
	FieldRefPath string `json:"fieldRefPath"`
}
type DownWardAPIVolume struct {
	Items []DownWardAPIVolumeItem `json:"items"`
}
type PvcVolume struct {
	ClaimName string `json:"claimName"`
}
type Volume struct {
	Name string `json:"name"`
	// emptyDir\hostPath\pvc\configMap\secret\downward
	Type               string             `json:"type"`
	ConfigMapRefVolume ConfigMapRefVolume `json:"configMapRefVolume"`
	SecretRefVolume    SecretRefVolume    `json:"secretRefVolume"`
	HostPathVolume     HostPathVolume     `json:"hostPathVolume"`
	DownWardAPIVolume  DownWardAPIVolume  `json:"downwardApiVolume"`
	PvcVolume          PvcVolume          `json:"pvcVolume"`
}
type DnsConfig struct {
	NameServers []string `json:"nameservers"`
}
type Networking struct {
	HostNetwork bool               `json:"hostNetwork"`
	HostName    string             `json:"hostname"`
	DnsPolicy   string             `json:"dnsPolicy"`
	DnsConfig   DnsConfig          `json:"dnsConfig"`
	HostAliases []base.ListMapItem `json:"hostAliases"`
}
type Resources struct {
	Enable     bool  `json:"enable"` // 是否开启资源限制
	MemRequest int32 `json:"MemRequest"`
	MemLimit   int32 `json:"MemLimit"`
	CpuRequest int32 `json:"CpuRequest"`
	CpuLimit   int32 `json:"CpuLimit"`
}
type VolumeMounts struct {
	MountName string `json:"name"`
	MountPath string `json:"mountPath"`
	ReadOnly  bool   `json:"readOnly"`
}
type ProbeCommand struct {
	Command []string `json:"command"`
}
type ProbeHttpGet struct {
	Scheme      string             `json:"scheme"`
	Host        string             `json:"host"` // 请求host,如果为空，则使用pod的IP地址
	Path        string             `json:"path"`
	Port        int32              `json:"port"`
	HttpHeaders []base.ListMapItem `json:"httpHeaders"`
}
type ProbeTcpSocket struct {
	Host string `json:"host"`
	Port int32  `json:"port"`
}
type ProbeTime struct {
	FailureThreshold    int32 `json:"failureThreshold"`
	InitialDelaySeconds int32 `json:"initialDelaySeconds"`
	PeriodSeconds       int32 `json:"periodSeconds"`
	SuccessThreshold    int32 `json:"successThreshold"`
	TimeoutSeconds      int32 `json:"timeoutSeconds"`
}
type ContainerProbe struct {
	Enable    bool           `json:"enable"` // 是否开启探针
	Type      string         `json:"type"`   // 探针类型
	Exec      ProbeCommand   `json:"exec"`
	HttpGet   ProbeHttpGet   `json:"httpGet"`
	TcpSocket ProbeTcpSocket `json:"tcpSocket"`
	ProbeTime
}
type ContainerPort struct {
	Name          string `json:"name"`
	ContainerPort int32  `json:"containerPort"`
	HostPort      int32  `json:"hostPort"`
}
type EnvVar struct {
	Name    string `json:"name"`
	RefName string `json:"refName"`
	Type    string `json:"type"` // env type
	// configmap secret default(key:value)
	Value string `json:"value"`
}
type EnvVarFromResource struct {
	Name    string `json:"name"`    // 资源名称
	RefType string `json:"refType"` // 资源类型 configmap secret
	Prefix  string `json:"prefix"`  // 资源前缀,用于表示环境变量前缀
}
type Container struct {
	Name            string               `json:"name"`
	Image           string               `json:"image"`
	ImagePullPolicy string               `json:"imagePullPolicy"`
	Tty             bool                 `json:"tty"`        //是否开启tty
	WorkingDir      string               `json:"workingDir"` //工作目录
	Command         []string             `json:"command"`    // 命令
	Args            []string             `json:"args"`       // 参数
	Envs            []EnvVar             `json:"envs"`       // 环境变量
	EnvsFrom        []EnvVarFromResource `json:"envsFrom"`
	Privileged      bool                 `json:"privileged"`     //是否特权模式
	Resources       Resources            `json:"resources"`      //资源限制
	VolumeMounts    []VolumeMounts       `json:"volumeMounts"`   //挂载卷
	StartupProbe    ContainerProbe       `json:"startupProbe"`   //启动探针
	LivenessProbe   ContainerProbe       `json:"livenessProbe"`  //存活探针
	ReadinessProbe  ContainerProbe       `json:"readinessProbe"` //就绪探针
	Ports           []ContainerPort      `json:"ports"`          //端口映射
}

type NodeSelectTermExpressions struct {
	Key      string                      `json:"key"`
	Operator corev1.NodeSelectorOperator `json:"operator"`
	Values   string
}
type NodeScheduling struct {
	// nodeName nodeSelector nodeAffinity
	Type         string                      `json:"type"`         //调度类型
	NodeName     string                      `json:"nodeName"`     //节点名称
	NodeSelector []base.ListMapItem          `json:"nodeSelector"` //节点选择器
	NodeAffinity []NodeSelectTermExpressions `json:"nodeAffinity"` //节点亲和性
}
type Pod struct {
	Base           Base                `json:"base"`
	Tolerations    []corev1.Toleration `json:"tolerations"`
	NodeScheduling NodeScheduling      `json:"nodeScheduling"`
	Volumes        []Volume            `json:"volumes"`
	Networking     Networking          `json:"networking"`
	InitContainers []Container         `json:"initContainers"`
	Containers     []Container         `json:"containers"`
	//
}
