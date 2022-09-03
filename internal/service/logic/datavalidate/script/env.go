package script

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

// GetEnvValueFromFile 由于获取不到vyos的环境变量，直接将环境变量读取出来
func GetEnvValueFromFile() map[string]string {

	var envMap map[string]string /*创建集合 */
	envMap = make(map[string]string)

	//实现从vyos中读取环境变量
	fileName := "/home/vyos/env.txt"
	file, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Open file error!", err)
		return nil
	}
	defer file.Close()

	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)

		envArr := strings.Split(line, "=")
		//**************这里最后会出现一个空，奇怪！！！
		if len(envArr) == 1 {
			break
		}
		envMap[envArr[0]] = envArr[1]

		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println("Read file error!", err)
				return nil
			}
		}
	}
	return envMap
}

// GetEnvValue 跟据env的key获取到value
func GetEnvValue(name string) (string, error) {

	//由于无法直接获取env，这里使用从文件读取之后放入map，再从map遍历
	envMap := GetEnvValueFromFile()
	for k, v := range envMap {
		if k == name {
			return v, nil
		}
	}
	return "", errors.New("not find env value！")

	//后面放入vyos中直接用下代码
	//return os.Getenv(name),nil

}
