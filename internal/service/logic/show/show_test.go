package show

import (
	"encoding/json"
	"fmt"
	"testing"
	"vyatta-cfg-go/internal/service/logic/activecfg"
)

func TestShow(t *testing.T) {

	var tree *activecfg.Node

	str := "{\"name\":\"\",\"value\":null,\"childNode\":[{\"name\":\"interfaces\",\"value\":null,\"childNode\":[{\"name\":\"ethernet\",\"value\":null,\"childNode\":[{\"name\":\"eth0\",\"value\":null,\"childNode\":[{\"name\":\"address\",\"value\":[\"123123\",\"dhcp\"],\"childNode\":null},{\"name\":\"description\",\"value\":[\"OUTSIDE\"],\"childNode\":null},{\"name\":\"ip\",\"value\":null,\"childNode\":[{\"name\":\"enable-arp-ignore\",\"value\":null,\"childNode\":null}]}]},{\"name\":\"eth1\",\"value\":null,\"childNode\":[{\"name\":\"description\",\"value\":[\"INSIDE\"],\"childNode\":null},{\"name\":\"address\",\"value\":[\"192.168.0.1/24\",\"123\"],\"childNode\":null}]}]},{\"name\":\"loopback\",\"value\":null,\"childNode\":[{\"name\":\"zx\",\"value\":null,\"childNode\":[{\"name\":\"ip\",\"value\":null,\"childNode\":null}]},{\"name\":\"lo\",\"value\":null,\"childNode\":[{\"name\":\"ip\",\"value\":null,\"childNode\":[{\"name\":\"source-validation\",\"value\":[\"strict\"],\"childNode\":null}]}]}]}]},{\"name\":\"system\",\"value\":null,\"childNode\":[{\"name\":\"name-server\",\"value\":[\"123\",\"456\",\"zxzx\",\"zxzx123\",\"zxzx123456\"],\"childNode\":null},{\"name\":\"login\",\"value\":null,\"childNode\":[{\"name\":\"user\",\"value\":null,\"childNode\":[{\"name\":\"vyos\",\"value\":null,\"childNode\":null}]}]},{\"name\":\"host-name\",\"value\":[\"1234\"],\"childNode\":null}]},{\"name\":\"service\",\"value\":null,\"childNode\":[{\"name\":\"ssh\",\"value\":null,\"childNode\":[{\"name\":\"disable-password-authentication\",\"value\":null,\"childNode\":null}]}]},{\"name\":\"nat\",\"value\":null,\"childNode\":[{\"name\":\"source\",\"value\":null,\"childNode\":[{\"name\":\"rule\",\"value\":null,\"childNode\":[{\"name\":\"100\",\"value\":null,\"childNode\":[{\"name\":\"source\",\"value\":null,\"childNode\":[{\"name\":\"address\",\"value\":[\"192.168.0.0/24\"],\"childNode\":null}]},{\"name\":\"outbound-interface\",\"value\":[\"eth0\"],\"childNode\":null},{\"name\":\"translation\",\"value\":null,\"childNode\":[{\"name\":\"address\",\"value\":[\"masquerade\"],\"childNode\":null}]}]}]}]}]}]}\n"

	json.Unmarshal([]byte(str), &tree)

	showJson, _ := Show(tree)
	fmt.Println(showJson)

}
