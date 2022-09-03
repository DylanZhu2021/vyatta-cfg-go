package compare

import (
	"strings"
	"vyatta-cfg-go/internal/service/logic/datavalidate/get"
)

type Compare struct {
	ComExp string
	Ret    string
}

// GetCompare 根据路径获取到Compare结构体
// /opt/vyatta/share/vyatta-cfg/templates/cluster/monitor-dead-interval/
//   ($VAR(@) >= 1 && $VAR(@) < 30000) ; "Data must be between 1 and 29,999"
func GetCompare(path string) (*Compare, error) {

	var compare *Compare

	compare = new(Compare)

	syntax, err := get.GetSyntaxExp(path)
	if err != nil {
		return nil, err
	}
	syntaxArr := strings.Split(syntax, ";")

	//获取到Ret
	compare.Ret = strings.Split(syntaxArr[1], "\"")[1]

	//获取到ComExp(比较表达式)
	//去掉首尾两个字符
	str1 := strings.Trim(syntaxArr[0], " (")
	str2 := strings.Trim(str1, ") ")
	compare.ComExp = str2

	return compare, nil
}
