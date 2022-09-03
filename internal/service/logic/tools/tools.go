package tools

import (
	"io/ioutil"
	"strings"
)

// GetPathArr 把路径的string转为[]string
//可以后空格分隔实现
func GetPathArr(path string) []string {
	countSplit := strings.Split(path, " ")
	return countSplit
}

// GetCheckPath add tmp root path
func GetCheckPath(path string) string {
	return TmpRootPath + "/" + path
}

// IsTag 检查节点是不是tag目录
// 传入：eg:interfaces/ethernet
func IsTag(path string) bool {
	path = GetCheckPath(path)
	files, _ := ioutil.ReadDir(path)

	for _, f := range files {
		if f.Name() == NodeTag {
			return true
		}
	}
	return false
}

//// IsTagNode 判断是不是tag节点
//// 传入节点的path字段 eg:"interfaces ethernet eth0" =>return true
//func IsTagNode(path string) bool {
//	var queue gqueue.QueueItf
//	var filePath string
//	filePath = ""
//	queue = gqueue.New()
//	queue.GetPathQue(path)
//
//	for !queue.IsEmpty() {
//
//	}
//	return false
//}

// DeleteStrSliceElms 删除切片指定字符串元素(传入传出都是[]string)
func DeleteStrSliceElms(sl []string, elm string) []string {

	j := 0
	for _, v := range sl {
		if v != elm {
			sl[j] = v
			j++
		}
	}
	return sl[:j]
}
