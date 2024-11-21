package utils

import "kubeimook/model/base"

func ToMap(item []base.ListMapItem) map[string]string {
	dataMap := make(map[string]string)
	for _, v := range item {
		dataMap[v.Key] = v.Value
	}
	return dataMap
}
func ToList(data map[string]string) []base.ListMapItem {
	list := make([]base.ListMapItem, 0)
	for k, v := range data {
		list = append(list, base.ListMapItem{
			Key: k, Value: v,
		})
	}
	return list
}
func ToListWithMapByte(data map[string][]byte) []base.ListMapItem {
	list := make([]base.ListMapItem, 0)
	for k, v := range data {
		list = append(list, base.ListMapItem{
			Key: k, Value: string(v),
		})
	}
	return list
}
