package command

import (
	"encoding/json"
	"fmt"
	"testing"
	"vyatta-cfg-go/internal/service/logic/command/operatetree"
)

//func TestCommandSet(t *testing.T) {
//
//	//setCmd := "set system ntp server time1.vyos.net pool"
//	//setCmd1 := "set interfaces ethernet eth0 ip enable-arp-ignore"
//	//setCmd := "set system host-name 123"
//	//setCmd := "set interfaces loopback lo ip source-validation strict"
//	//setCmd := "set interfaces loopback zx"
//	//setCmd1 := "set system host-name 123zxzx"
//	//setCmd := "set interfaces ethernet eth1 address"
//	//setCmd := "set protocols bgp address-family ipv4-vpn network zx"
//
//	cmds := []string{
//		"set interfaces loopback lo ip source-validation strict",
//		"set interfaces loopback zx",
//		"set system name-server 8.8.8.8",
//		"set system name-server 9.9.9.9",
//		"set system host-name 123",
//		"set system host-name 1234",
//		"set interfaces ethernet eth0 ip enable-arp-ignore",
//	}
//	tree := GetSetTree()
//
//	jsonstr, _ := json.Marshal(tree)
//	fmt.Printf("%v\n", string(jsonstr))
//	finalTree := tree
//
//	for k, cmd := range cmds {
//		err := Command(tree, cmd)
//		if err != nil {
//			fmt.Println(k)
//			fmt.Println("set:", err)
//			return
//		}
//	}
//
//	//err := Command(tree, setCmd)
//	//if err != nil {
//	//	fmt.Println("set:", err)
//	//	return
//	//}
//
//	jsonstr, _ = json.Marshal(finalTree)
//	fmt.Printf("%v\n", string(jsonstr))
//}

func TestCommand(t *testing.T) {
	var tree *operatetree.Node
	tree = new(operatetree.Node)

	cmds := []string{
		//"set interfaces loopback lo ip source-validation strict",
		//"delete interfaces loopback lo ip source-validation strict",
		"set interfaces loopback zx ip",
		// "set system name-server 123",
		//"delete system name-server 123",
		"set system name-server 456",
		"set system name-server zxzx",
		"set system name-server zxzx123",
		"set system name-server zxzx123456",
		//"set system host-name 123",
		//"delete system host-name 123",
		"set system host-name 1234",
		//"set interfaces ethernet eth0 ip enable-arp-ignore",
		//"delete interfaces ethernet eth0 ip enable-arp-ignore",
		"set interfaces ethernet eth0 address dhcp",
		"set interfaces ethernet eth0 description OUTSIDE",
		"set interfaces ethernet eth1 address 192.168.0.1/24",
		"set interfaces ethernet eth1 description INSIDE",
		"set nat source rule 100 outbound-interface eth0",
		"set nat source rule 100 source address 192.168.0.0/24",
		"set nat source rule 100 translation address masquerade",
		//"set system login user vyos",
		//"delete system login user vyos",
		"set service ssh disable-password-authentication",
	}
	for k, cmd := range cmds {
		err := Command(tree, cmd)
		if err != nil {
			fmt.Printf("%v,err:%v", k, err)
			return
		}
	}
	str, err := json.Marshal(tree)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(str))

	//err := Command(tree, "set interfaces loopback lo ip source-validation strict")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//str, err := json.Marshal(tree)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(string(str))
	//err = Command(tree, "delete interfaces loopback lo ip source-validation strict")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//str, err = json.Marshal(tree)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(string(str))
	//err = Command(tree, "delete interfaces loopback lo ip")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//str, err = json.Marshal(tree)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(string(str))
	//err = Command(tree, "set interfaces ethernet eth0 ip enable-arp-ignore")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//str, err = json.Marshal(tree)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(string(str))

}
