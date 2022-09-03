package commit

import (
	"errors"
	"fmt"
	"vyatta-cfg-go/internal/service/logic/activecfg"
	"vyatta-cfg-go/internal/service/logic/command"
	"vyatta-cfg-go/internal/service/logic/tmplate"
	"vyatta-cfg-go/internal/service/logic/tools"
	"vyatta-cfg-go/internal/service/logic/tools/gqueue"
)

func Delete(tree *activecfg.Node, path string, value string) error {

	var (
		isExist bool
		queue   gqueue.QueueItf
		pnode   *activecfg.Node
		cnode   *activecfg.Node
	)
	//初始化路劲队列
	queue = gqueue.New()
	queue.GetPathQue(path)

	cnode = tree

	for !queue.IsEmpty() {
		isExist = false
		for _, node := range cnode.ChildNode {
			if node.Name == queue.Peek() {
				isExist = true
				pnode = cnode
				cnode = node
				continue
			}
		}
		if !isExist {
			return errors.New(fmt.Sprintf("operate failed;path:%v", path))
		}
		queue.Pop()
	}
	if value != "" {
		//有值节点并且单值节点
		if !tmplate.IsMulti(path) {
			if value != cnode.Value[0] {
				return errors.New(fmt.Sprintf("delete value:%v failed(value is nit exist); path:%v", value, path))
			}
			cnode.Value = nil
			return nil
		}
		isExist = false
		for _, v := range cnode.Value {
			if value == v {
				isExist = true
				break
			}
		}
		if !isExist {
			return errors.New(fmt.Sprintf("delete value:%v failed(value is nit exist); path:%v", value, path))
		}
		cnode.Value = tools.DeleteStrSliceElms(cnode.Value, value)
		return nil
	}
	pnode.ChildNode = command.DeleteStructSliceElms(pnode.ChildNode, cnode)
	return nil
}
