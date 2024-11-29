package utils

import "kubeimook/model/base"

func ToMap(item []base.ListMapItem) map[string]string {
	dataMap := make(map[string]string)
	for _, v := range item {
		dataMap[v.Key] = v.Value
	}
	return dataMap
}

// ToList 函数将给定的 map[string]string 类型数据转换为 base.ListMapItem 类型的切片。
//
// 参数:
//
//	data - map[string]string 类型，表示要转换的键值对数据。
//
// 返回值:
//
//	[]base.ListMapItem - 返回一个包含 base.ListMapItem 类型的切片，每个元素都包含原始 map 中的一个键值对。
//
// 说明:
// 该函数首先创建一个空的 base.ListMapItem 类型切片。然后，它遍历输入的 map 数据，
// 对于每个键值对，它创建一个新的 base.ListMapItem 对象，并将其添加到切片中。
// 最后，返回填充好的切片。
func ToList(data map[string]string) []base.ListMapItem {
	list := make([]base.ListMapItem, 0)
	for k, v := range data {
		list = append(list, base.ListMapItem{
			Key: k, Value: v,
		})
	}
	return list
}

// ToListWithMapByte 函数将给定的 map[string][]byte 类型数据转换为 base.ListMapItem 类型的切片。
//
// 参数:
//
//	data - map[string][]byte 类型，表示包含字符串键和字节切片值的映射。
//
// 返回值:
//
//	[]base.ListMapItem - 返回一个包含 base.ListMapItem 类型的切片，其中每个元素都包含原始映射中的一个键值对。
//	键保持不变，而值则从字节切片转换为字符串。
//
// 说明:
// 该函数首先创建一个空的 base.ListMapItem 类型切片。然后，它遍历输入的 map 数据，
// 对于每个键值对，它将字节切片值转换为字符串，并创建一个新的 base.ListMapItem 对象，
// 将其添加到切片中。最后，返回填充好的切片。
func ToListWithMapByte(data map[string][]byte) []base.ListMapItem {
	list := make([]base.ListMapItem, 0)
	for k, v := range data {
		list = append(list, base.ListMapItem{
			Key: k, Value: string(v),
		})
	}
	return list
}
