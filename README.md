# inc

Internal Network Channel 

做内网服务器到公网的映射访问，用于解决内网服务器没有公网IP或者无法进行端口访问的场景


## 使用 

```
string internal_ip    = 1; // 内网ip
int32  internal_port  = 2; // 内网代理端口
int32  external_port  = 3; // 外网端口
string description    = 4; // 描述
```

### 代理内网 ssh

1. 启动 inc-master:

"./inc-master -h x.x.x.x -p 9852"

2. 启动 inc-slave(并注册代理):

`./inc-slacve -a x.x.x.x:9852 -r "127.0.0.1 22 2201 ssh" `

3. 使用ssh客户端连接（例用户名：test）

`ssh -P 2201 test@x.x.x.x`


