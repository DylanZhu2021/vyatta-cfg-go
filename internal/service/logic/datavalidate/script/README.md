# 使用脚本校验
* > ./container/name/node.def:4:syntax:expression: exec "${vyos_libexec_dir}/validate-value --regex \'[-a-zA-Z0-9]+\'   --value \'$VAR(@)\'"; "Container name must be alphanumeric and can contain hyphens"
  ./container/name/node.tag/cap-add/node.def:11:syntax:expression: exec "${vyos_libexec_dir}/validate-value --regex \'(net-admin|net-bind-service|net-raw|setpcap|sys-admin|sys-time)\'   --value \'$VAR(@)\'"; "I
  nvalid value"
  ./container/name/node.tag/environment/node.def:4:syntax:expression: exec "${vyos_libexec_dir}/validate-value --regex \'[-_a-zA-Z0-9]+\'   --value \'$VAR(@)\'"; "Environment variable name must be alphanumeric
  and can contain hyphen and underscores"
  ./container/name/node.tag/memory/node.def:5:syntax:expression: exec "${vyos_libexec_dir}/validate-value  --exec \"${vyos_validators_dir}/numeric --range 0-16384\"  --value \'$VAR(@)\'"; "Container memory must
  be in range 0 to 16384 MB"
* 获取到syntax字段
* 将环境变量替换为vyo环境变量的具体值（由于goland获取不到vyos环境变量的原因，直接将环境变量读出来）
* （后面可以修改env.go中的GetEnvValue函数可实现直接从vyos获取环境变量）
* 使用正则匹配替换$VAR(@)
* 执行脚本，获取返回值
* 错误值返回Script结构体中的Ret字段