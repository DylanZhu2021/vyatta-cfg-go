package pattern

import (
	"regexp"
	"strings"
	"vyatta-cfg-go/internal/service/logic/datavalidate/get"
)

type Pattern struct {
	RegExp string
	Ret    string
}

// GetPattern 获取到pattern校验方式的pattern结构体
// pattern $VAR(@) "^[[:alnum:]][-_[:alnum:]]*$"
// pattern $VAR(@) "^[^-]" ; "Match queue name cannot start with \"-\""
func GetPattern(path string) (*Pattern, error) {

	var pattern *Pattern
	pattern = new(Pattern)

	//获取syntax expression
	syntax, err := get.GetSyntaxExp(path)
	if err != nil {
		return nil, err
	}

	syntaxArr := strings.Split(syntax, ";")

	//由于有些pattern没有错误信息字段做如下判断
	if len(syntaxArr) == 1 {
		pattern.Ret = ""
	} else {
		//获取错误返回的信息
		//使用正则匹配：\\\",替换掉\"
		//先去掉首尾的"
		str1 := syntaxArr[1][2 : len(syntaxArr[1])-1]

		//正则匹配替换掉\"
		reg, err := regexp.Compile("\\\\\\\"")
		if err == nil {
			pattern.Ret = reg.ReplaceAllString(str1, "")
		}
	}

	//对pattern模式匹配字段获取到正则表达式
	reg, err := regexp.Compile("\".*\"")
	if err != nil {
		return nil, err
	}
	regexp := strings.Split(reg.FindAllString(syntaxArr[0], -1)[0], "\"")
	pattern.RegExp = regexp[1]

	return pattern, nil
}
