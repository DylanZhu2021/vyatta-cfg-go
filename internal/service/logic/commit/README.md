# Commit
* commit之前的操作:  **操作对象是cmds**
  * 1.将cmds转化为操作树：command函数
  * 2.操作树过滤出节点放入优先级队列(如果操作树中的节点有不符合要求的直接设置失败)
  * 3.优先级队列的节点合并到active config

* commit将优先级队列中的节点修改到ACfg里面的操作
  * 根据优先级队列中的路径，添加到ACfg中
  * 实现逻辑参考之前的set和del操作