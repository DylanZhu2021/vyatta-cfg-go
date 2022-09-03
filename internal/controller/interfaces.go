package controller

import (
	"context"
	"fmt"
	"vyatta-cfg-go/api/v1/vinterfaces"
	"vyatta-cfg-go/internal/model"
	"vyatta-cfg-go/internal/service"
)

/*
如果要显示/json和/swagger这两个页面的话，需要注意：
1.Handler函数必须是如下格式：func Xxx(ctx context.Context,req XxxReq) (res XxxRes, err error)
2.此外不是必须使用对象注册
(这里统一使用对象注册)
*/

type Interfaces struct{}

func (c *Interfaces) HandlerInterfaces(ctx context.Context,
	req *vinterfaces.ModifyInterfacesReq) (res *vinterfaces.ModifyInterfacesRes, err error) {

	ret, err := service.Modify().DoModifyInterfaces(ctx, model.DoModify{
		Tag:     req.Tag,
		Data:    req.Data,
		Operate: req.Operate,
		Leaf:    req.Leaf,
	})
	if err != nil {
		return nil, err
	}
	res = &vinterfaces.ModifyInterfacesRes{
		Ret: ret,
	}
	fmt.Println(res.Ret)
	return res, nil
}
