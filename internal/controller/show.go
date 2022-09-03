package controller

import (
	"context"
	"vyatta-cfg-go/api/v1/vshow"
	"vyatta-cfg-go/internal/model"
	"vyatta-cfg-go/internal/service"
)

type Show struct{}

func (c *Show) HandlerShow(ctx context.Context,
	req *vshow.ShowReq) (res *vshow.ShowRes, err error) {
	ret, err := service.Modify().DoShow(ctx, model.DoModify{})
	if err != nil {
		return res, err
	}
	res = &vshow.ShowRes{
		Ret: ret,
	}
	return res, nil
}
