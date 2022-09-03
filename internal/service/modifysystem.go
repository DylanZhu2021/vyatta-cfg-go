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

func (m sModify) DoModifySystem(ctx context.Context, modify model.DoModify) (string, error) {
	var acfTree *activecfg.Node
	var treeBuf bytes.Buffer
	//"set system host-name 123123"
	cmds := []string{
		fmt.Sprintf("%s system %s %s", modify.Operate, modify.Leaf, modify.Data),
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
