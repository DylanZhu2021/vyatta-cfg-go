# 使用备选判断逻辑
* > ./load-balancing/wan/interface-health/node.tag/test/node.tag/type/node.def:5:syntax:expression: $VAR(@) in "ping", "ttl", "user-defined";
  ./load-balancing/wan/rule/node.tag/limit/period/node.def:4:syntax:expression: $VAR(@) in "second", "minute", "hour"; "period must be second or minute or hour"
  ./load-balancing/wan/rule/node.tag/limit/threshold/node.def:4:syntax:expression: $VAR(@) in "below", "above"; "limit can apply above or below rate"
  ./traffic-policy/priority-queue/node.tag/class/node.tag/queue-type/node.def:3:syntax:expression: $VAR(@) in "fq-codel", "fair-queue", "priority", "drop-tail", "random-detect";\
  ./traffic-policy/priority-queue/node.tag/default/queue-type/node.def:3:syntax:expression: $VAR(@) in "fq-codel", "fair-queue", "priority", "drop-tail", "random-detect";\
  ./traffic-policy/round-robin/node.tag/class/node.tag/queue-type/node.def:3:syntax:expression: $VAR(@) in "fq-codel", "fair-queue", "priority", "drop-tail"; \
  ./traffic-policy/round-robin/node.tag/default/queue-type/node.def:3:syntax:expression: $VAR(@) in "fq-codel", "fair-queue", "priority", "drop-tail";\
  ./traffic-policy/shaper/node.tag/class/node.tag/queue-type/node.def:3:syntax:expression: $VAR(@) in "fq-codel", "fair-queue", "priority", "drop-tail", "random-detect";\
  ./traffic-policy/shaper/node.tag/default/queue-type/node.def:3:syntax:expression: $VAR(@) in "fq-codel", "fair-queue", "priority", "drop-tail", "random-detect";\
* 使用正则表达式获取到syntax字段中的备选值
* 然后判断需要设置的值是否与备选值一致
* 错误值返回SelectIn结构体中的Ret字段