package tmplate

import (
	"container/list"
	"errors"
	"fmt"
	"vyatta-cfg-go/internal/service/logic/tools"
	"vyatta-cfg-go/internal/service/logic/tools/gqueue"
)

/*
	Multi     bool
	Tag       bool
	Type      string
	Help      string
	Comp_help string
	Val_help  string
	Allowed   string
	Syntax    string
	Commit    string
	Priority  string
	DefValue  string
	Begin     string
	End       string
	Create    string
	Update    string
	Delete    string
*/

// GetTmpNode 根据path（queue）得到对应的模板节点
func GetTmpNode(pathQue gqueue.QueueItf) (*NodeTmp, error) {
	//******
	tmpPath := ""
	isTag := false
	count := pathQue.Length()

	pathQue.Iterator(func(e *list.Element) bool {
		if tmpPath == "" {
			tmpPath += fmt.Sprintf("%v", e.Value)
		} else {
			tmpPath += "/" + fmt.Sprintf("%v", e.Value)
		}
		if tools.IsTag(tmpPath) {
			isTag = true
			return false
		}
		count--
		return true
	})
	//******
	var tmpPathAll string
	if isTag && count == 2 { //
		tmpPathAll = pathQue.GetFilePath()
	} else {
		tmpPathAll = pathQue.GetFilePath() + "/" + tools.NodeDef
	}

	if CheckPathValidInTmp(pathQue.GetPathString()) {
		var tmpNode *NodeTmp
		tmpNode = ReadTmpToStruct(tmpPathAll)
		return tmpNode, nil
	} else {
		return nil, errors.New("path(queue) is not valid！")
	}
}

// GetTmpTag 返回模板节点tag值
func (tmpNode *NodeTmp) GetTmpTag() bool {
	return tmpNode.Tag
}

// GetTmpMulti 返回模板节点Multi值
func (tmpNode *NodeTmp) GetTmpMulti() bool {
	return tmpNode.Multi
}

// GetTmpType 返回模板节点默认值
func (tmpNode *NodeTmp) GetTmpType() string {
	return tmpNode.Type
}

// GetTmpDefault 返回模板节点的默认值
func (tmpNode *NodeTmp) GetTmpDefault() string {
	return tmpNode.DefValue
}

func (tmpNode *NodeTmp) GetTmpHelp() string {
	return tmpNode.DefValue
}
func (tmpNode *NodeTmp) GetTmpComHelp() string {
	return tmpNode.Comp_help
}
func (tmpNode *NodeTmp) GetTmpValHelp() string {
	return tmpNode.Val_help
}
func (tmpNode *NodeTmp) GetTmpSyntax() string {
	return tmpNode.Syntax
}
func (tmpNode *NodeTmp) GetTmpPrior() string {
	return tmpNode.Priority
}
