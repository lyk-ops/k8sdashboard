package node

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeimook/global"
	node_res "kubeimook/model/node/response"
	"strings"
)

type NodeService struct {
}

/*
NAME              STATUS                     ROLES     AGE      VERSION   INTERNAL-IP       EXTERNAL-IP   OS-IMAGE                KERNEL-VERSION                CONTAINER-RUNTIME   LABELS
192.168.150.113   Ready,SchedulingDisabled   master    2y347d   v1.20.5   192.168.150.113   <none>        CentOS Linux 7 (Core)   3.10.0-1127.19.1.el7.x86_64   docker://20.10.5    beta.kubernetes.io/arch=amd64,beta.kubernetes.io/os=linux,kubernetes.io/arch=amd64,kubernetes.io/hostname=192.168.150.113,kubernetes.io/os=linux,kubernetes.io/role=master
192.168.150.114   Ready,SchedulingDisabled   master    2y347d   v1.20.5   192.168.150.114   <none>        CentOS Linux 7 (Core)   3.10.0-1127.19.1.el7.x86_64   docker://20.10.5    beta.kubernetes.io/arch=amd64,beta.kubernetes.io/os=linux,kubernetes.io/arch=amd64,kubernetes.io/hostname=192.168.150.114,kubernetes.io/os=linux,kubernetes.io/role=master
192.168.150.116   Ready,SchedulingDisabled   master    2y347d   v1.20.5   192.168.150.116   <none>        CentOS Linux 7 (Core)   3.10.0-1127.19.1.el7.x86_64   docker://20.10.5    beta.kubernetes.io/arch=amd64,beta.kubernetes.io/os=linux,kube-ovn/role=master,kubernetes.io/arch=amd64,kubernetes.io/hostname=192.168.150.116,kubernetes.io/os=linux,kubernetes.io/role=master,node-role.kubernetes.io/master=master
192.168.150.122   Ready,SchedulingDisabled   ingress   2y347d   v1.20.5   192.168.150.122   <none>        CentOS Linux 7 (Core)   3.10.0-1127.19.1.el7.x86_64   docker://20.10.5    beta.kubernetes.io/arch=amd64,beta.kubernetes.io/os=linux,kubernetes.io/arch=amd64,kubernetes.io/hostname=192.168.150.122,kubernetes.io/os=linux,kubernetes.io/role=ingress
192.168.150.123   Ready                      ingress   2y347d   v1.20.5   192.168.150.123   <none>        CentOS Linux 7 (Core)   3.10.0-1127.19.1.el7.x86_64   docker://20.10.5    beta.kubernetes.io/arch=amd64,beta.kubernetes.io/os=linux,kubernetes.io/arch=amd64,kubernetes.io/hostname=192.168.150.123,kubernetes.io/os=linux,kubernetes.io/role=ingress
192.168.150.124   Ready                      web       2y347d   v1.20.5   192.168.150.124   <none>        CentOS Linux 7 (Core)   3.10.0-1127.19.1.el7.x86_64   docker://20.10.5    beta.kubernetes.io/arch=amd64,beta.kubernetes.io/os=linux,kubernetes.io/arch=amd64,kubernetes.io/hostname=192.168.150.124,kubernetes.io/jenkins=jenkins,kubernetes.io/os=linux,kubernetes.io/role=web
192.168.150.125   Ready                      web       2y347d   v1.20.5   192.168.150.125   <none>        CentOS Linux 7 (Core)   3.10.0-1127.19.1.el7.x86_64   docker://20.10.5    beta.kubernetes.io/arch=amd64,beta.kubernetes.io/os=linux,kubernetes.io/arch=amd64,kubernetes.io/hostname=192.168.150.125,kubernetes.io/os=linux,kubernetes.io/role=web
192.168.150.130   Ready                      web       2y197d   v1.20.5   192.168.150.130   <none>        CentOS Linux 7 (Core)   3.10.0-1127.19.1.el7.x86_64   docker://20.10.5    beta.kubernetes.io/arch=amd64,beta.kubernetes.io/os=linux,kubernetes.io/arch=amd64,kubernetes.io/hostname=192.168.150.130,kubernetes.io/os=linux,kubernetes.io/role=web
*/
//func (*NodeService) GetNodeList(keyword string) ([]node_res.Node, error) {
//	ctx := context.TODO()
//	list, err := global.KubeConfigSet.CoreV1().Nodes().List(ctx, metav1.ListOptions{
//		FieldSelector: "metadata.name=" + keyword,
//	})
//	if err != nil {
//		return nil, err
//	}
//	//fmt.Printf("list:%v\n", list)
//	nodeResList := make([]node_res.Node, 0)
//	for _, item := range list.Items {
//		if strings.Contains(item.Name, keyword) {
//			nodeRes := nodeConvert.GetNodeResItem(item)
//			nodeResList = append(nodeResList, nodeRes)
//		}
//
//	}
//	return nodeResList, nil
//}
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
