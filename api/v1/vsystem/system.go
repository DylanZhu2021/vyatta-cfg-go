package vsystem

import (
	"github.com/gogf/gf/v2/frame/g"
)

type ModifySystemReq struct {
	g.Meta  `path:"/system/:leaf" method:"post" tags:"SystemService" summary:"modify: system host-name or name-server"`
	Leaf    string `json:"leaf"  p:"leaf" v:"required#operate must given" dc:"leaf node,in this service,leaf select in host-name,name-server"`
	Operate string `json:"operate" v:"required|in:set,delete#operate must given|operate must select in set and delete" dc:"operate to node,set or delete"`
	Data    string `json:"data" v:"required" dc:"value you want to set or delete"`
}

type ModifySystemRes struct {
	Ret string `json:"ret,omitempty"`
}
