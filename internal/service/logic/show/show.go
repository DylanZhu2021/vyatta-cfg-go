package show

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/util/gutil"
	"vyatta-cfg-go/internal/service/logic/activecfg"
)

// Show 返回active config的json字符串
func Show(tree *activecfg.Node) (string, error) {

	var out bytes.Buffer
	var data map[string]interface{}

	data, _ = activecfg.TreeToMap(tree)

	gutil.Dump(data)
	mapstr, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	err = json.Indent(&out, mapstr, "", "\t")
	if err != nil {
		return "", err
	}

	showJson := out.String()

	//todo del
	fmt.Println(showJson)

	return showJson, err
}
