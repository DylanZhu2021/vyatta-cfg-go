package command

import (
	"fmt"
	"strconv"
	"strings"
	"vyatta-cfg-go/internal/service/logic/activecfg"
	"vyatta-cfg-go/internal/service/logic/command/operatetree"
	"vyatta-cfg-go/internal/service/logic/tmplate"
	"vyatta-cfg-go/internal/service/logic/tools"
	"vyatta-cfg-go/internal/service/logic/tools/gqueue"
)

// GetCmdPathValueNode 获取命令中的路径（针对有值节点）
//逻辑：命令转为队列，将第一个元素出队,不保留最后一个元素
func GetCmdPathValueNode(setCmd string) string {
	var queue gqueue.QueueItf
	var pathQue gqueue.QueueItf

	queue = gqueue.New()
	pathQue = gqueue.New()
	queue.GetPathQue(setCmd)
	queue.Pop()
	length := queue.Length()
	for i := 0; i < length-1; i++ {
		pathQue.Push(queue.Pop())
	}
	return pathQue.GetPathString()
}

// GetCmdPathWithoutValueNode 获取命令中的路径（针对无值节点）
//逻辑：命令转为队列，将第一个元素出队,保留最后一个元素
func GetCmdPathWithoutValueNode(setCmd string) string {
	var queue gqueue.QueueItf
	var pathQue gqueue.QueueItf

	queue = gqueue.New()
	pathQue = gqueue.New()
	queue.GetPathQue(setCmd)
	queue.Pop()
	length := queue.Length()
	for i := 0; i < length; i++ {
		pathQue.Push(queue.Pop())
	}
	return pathQue.GetPathString()
}

// GetCmdValue 从一个set命令中获取到需要设置的值
func GetCmdValue(setCmd string, isValueNode bool) string {

	//判断是不是无值节点
	if isValueNode {
		return ""
	}

	setCmdArr := tools.GetPathArr(setCmd)
	return setCmdArr[len(setCmdArr)-1]
}

//// GetSetTree 获取WCfg的树
//func GetSetTree() *Operatetree.Node {
//
//	jsonstr := "{\n  \"name\": \"\",\n  \"value\": null,\n  \"childNode\": [\n    {\n      \"name\": \"interfaces\",\n     " +
//		" \"value\": null,\n      \"childNode\": [\n        {\n          \"name\": \"ethernet\",\n          \"value\": null,\n         " +
//		" \"childNode\": [\n            {\n              \"name\": \"eth0\",\n              \"value\": null,\n              \"childNode\": [\n    " +
//		"            {\n                  \"name\": \"address\",\n                  \"value\": [\n                    \"dhcp\",\n                    " +
//		"\"192.168.0.1/24\",\n                    \"101010\"\n                  ],\n                  \"childNode\": null\n                },\n      " +
//		"          {\n                  \"name\": \"description\",\n                  \"value\": [\n                    \"OUTSIDE\"\n                 " +
//		" ],\n                  \"childNode\": null\n                },\n                {\n                  \"name\": \"hw-id\",\n                 " +
//		" \"value\": [\n                    \"00:0c:29:8f:74:b1\"\n                  ],\n                  \"childNode\": null\n                }\n     " +
//		"         ]\n            }\n          ]\n        },\n        {\n          \"name\": \"loopback\",\n          \"value\": null,\n         " +
//		" \"childNode\": [\n            {\n              \"name\": \"lo\",\n              \"value\": null,\n              \"childNode\": [\n        " +
//		"        {\n                  \"name\": \"address\",\n                  \"value\": [\n                    \"123zxzx\"\n                 " +
//		" ],\n                  \"childNode\": null\n                }\n              ]\n            }\n          ]\n        }\n      ]\n    },\n    {\n   " +
//		"   \"name\": \"system\",\n      \"value\": null,\n      \"childNode\": [\n        {\n          \"name\": \"ntp\",\n          \"value\": null,\n    " +
//		"      \"childNode\": [\n            {\n              \"name\": \"server\",\n              \"value\": null,\n              \"childNode\": [\n        " +
//		"        {\n                  \"name\": \"time1.vyos.net\",\n                  \"value\": null,\n                  \"childNode\": [\n                  " +
//		"  {\n                      \"name\": \"pool\",\n                      \"value\": null,\n                      \"childNode\": null\n                   " +
//		"" + " }\n                  ]\n                }\n              ]\n            }\n          ]\n        }\n      ]\n    }\n  ]\n}"
//
//	var tree *Operatetree.Node
//	tree = new(Operatetree.Node)
//	json.Unmarshal([]byte(jsonstr), tree)
//
//	return tree
//}

