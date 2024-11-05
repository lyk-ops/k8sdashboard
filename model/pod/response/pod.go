package response

//NAME       名称
//READY      状态 1/2
//STATUS     状态 Running/Pending
//RESTARTS   重启次数
//AGE       时间
//IP        IP地址
//NODE
//NOMINATED NODE
//NODE        节点
//READINESS READY状态
//GATES     状态

type PodListItem struct {
	Name     string `json:"name"`
	Ready    string `json:"ready"`
	Status   string `json:"status"`
	Restarts int32  `json:"restarts"`
	Age      int64  `json:"age"`
	IP       string `json:"IP"`
	Node     string `json:"node"`
}
