package controller

import (
	"context"
	"fmt"
	"vyatta-cfg-go/api/v1/vprotocols"
	"vyatta-cfg-go/internal/model"
	"vyatta-cfg-go/internal/service"
)

type Protocols struct{}

func (c *Protocols) HandlerProtocolsctx(ctx context.Context,
	req *vprotocols.ModifyProtocolsReq) (res *vprotocols.ModifyProtocolsRes, err error) {

	ret, err := service.Modify().DoModifyProtocols(ctx, model.DoModify{
		Data:    req.Data,
		Operate: req.Operate,
	})
	if err != nil {
		return res, err
	}
	res = &vprotocols.ModifyProtocolsRes{
		Ret: ret,
	}
	fmt.Println(res.Ret)
	return res, nil

}
