package common

const (
	defaultDBConfigTemplate = `
---
server:
  name: db-demo                 # 项目名称，自定义
  version: 1.0.1                # 配置版本
dbconfigs:
   - name: 204                  # 数据源，自定义，唯一
     config:
      mode: true                # 数据源开关
      driver: mysql             # 数据库类型
      host: 192.168.1.204       # 数据库地址
      port: 3306                # 数据库端口
      username: web             # 数据库用户名
      password: 123456          # 数据库密码
      databasename: db_dingzuo  # 数据库名字
      lifetime: 10              # 连接失效时间
      max-open-num: 5     
      max-idle-num: 5       
`
   defaultConfigFilePath = "./resources/db-config.yaml"
   ConfigEnvName = "fast_config_path"
)
