package service

import (
	"context"
	"encoding/json"
	"vyatta-cfg-go/internal/model"
	"vyatta-cfg-go/internal/service/logic/activecfg"
	"vyatta-cfg-go/internal/service/logic/show"
)

func (m sModify) DoShow(ctx context.Context, modify model.DoModify) (string, error) {

	treeStr := activecfg.GetTree()
	var tree *activecfg.Node
	err := json.Unmarshal([]byte(treeStr), &tree)
	if err != nil {
		return "", err
	}

	return show.Show(tree)
}
