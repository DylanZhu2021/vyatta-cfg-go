package controller

import (
	"context"
	"fmt"
	"vyatta-cfg-go/api/v1/vservice"
	"vyatta-cfg-go/internal/model"
	"vyatta-cfg-go/internal/service"
)

type Service struct{}

func (c *Service) HandlerServiceP(ctx context.Context,
	req *vservice.ModifyServicePReq) (res *vservice.ModifyServicePRes, err error) {

	ret, err := service.Modify().DoModifyServiceP(ctx, model.DoModify{
		Data:    req.Data,
		Operate: req.Operate,
	})
	if err != nil {
		return res, err
	}
	res = &vservice.ModifyServicePRes{
		Ret: ret,
	}
	fmt.Println(res.Ret)
	return res, nil
}

func (c *Service) HandlerServiceAAU(ctx context.Context,
	req *vservice.ModifyServiceAAUReq) (res *vservice.ModifyServiceAAURes, err error) {
	ret, err := service.Modify().DoModifyServiceAAU(ctx, model.DoModify{
		Data:    req.Data,
		Operate: req.Operate,
	})
	if err != nil {
		res = &vservice.ModifyServiceAAURes{
			Ret: fmt.Sprintf("%s", err),
		}
		return res, err
	}
	res = &vservice.ModifyServiceAAURes{
		Ret: ret,
	}
	fmt.Println(res.Ret)
	return res, nil
}
