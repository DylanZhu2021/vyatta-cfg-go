package vservice

import "github.com/gogf/gf/v2/frame/g"

type ModifyServiceAAUReq struct {
	g.Meta  `path:"/service/ssh/access-control/allow/user" method:"post" tags:"ServiceService" summary:"modify: service ssh access-control allow user"`
	Operate string `json:"operate" v:"required|in:set,delete#operate must given|operate must select in set and delete" dc:"operate to node,set or delete"`
	Data    string `json:"data" v:"required" dc:"value you want to set or delete"`
}

type ModifyServiceAAURes struct {
	Ret string `json:"ret"`
}

type ModifyServicePReq struct {
	g.Meta  `path:"/service/ssh/port" method:"post" tags:"SystemService" summary:"modify: service ssh port"`
	Operate string `json:"operate"p:"operate" v:"required|in:set,delete#operate must given|operate must select in set and delete" dc:"operate to node,set or delete"`
	Data    string `json:"data"p:"data" v:"required|between:0,65535#data(port)must given|data(port) must between 0 and 65535" dc:"value you want to set or delete"`
}

type ModifyServicePRes struct {
	Ret string `json:"ret"`
}
