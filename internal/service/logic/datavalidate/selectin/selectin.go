package selectin

import (
	"regexp"
	"strings"
	"vyatta-cfg-go/internal/service/logic/datavalidate/get"
)

type SelectIn struct {
	Sel []string
	Ret string
}

func GetSelect(path string) (*SelectIn, error) {

	var selectin *SelectIn
	var sel []string

	selectin = new(SelectIn)

	syntax, err := get.GetSyntaxExp(path)
	if err != nil {
		return nil, err
	}

	syntaxArr := strings.Split(syntax, ";")

	//获取到Ret
	selectin.Ret = strings.Split(syntaxArr[1], "\"")[1]

	//获取到sel字符串数组
	reg, _ := regexp.Compile("\".*\"")
	strArr := reg.FindAllString(syntaxArr[0], -1)
	selStr := strings.Split(strArr[0], ",")

	for i := 0; i < len(selStr); i++ {
		str1 := strings.Split(selStr[i], "\"")[1]
		sel = append(sel, str1)
	}

	selectin.Sel = sel
	return selectin, nil
}
