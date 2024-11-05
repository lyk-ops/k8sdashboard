package response

type Namespace struct {
	Name              string `json:"name"`
	CreationTimestamp string `json:"creationTimestamp"`
	Status            string `json:"status"`
}
