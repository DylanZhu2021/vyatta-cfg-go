package tmplate

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

// NodeTmp 定义节点结构体(node.def)
type NodeTmp struct {
	Multi     bool
	Tag       bool
	Type      string
	Help      string
	Comp_help string
	Val_help  string
	Allowed   string
	Syntax    string
	Commit    string
	Priority  string
	DefValue  string
	Begin     string
	End       string
	Create    string
	Update    string
	Delete    string
}

/*
获取节点模板的值，存放于一个节点的结构体中，
这个结构体就是模板节点的个属性的值，用于后面的节点校验等操作
*/

// ReadTemplate 读取template文件node.def值,返回string（原文件的内容）
func ReadTemplate(path string) string {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("read file fail", err)
		return ""
	}
	defer f.Close()

	fd, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("read to fd fail", err)
		return ""
	}

	return string(fd)
}

// ReadTmpToStruct 把文件读取到的string转成node结构体
// 传入文件的路径，返回模板节点结构体
func ReadTmpToStruct(filename string) *NodeTmp {
	var node NodeTmp

	node.Tag = false
	node.Multi = false

	fieldArr := ReadLineFile(filename)
	for k, v := range fieldArr {
		if k == "type" {
			node.Type = v
		}
		if k == "help" {
			node.Help = v
		}
		if k == "comp_help" {
			node.Comp_help = v
		}
		if k == "val_help" {
			node.Val_help = v
		}
		if k == "allowed" {
			node.Allowed = v
		}
		if k == "syntax" {
			node.Syntax = v
		}
		if k == "commit" {
			node.Commit = v
		}
		if k == "priority" {
			node.Priority = v
		}
		if k == "default" {
			node.DefValue = v
		}
		if k == "begin" {
			node.Begin = v
		}
		if k == "end" {
			node.End = v
		}
		if k == "create" {
			node.Create = v
		}
		if k == "update" {
			node.Update = v
		}
		if k == "delete" {
			node.Delete = v
		}
		if k == "multi" {
			node.Multi = true
		}
		if k == "tag" {
			node.Tag = true
		}
	}

	return &node
}

// ReadLineFile 按行读取文件，将一行数据作为一个map，K是字段名字，V是字段值
func ReadLineFile(fileName string) map[string]string {
	var fieldArr map[string]string
	var lastName string
	var commitArr []string

	fieldArr = make(map[string]string)

	if file, err := os.Open(fileName); err != nil {
		panic(err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			//匹配正则表达式，判断是不是拥有字段名字和字段值的一行
			reg, _ := regexp.Compile("\\w:[\\w]?")

			if reg.MatchString(scanner.Text()) {
				//获取字段名字
				lastName = GetFieldName(scanner.Text())
				fieldArr[lastName] = GetFiledValue(scanner.Text())
			} else {
				//如果没有字段名字和字段值，说明是上一行字段值的一部分，做字符串拼接操作
				fieldArr[lastName] = strings.Trim(fieldArr[lastName], "\\")
				text := strings.Trim(scanner.Text(), "\t")
				fieldArr[lastName] = fieldArr[lastName] + text //+ " "
				if lastName == "commit" {
					//如果字段名为commit的话，将字段值放入字符串数组
					if len(commitArr) != 0 && commitArr[len(commitArr)-1] == fieldArr[lastName] {
						continue
					}
					commitArr = append(commitArr, fieldArr[lastName])
				}

			}

		}
		//单独处理commit字段
		commitStr := ""
		for _, v := range commitArr {
			if commitStr == "" {
				commitStr += v
			} else {
				commitStr += "+" + v
			}
		}
		fieldArr["commit"] = commitStr
	}
	return fieldArr
}

// GetFieldName 返回文件一行中的字段名字
func GetFieldName(filelinestr string) string {
	countSplit := strings.Split(filelinestr, ":")
	return countSplit[0]
}

// GetFiledValue 返回文件一行中的字段值
func GetFiledValue(filelinestr string) string {
	countSplit := strings.Split(filelinestr, ":")
	if len(countSplit) <= 2 {
		if countSplit[0] == "tag" || countSplit[0] == "mulit" {
			countSplit[1] = "true"
		}
	} else {
		if countSplit[0] == "tag" || countSplit[0] == "mulit" {
			countSplit[1] = "true"
		}
		for i := 0; i <= len(countSplit)-2; i++ {
			if i == 0 {
				countSplit[1] = countSplit[i+1]
				continue
			}
			countSplit[1] = countSplit[1] + ":" + countSplit[i+1]
		}

	}
	return countSplit[1]
}
