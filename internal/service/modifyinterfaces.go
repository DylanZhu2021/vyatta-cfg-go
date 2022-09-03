package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"vyatta-cfg-go/internal/model"
	"vyatta-cfg-go/internal/service/logic/activecfg"
	"vyatta-cfg-go/internal/service/logic/commit"
)

func (m sModify) DoModifyInterfaces(ctx context.Context, modify model.DoModify) (string, error) {
	var acfTree *activecfg.Node
	var treeBuf bytes.Buffer

	//"set interfaces ethernet %s address %s"
	cmds := []string{
		fmt.Sprintf("%s interfaces ethernet %s %s %s", modify.Operate, modify.Tag, modify.Leaf, modify.Data),
	}
	fmt.Println(cmds)

	treejson := activecfg.GetTree()

	if err := json.Unmarshal([]byte(treejson), &acfTree); err != nil {
		return "", err
	}

	if err := commit.Commit(cmds, acfTree); err != nil {
		return "", err
	}

	treebyte, _ := json.Marshal(acfTree)

	//格式化json字符串
	if err := json.Indent(&treeBuf, treebyte, "", "    "); err != nil {
		return "", err
	}

	return treeBuf.String(), nil
}
