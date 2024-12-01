package node

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/json"
	"kubeimook/global"
	node_req "kubeimook/model/node/request"
	node_res "kubeimook/model/node/response"
	"strings"
)

type NodeService struct {
}

// GetNodeList 方法根据给定的关键字获取 Kubernetes 节点列表，并将节点信息转换为自定义的 node_res.Node 类型列表。
//
// 参数:
//
//	keyword - string 类型，表示用于搜索节点的关键字。
//
// 返回值:
//
//	([]node_res.Node, error) - 返回一个包含自定义 node_res.Node 类型对象的切片和一个 error 对象。
//	如果操作成功，将返回包含匹配节点的自定义节点对象切片和 nil；如果操作失败，将返回 nil 和相应的错误。
//
// 说明:
// 该方法首先创建一个上下文对象（虽然在这里使用 context.TODO() 作为占位符）。
// 然后，使用全局的 Kubernetes 客户端配置（global.KubeConfigSet）获取 CoreV1 API 的 Nodes 接口，
// 并调用 List 方法列出所有节点。如果列表操作失败，则立即返回错误。
//
// 接下来，方法创建一个空的 node_res.Node 类型切片用于存储转换后的节点信息。
// 遍历列出的所有节点，如果节点的名称包含给定的关键字，则调用 nodeConvert.GetNodeResItem 方法将该节点转换为自定义的 node_res.Node 类型对象，
// 并将其添加到结果切片中。
//
// 最后，返回包含所有匹配节点的自定义节点对象切片和可能的错误。
func (*NodeService) GetNodeList(keyword string) ([]node_res.Node, error) {
	ctx := context.TODO()
	list, err := global.KubeConfigSet.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	nodeResList := make([]node_res.Node, 0)
	for _, item := range list.Items {
		if strings.Contains(item.Name, keyword) {
			nodeRes := nodeConvert.GetNodeResItem(item)
			nodeResList = append(nodeResList, nodeRes)
		}
	}
	return nodeResList, err
}

func (*NodeService) GetNodeDetail(name string) (*node_res.Node, error) {

	node, err := global.KubeConfigSet.CoreV1().Nodes().Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	detail := nodeConvert.GetNodeDetail(*node)
	return &detail, nil

}
func (*NodeService) UpdateNodeLabel(updateLabels node_req.UpdatedLabel) error {
	labelsMap := make(map[string]string, 0)
	for _, label := range updateLabels.Labels {
		labelsMap[label.Key] = label.Value
	}
	labelsMap["$patch"] = "replace"
	patchData := map[string]any{
		"metadata": map[string]any{
			"labels": labelsMap,
		},
	}
	patchDataBytes, _ := json.Marshal(&patchData)
	global.KubeConfigSet.CoreV1().Nodes().Patch(
		context.TODO(),
		updateLabels.Name,
		types.StrategicMergePatchType,
		patchDataBytes,
		metav1.PatchOptions{},
	)
	return nil
}
func (*NodeService) UpdateNodeTaint(updateTaint node_req.UpdatedTaint) error {
	patchData := map[string]any{
		"spec": map[string]any{
			"taints": updateTaint.Taints,
		},
	}
	patchDataBytes, _ := json.Marshal(&patchData)
	global.KubeConfigSet.CoreV1().Nodes().Patch(
		context.TODO(),
		updateTaint.Name,
		types.StrategicMergePatchType,
		patchDataBytes,
		metav1.PatchOptions{},
	)
	return nil
}
