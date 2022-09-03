package vshow

import "github.com/gogf/gf/v2/frame/g"

type ShowReq struct {
	g.Meta `path:"/show" method:"Get" tags:"ShowService" summary:"show configuration"`
}

type ShowRes struct {
	Ret string `json:"ret"`
}
