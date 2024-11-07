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
