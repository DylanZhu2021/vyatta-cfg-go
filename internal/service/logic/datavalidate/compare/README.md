# 执行比较大小的逻辑校验
* >./cluster/dead-interval/node.def:4:syntax:expression: ($VAR(@) >= 1 && $VAR(@) < 30000) ; "Value must be between 1 and 29,999"
  ./cluster/monitor-dead-interval/node.def:4:syntax:expression: ($VAR(@) >= 1 && $VAR(@) < 30000) ; "Value must be between 1 and 29,999"
  ./load-balancing/wan/interface-health/node.tag/failure-count/node.def:3:syntax:expression: $VAR(@) >= 0 && $VAR(@) <= 10; "failure count must be between 1-10"
  ./load-balancing/wan/interface-health/node.tag/success-count/node.def:3:syntax:expression: $VAR(@) >= 0 && $VAR(@) <= 10; "success count must be between 1 and 10"
  ./load-balancing/wan/interface-health/node.tag/test/node.tag/resp-time/node.def:3:syntax:expression: $VAR(@) >= 0 && $VAR(@) <= 30; "response must be between 1 and 30 seconds"
  ./load-balancing/wan/interface-health/node.tag/test/node.tag/ttl-limit/node.def:6:syntax:expression: $VAR(@) >=1 && $VAR(@) <= 254; "ttl limit hops must be between 1 and 254"
  ./load-balancing/wan/rule/node.def:4:syntax:expression: $VAR(@) > 0 && $VAR(@) <= 9999; "firewall rule number must be between 1 and 9999"
  ./load-balancing/wan/rule/node.tag/interface/node.tag/weight/node.def:3:syntax:expression: $VAR(@) >= 0 && $VAR(@) <= 255; "Interface weight must be between 1 and 255"
  ./system/config-management/commit-revisions/node.def:16:syntax:expression: $VAR(@) >= 0 && $VAR(@) <= 65535 ; \
* 获取到syntax中的expression字段
* 然后使用正则表达式对$VAR(@)进行替换
* 然后将字符串作为go语句执行，获取返回结果
* 结果错误返回Compare结构体中的Ret字段