//// GetDelTree 获取WCfg的树
//func GetDelTree() *Operatetree.Node {
//
//	jsonstr := "{\"name\":\"\",\"value\":null,\"childNode\":[{\"name\":\"interfaces\",\"value\":null," +
//		"\"childNode\":[{\"name\":\"ethernet\",\"value\":null,\"childNode\":[{\"name\":\"eth0\",\"value\":" +
//		"null,\"childNode\":[{\"name\":\"address\",\"value\":[\"dhcp\",\"192.168.0.1/24\",\"101010\"],\"childNode\"" +
//		":null},{\"name\":\"description\",\"value\":[\"OUTSIDE\"],\"childNode\":null},{\"name\":\"hw-id\",\"value\":" +
//		"[\"00:0c:29:8f:74:b1\"],\"childNode\":null},{\"name\":\"ip\",\"value\":null,\"childNode\":[{\"name\":\"enabl" +
//		"e-arp-ignore\",\"value\":null,\"childNode\":null}]}]}]},{\"name\":\"loopback\",\"value\":null,\"childNode\":" +
//		"[{\"name\":\"lo\",\"value\":null,\"childNode\":[{\"name\":\"address\",\"value\":[\"123zxzx\"],\"childNode\":" +
//		"null},{\"name\":\"ip\",\"value\":null,\"childNode\":[{\"name\":\"source-validation\",\"value\":[\"strict\"]," +
//		"\"childNode\":null}]}]},{\"name\":\"zx\",\"value\":null,\"childNode\":null}]}]},{\"name\":\"system\",\"value\":nul" +
//		"l,\"childNode\":[{\"name\":\"ntp\",\"value\":null,\"childNode\":[{\"name\":\"server\",\"value\":null,\"childNode\":" +
//		"[{\"name\":\"time1.vyos.net\",\"value\":null,\"childNode\":[{\"name\":\"pool\",\"value\":null,\"childNode\":null}]}]}" +
//		"]},{\"name\":\"name-server\",\"value\":[\"8.8.8.8\",\"9.9.9.9\"],\"childNode\":null},{\"name\":\"host-name\",\"value\":" +
//		"[\"1234\"],\"childNode\":null}]}]}\n"
//	var tree *operatetree.Node
//	tree = new(operatetree.Node)
//	json.Unmarshal([]byte(jsonstr), tree)
//
//	return tree
//}

// AddNodeFromTmp 往操作树里面加一个节点！！
func AddNodeFromTmp(queue gqueue.QueueItf, value string, tmpPath string, isNoValueNode bool, operate string) (*operatetree.Node, error) {
	var (
		root, node *operatetree.Node
	)

	root = new(operatetree.Node)
	node = new(operatetree.Node)

	//root是最后返回的节点，在变动的是node
	node = root

	for {
		prio := GetPriority(tmpPath)
		if queue.IsEmpty() {
			break
		}

		pathArr := strings.Split(tmpPath, " ")
		pathFile := strings.Join(pathArr, "/")
		if queue.Length() == 1 {
			if !isNoValueNode {
				if tools.IsTag(pathFile) {
					var operateArr []string
					var v []string
					operateArr = append(operateArr, operate)
					v = append(v, value)
					node.ChildNode = append(node.ChildNode, &operatetree.Node{
						Name:      fmt.Sprintf("%v", queue.Peek()),
						Value:     v,
						ChildNode: nil,
						Priority:  prio,
						Operate:   operateArr,
						Path:      tmpPath,
						IsLeaf:    true,
					})
					node = node.ChildNode[0]
				} else {
					var operateArr []string
					var v []string
					operateArr = append(operateArr, operate)
					v = append(v, value)
					node.ChildNode = append(node.ChildNode, &operatetree.Node{
						Name:      fmt.Sprintf("%v", queue.Peek()),
						Value:     v,
						ChildNode: nil,
						Priority:  prio,
						Operate:   operateArr,
						Path:      tmpPath + " " + fmt.Sprintf("%v", queue.Peek()),
						IsLeaf:    true,
					})
					node = node.ChildNode[0]
				}

			} else {
				//如果是无值节点
				var operateArr []string
				operateArr = append(operateArr, operate)
				node.ChildNode = append(node.ChildNode, &operatetree.Node{
					Name:      fmt.Sprintf("%v", queue.Peek()),
					Value:     nil,
					ChildNode: nil,
					Priority:  prio,
					Operate:   operateArr,
					Path:      strings.Trim(tmpPath+" "+fmt.Sprintf("%v", queue.Peek()), " "),
					IsLeaf:    true,
				})
				node = node.ChildNode[0]
			}
		} else {
			node.ChildNode = append(node.ChildNode, &operatetree.Node{
				Name:      fmt.Sprintf("%v", queue.Peek()),
				Value:     nil,
				ChildNode: nil,
				Operate:   nil,
				Priority:  prio,
				Path:      strings.Trim(tmpPath+" "+fmt.Sprintf("%v", queue.Peek()), " "),
				IsLeaf:    false,
			})
			node = node.ChildNode[0]
		}
		if tmpPath == "" {
			tmpPath += fmt.Sprintf("%v", queue.Pop())
		} else {
			tmpPath += " " + fmt.Sprintf("%v", queue.Pop())
		}
	}
	return root, nil
}

// DeleteStructSliceElms 删除切片指定元素(传入传出都是Node的切片)
func DeleteStructSliceElms(sl []*activecfg.Node, elm *activecfg.Node) []*activecfg.Node {

	j := 0
	for _, v := range sl {
		if v.Name != elm.Name {
			sl[j] = v
			j++
		}
	}
	return sl[:j]
}

// DeleteSlice3 删除指定元素。

// GetPriority 获取节点的优先级
//传入路径 eg: "interfaces ethernet"
func GetPriority(path string) int {

	if path == "" {
		return 0
	}
	var queue gqueue.QueueItf

	queue = gqueue.New()

	queue.GetPathQue(path)
	tmpNode, err := tmplate.GetTmpNode(queue)
	if err != nil {
		return 0
	}
	//由于优先级的字符串前面有个空格，先去掉空格，不然转成int型后为0
	prioStr := strings.Trim(tmpNode.Priority, " ")
	prioInt, _ := strconv.Atoi(prioStr)
	return prioInt
}

// Append 切片添加，去重复
func Append(arr []string, elem string) []string {

	for _, v := range arr {
		if v == elem {
			return arr
		}
	}
	return append(arr, elem)
}
