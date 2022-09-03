# 本文件执行的校验
* syntax：syntax:expression: $VAR(@) == "auto" || \exec "/opt/vyatta/sbin/vyatta-qos-util.pl --rate $VAR(@)"
* 执行逻辑：直接判断需要校验的值是不是"auto",如果不是，就使用pl脚本进行校验
* pl脚本执行逻辑：
  * 直接执行pl脚本的命令
  * 获取返回值