package tmplate

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"vyatta-cfg-go/internal/service/logic/tools"
	"vyatta-cfg-go/internal/service/logic/tools/gqueue"
)

// CheckPathValidInTmp 检查path路劲是否在模板文件中有效
func CheckPathValidInTmp(path string) bool {
	//获取到path的queue
	var queue gqueue.QueueItf
	var dirPath string
	queue = gqueue.New()
	queue.GetPathQue(path)

	dirPath = queue.GetFilePath()
	_, err := ioutil.ReadDir(dirPath)
	if err != nil {
		//表示没有对应的目录
		return false
	}
	return true

}

// IsMulti 检查节点的Multi属性
//path eg :interfaces ethernet
func IsMulti(path string) bool {
	//获取到path的queue
	var queue gqueue.QueueItf
	queue = gqueue.New()
	queue.GetPathQue(path)

	tmpNode, err := GetTmpNode(queue)
	if err != nil {
		log.Panicf("检查节点Multi出错:%v", err)
		return false
	}
	return tmpNode.GetTmpMulti()
}

// IsLeaf 根据模板文件判断节点是不是叶子节点
//传入的path eg:interfaces ethernet
// IsLeaf
func IsLeaf(path string) bool {
	//先将path转为含有node.tag的路径
	var queue gqueue.QueueItf
	queue = gqueue.New()
	queue.GetPathQue(path)
	path = queue.GetFilePath()

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return false
	}
	count := len(files)
	var filename string
	for _, v := range files {
		filename = v.Name()
	}
	if count == 1 && filename == tools.NodeDef {
		return true
	}
	return false
}

// IsNoValueNode 判断节点是否是拥有值的节点
//无值节点的判定：
//1.先判断是不是叶子节点
//2.没有type的字段的节点就是无值节点
//返回true为无值节点
// 传入：interfaces ethernet eth0 address
func IsNoValueNode(setPath string) (bool, error) {

	//获取到set命令的queue
	var queue gqueue.QueueItf
	var pathQueue gqueue.QueueItf
	pathQueue = gqueue.New()
	queue = gqueue.New()
	queue.GetPathQue(setPath)

	pathArr := strings.Split(setPath, " ")
	pathFile := strings.Join(pathArr, "/")
	if tools.IsTag(pathFile) {
		return false, nil
	}
	tmpPath := ""

	for {
		//如果队空出队
		if queue.IsEmpty() {
			break
		}
		if tools.IsTag(tmpPath) && queue.Length()-2 == 0 {
			return false, nil
		}
		//如果不是叶子节点进入下一轮循环
		if !IsLeaf(tmpPath) {
			if tmpPath == "" {
				tmpPath += fmt.Sprintf("%v", queue.Pop())
			} else {
				tmpPathArr := strings.Split(tmpPath, " ")
				tmpPathFile := strings.Join(tmpPathArr, "/")
				if tools.IsTag(tmpPathFile) {
					tmpPath += " " + tools.NodeTag
					queue.Pop()
				} else {
					tmpPath += " " + fmt.Sprintf("%v", queue.Pop())
				}
			}
			continue
		} else {
			break
		}

	}

	pathQueue.GetPathQue(tmpPath)
	node, err := GetTmpNode(pathQueue)
	if err != nil {
		return false, err
	}
	if node.Type == "" {
		return true, nil
	}
	return false, nil

}
