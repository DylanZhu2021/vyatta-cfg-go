package vprotocols

import "github.com/gogf/gf/v2/frame/g"

type ModifyProtocolsReq struct {
	g.Meta  `path:"/protocols/static/route" method:"post" tags:"ProtocolsService" summary:"modify: protocols static route"`
	Operate string `json:"operate"v:"required|in:set,delete#operate must given|operate must select in set and delete" dc:"operate to node,set or delete"`
	Data    string `json:"data"v:"required#data must given" dc:"value you want to set or delete"`
}

type ModifyProtocolsRes struct {
	Ret string `json:"ret"`
}
