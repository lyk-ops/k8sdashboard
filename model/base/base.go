package base

type ListMapItem struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (*ListMapItem) ToMap(item []ListMapItem) map[string]string {
	dataMap := make(map[string]string)
	for _, v := range item {
		dataMap[v.Key] = v.Value
	}
	return dataMap
}
func (*ListMapItem) ToList(data map[string]string) []ListMapItem {
	list := make([]ListMapItem, 0)
	for k, v := range data {
		list = append(list, ListMapItem{
			Key: k, Value: v,
		})
	}
	return list
}
