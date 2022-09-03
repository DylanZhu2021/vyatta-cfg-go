package controller

import (
	"context"
	"fmt"
	"vyatta-cfg-go/api/v1/vsystem"
	"vyatta-cfg-go/internal/model"
	"vyatta-cfg-go/internal/service"
)

type System struct{}

func (c *System) HandlerSystem(ctx context.Context,
	req *vsystem.ModifySystemReq) (res *vsystem.ModifySystemRes, err error) {
	ret, err := service.Modify().DoModifySystem(ctx, model.DoModify{
		Data:    req.Data,
		Operate: req.Operate,
		Leaf:    req.Leaf,
	})
	if err != nil {
		return res, err
	}
	res = &vsystem.ModifySystemRes{
		Ret: ret,
	}
	fmt.Println(res.Ret)
	return res, nil
}
