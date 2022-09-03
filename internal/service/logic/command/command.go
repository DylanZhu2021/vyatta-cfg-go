package command

import (
	"fmt"
	"strings"
	"vyatta-cfg-go/internal/service/logic/command/operatetree"
	"vyatta-cfg-go/internal/service/logic/tmplate"
	"vyatta-cfg-go/internal/service/logic/tools"
	"vyatta-cfg-go/internal/service/logic/tools/gqueue"
)

/*
	生成操作树!!
*/

// Command 生成操作树
func Command(finalTree *operatetree.Node, cmd string) error {

	//获取到WCfg的树
	var queue gqueue.QueueItf
	var path string
	var isNoValueNode bool

	tree := finalTree

	queue = gqueue.New()

	pathArr := strings.Split(cmd, " ")
	//获取到命令的操作
	operate := pathArr[0]
	pathArr = pathArr[1:]
	cmdPath := strings.Join(pathArr, " ")

	//处理tag节点
	tagPath := strings.Join(pathArr[0:len(pathArr)-1], "/")
	if tools.IsTag(tagPath) {
		isNoValueNode = false
	} else {
		//判断是不是无值节点
		ret, err := tmplate.IsNoValueNode(cmdPath)
		isNoValueNode = ret
		if err != nil {
			fmt.Println(err)
			return err
		}

	}

	if !isNoValueNode {
		//若不是无值节点
		//获取命令的路径及对应的队列
		path = GetCmdPathValueNode(cmd)
		queue.GetPathQue(path)
	} else {
		path = GetCmdPathWithoutValueNode(cmd)
		queue.GetPathQue(path)
	}
	//获取设置的值
	value := GetCmdValue(cmd, isNoValueNode)

	tmpPath := ""
	flag := false
	//isTag := false

	for !queue.IsEmpty() {
		//匹配到了叶子节点，则退出
		if tmplate.IsLeaf(tmpPath) {
			break
		}

		for _, v := range tree.ChildNode {
			if v.Name == queue.Peek() {
				flag = true
				tree = v
				if tmpPath == "" {
					tmpPath += fmt.Sprintf("%v", queue.Pop())
				} else {
					tmpPath += " " + fmt.Sprintf("%v", queue.Pop())
				}
				break
			}
		}

		if flag {
			flag = false
			continue
		} else {
			//树中未匹配成功
			//检查模板中是否匹配
			addNode, err := AddNodeFromTmp(queue, value, tmpPath, isNoValueNode, operate)
			if err != nil {
				return err
			}

			tree.ChildNode = append(tree.ChildNode, addNode.ChildNode[0])
			return nil
		}
	}

	prio := GetPriority(tmpPath)

	if value != "" {
		tree.Value = append(tree.Value, value) //valueArr
	} else {
		tree.Value = nil
	}
	tree.Operate = append(tree.Operate, operate)
	tree.Priority = prio
	tree.Path = tmpPath
	//if isTag {
	//	tree.Path = tmpPath
	//}else{
	//
	//}
	fmt.Println(tree.Path)
	tree.IsLeaf = true

	return nil
}
