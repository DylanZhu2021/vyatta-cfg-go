package get

import (
	"strings"
	"vyatta-cfg-go/internal/service/logic/tmplate"
	"vyatta-cfg-go/internal/service/logic/tools/gqueue"
)

// GetSyntaxExp 通过模板获取到syntax
// 传入的path就是路径，通过路径获取到模板
// path:eg:interfaces ethernet eth0 address
/*
对一下内容解析
 path :/opt/vyatta/share/vyatta-cfg/templates/protocols/bgp/local-as/
syntax:
		syntax:expression:
					exec "${vyos_libexec_dir}/validate-value
						--exec \"${vyos_validators_dir}/numeric --range 1-4294967294\"
						--value \'$VAR(@)\'"; "Invalid value"
最终返回：
		exec "${vyos_libexec_dir}/validate-value
						--exec \"${vyos_validators_dir}/numeric --range 1-4294967294\"
						--value \'$VAR(@)\'"; "Invalid value"
模板路径：/opt/vyatta/share/vyatta-cfg/templates/protocols/bgp/local-as/
解析执行脚本的命令：
1.env环境变量：从vyos中获取到env然后根据正则匹配将脚本命令中的环境变量替换，然后执行脚本呢变量
		TODO：正则匹配？？？？
2.\'$VAR(@)\'替换为传入的具体值！！也是用正则怕匹配
3.关于脚本执行命令的获取：可以通过对"（引号，英文）及进行分割，然后去掉首位，然后再用引号粘连在一起！！
*/
func GetSyntaxExp(path string) (string, error) {

	var queue gqueue.QueueItf
	var validStr string

	queue = gqueue.New()
	queue.GetPathQue(path)

	//得到模板节点
	NodeTmp, err := tmplate.GetTmpNode(queue)
	if err != nil {
		return "", err // errors.New("get template failed！")
	}

	//得到syntax字段
	syntax := NodeTmp.Syntax
	strSplit := strings.Split(syntax, ":")

	if len(strSplit) == 2 {
		validStr = strSplit[len(strSplit)-1]
	} else {
		for k, v := range strSplit {
			if k == 0 {
				validStr = v
				continue
			}
			validStr = validStr + ":" + v
		}
	}

	return validStr, nil
}
