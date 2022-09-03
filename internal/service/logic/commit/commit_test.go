package commit

import (
	"encoding/json"
	"fmt"
	"testing"
	"vyatta-cfg-go/internal/service/logic/activecfg"
)

func TestCommit(t *testing.T) {
	setcmds := []string{
		//"set interfaces loopback lo ip source-validation strict",
		//"set interfaces loopback zx ip",
		//"set system name-server 123",
		//"set system name-server 456",
		//"set system name-server zxzx",
		//"set system name-server zxzx123",
		//"set system name-server zxzx123456",
		//"set system host-name 1234",
		//"set interfaces ethernet eth0 ip enable-arp-ignore",
		//"set interfaces ethernet eth0 address dhcp",
		//"set interfaces ethernet eth0 description OUTSIDE",
		//"set interfaces ethernet eth1 address 192.168.0.1/24",
		//"set interfaces ethernet eth1 description INSIDE",
		//"set nat source rule 100 outbound-interface eth0",
		//"set nat source rule 100 source address 192.168.0.0/24",
		//"set nat source rule 100 translation address masquerade",
		//"set system login user vyos",
		//"set interfaces ethernet eth0",
		//"set service ssh disable-password-authentication",
		//"set interfaces ethernet eth1 address 123",
		//"delete system host-name 1234",
		//"delete interfaces loopback lo ip source-validation strict",
		//"delete system name-server 123",
		//"delete system name-server zxzx",
		//"delete interfaces ethernet eth0 ip enable-arp-ignore",
		//"delete system login user vyos",
		//"delete service ssh disable-password-authentication", //无值节点
		//"delete interfaces ethernet eth0",
		"set interfaces ethernet eth2 address 123zxzxzxzx",
	}
	//delcmds := []string{
	//	//"delete system host-name 1234",
	//	//"delete interfaces loopback lo ip source-validation strict",
	//	//"delete system name-server 123",
	//	//"delete system name-server zxzx",
	//	////"delete interfaces ethernet eth0 ip enable-arp-ignore",
	//	//"delete system login user vyos",
	//	////"delete service ssh disable-password-authentication", //无值节点
	//	//"delete interfaces ethernet eth0",
	//}

	var tree *activecfg.Node
	str := activecfg.GetTree()
	fmt.Println("json:" + str)

	//str := "{\"name\":\"\",\"value\":null,\"childNode\":[{\"name\":\"interfaces\",\"value\":null,\"childNode\":[{\"name\":\"ethernet\",\"value\":null,\"childNode\":[{\"name\":\"eth0\",\"value\":null,\"childNode\":[{\"name\":\"address\",\"value\":[\"123123\",\"dhcp\",\"dhcp\"],\"childNode\":null},{\"name\":\"description\",\"value\":[\"OUTSIDE\"],\"childNode\":null},{\"name\":\"ip\",\"value\":null,\"childNode\":[{\"name\":\"enable-arp-ignore\",\"value\":null,\"childNode\":null}]}]},{\"name\":\"eth1\",\"value\":null,\"childNode\":[{\"name\":\"description\",\"value\":[\"INSIDE\"],\"childNode\":null},{\"name\":\"address\",\"value\":[\"192.168.0.1/24\",\"192.168.0.1/24\"],\"childNode\":null}]}]},{\"name\":\"loopback\",\"value\":null,\"childNode\":[{\"name\":\"zx\",\"value\":null,\"childNode\":[{\"name\":\"ip\",\"value\":null,\"childNode\":null}]},{\"name\":\"lo\",\"value\":null,\"childNode\":[{\"name\":\"ip\",\"value\":null,\"childNode\":[{\"name\":\"source-validation\",\"value\":[\"strict\"],\"childNode\":null}]}]}]}]},{\"name\":\"system\",\"value\":null,\"childNode\":[{\"name\":\"name-server\",\"value\":[\"123\",\"456\",\"zxzx\",\"zxzx123\",\"zxzx123456\",\"123\",\"456\",\"zxzx\",\"zxzx123\",\"zxzx123456\"],\"childNode\":null},{\"name\":\"host-name\",\"value\":[\"1234\"],\"childNode\":null},{\"name\":\"login\",\"value\":null,\"childNode\":[{\"name\":\"user\",\"value\":null,\"childNode\":[{\"name\":\"vyos\",\"value\":null,\"childNode\":null}]}]}]},{\"name\":\"service\",\"value\":null,\"childNode\":[{\"name\":\"ssh\",\"value\":null,\"childNode\":[{\"name\":\"disable-password-authentication\",\"value\":null,\"childNode\":null}]}]},{\"name\":\"nat\",\"value\":null,\"childNode\":[{\"name\":\"source\",\"value\":null,\"childNode\":[{\"name\":\"rule\",\"value\":null,\"childNode\":[{\"name\":\"100\",\"value\":null,\"childNode\":[{\"name\":\"source\",\"value\":null,\"childNode\":[{\"name\":\"address\",\"value\":[\"192.168.0.0/24\"],\"childNode\":null}]},{\"name\":\"outbound-interface\",\"value\":[\"eth0\"],\"childNode\":null},{\"name\":\"translation\",\"value\":null,\"childNode\":[{\"name\":\"address\",\"value\":[\"masquerade\"],\"childNode\":null}]}]}]}]}]}]}\n"
	json.Unmarshal([]byte(str), &tree)
	err := Commit(setcmds, tree)
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonByte, _ := json.Marshal(tree)
	fmt.Println(string(jsonByte))
	//
	//err = Commit(delcmds, tree)
	//if err != nil {
	//	fmt.Println(err)
	//	//return
	//}
	//jsonByte, _ = json.Marshal(tree)
	//fmt.Println(string(jsonByte))
}
