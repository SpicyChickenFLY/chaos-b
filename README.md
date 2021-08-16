# auto-mysql(MySQL Automatic Installation)

This Project is use to automate MySQL Initializing Procedure by GOLANG

# New Feature
* 增加了用户自定义服务器实例格式化参数输入
* 提供格式化参数的解析
* 支持多服务器多实例的自动初始化/启停/建立主从
* 支持自定义修改配置参数

## Preparation
```bash
systemctl stop firewalld.service # 停止firewall
systemctl disable firewalld.service # 禁止firewall开机启动

vim /etc/selinux/config # 关闭SELINUX
# 将SELINUX=enforcing改为SELINUX=disabled

iptables -F # 清除和关闭iptables
iptables-save # 保存
```

## Usage
``` bash
# SRC_SQL_FILE_FOLDER is where you put your mysql
# DST_FILE_FOLDER is where you want to install the mysql
# SRC_CNF_FILE_FOLDER is your configure file 
# root@host:~/$ automysql [ -m [single/multi] -s <SRC_SQL_FILE>, -d <DST_SQL_PATH>, -c <SRC_CNF_FILE> ] # WARNING!!! NOW DEPRECATED

go run main.go -m standard -p 123456  -i "root:123456@192.168.1.15:22#3306|3307;root:123456@192.168.1.14:22"
```

Default value as follow

```bash
SRC_SQL_FILE = "./src/mysql.tar.gz"
DST_SQL_PATH = "/usr/local/mysql"
SRC_CNF_FILE = "./src/my.cnf"
```

## Warning

1. Please **don't** assign a password for "mysql" user (e.g. **passwd mysql** )
2. Make sure the installing path can be accessed by general users(*will be automated in future*)
3. do not touch any related file while running this script



## Program Interface

User Interface

![](static/imgs/1.png)

Running Log

![](static/imgs/2.png)




## Further Work

1. 封装成完整的包
2. 搭建可视化界面（Gin）提供方便的参数配置与自动安装




## FAQ

1.有啥用？

省点打cd,chown,mv指令的时间

2.有什么潜在的问题？

如果你把mysql装在奇奇怪怪的路径上（不保证当前用户可达的那种），初始化不一定能成功
