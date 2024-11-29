package node

import (
	corev1 "k8s.io/api/core/v1"
	"kubeimook/model/base"
	node_res "kubeimook/model/node/response"
)

type NodeK8s2Res struct {
}

func getNodeStatus(nodeCondition []corev1.NodeCondition) string {
	nodeStatus := "NotReady"
	for _, condition := range nodeCondition {
		if condition.Type == "Ready" && condition.Status == "True" {
			nodeStatus = "Ready"
			break
		}
	}
	return nodeStatus
}
func getNodeIp(addresses []corev1.NodeAddress, addressType corev1.NodeAddressType) string {
	for _, address := range addresses {
		if address.Type == addressType && address.Address != "" {
			return address.Address
		}
	}
	return "<none>"
}

// GetNodeResItem 方法将 Kubernetes 的 Node 对象转换为自定义的 node_res.Node 类型对象。
//
// 参数:
//
//	nodeK8s - corev1.Node 类型的参数，表示要转换的 Kubernetes 节点对象。
//
// 返回值:
//
//	node_res.Node - 转换后的自定义节点对象，包含了节点的详细信息。
//
// 说明:
// 该方法首先从输入的 Kubernetes 节点对象中提取相关信息，然后创建一个自定义的 node_res.Node 类型对象，
// 并填充该对象的各个字段，最后返回这个自定义对象。
func (*NodeK8s2Res) GetNodeResItem(nodeK8s corev1.Node) node_res.Node {
	nodeInfo := nodeK8s.Status.NodeInfo
	return node_res.Node{
		Name:             nodeK8s.GetName(),
		Status:           getNodeStatus(nodeK8s.Status.Conditions),
		Age:              nodeK8s.CreationTimestamp.Unix(),
		InternalIp:       getNodeIp(nodeK8s.Status.Addresses, corev1.NodeInternalIP),
		ExternalIp:       getNodeIp(nodeK8s.Status.Addresses, corev1.NodeExternalIP),
		OsImage:          nodeInfo.OSImage,
		Version:          nodeInfo.KubeletVersion,
		KernelVersion:    nodeInfo.KernelVersion,
		ContainerRuntime: nodeInfo.ContainerRuntimeVersion,
	}
}
func mapToList(m map[string]string) []base.ListMapItem {
	res := make([]base.ListMapItem, 0)
	for k, v := range m {
		res = append(res, base.ListMapItem{
			Key:   k,
			Value: v,
		})
	}
	return res
}

// GetNodeDetail 方法用于获取 Kubernetes 节点的详细信息，并将其转换为自定义的 node_res.Node 类型对象。
//
// 参数:
//
//	nodeK8s - corev1.Node 类型的参数，表示要获取详细信息的 Kubernetes 节点对象。
//
// 返回值:
//
//	node_res.Node - 转换后的自定义节点对象，包含了节点的详细信息。
//
// 说明:
// 该方法首先调用 GetNodeResItem 方法获取节点的基本信息，并填充到 node_res.Node 类型的对象中。
// 然后，它计算并设置节点的 Labels 和 Taints 信息，最终返回完整的自定义节点对象。
//
// - Labels: 将 Kubernetes 节点的 Labels 转换为 map[string]string 类型的列表。
// - Taints: 直接使用 Kubernetes 节点的 Taints 信息。
func (this *NodeK8s2Res) GetNodeDetail(nodeK8s corev1.Node) node_res.Node {
	nodeRes := this.GetNodeResItem(nodeK8s)
	//计算label和taint
	nodeRes.Labels = mapToList(nodeK8s.Labels)
	nodeRes.Taints = nodeK8s.Spec.Taints
	return nodeRes
}
