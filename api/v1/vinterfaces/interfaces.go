package vinterfaces

import "github.com/gogf/gf/v2/frame/g"

type ModifyInterfacesReq struct {
	g.Meta  `path:"/interfaces/ethernet/:tag/:leaf" method:"post" tags:"InterfaceService" summary:"modify: interfaces ethernet tag address(mac...)"`
	Tag     string `json:"tag"p:"tag" v:"required|regex:(eth|lan)[0-9]+|(eno|ens|enp|enx).+#tag must given|ethernet name is not validate" dc:"name of tag node"`
	Operate string `json:"operate"v:"required|in:set,delete#operate must given|operate must select in set and delete" dc:"operate to node,set or delete"`
	Data    string `json:"data"v:"required#data must given" dc:"value you want to set or delete"`
	Leaf    string `json:"leaf"p:"leaf" v:"required#leaf must given" dc:"leaf node,in this service,leaf select in mac,address,mtu"`
}
type ModifyInterfacesRes struct {
	Ret string `json:"ret"`
}
