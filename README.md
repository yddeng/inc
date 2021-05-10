# intun

Internal Network Tunnel 

做内网服务器到公网的映射访问，用于解决内网服务器没有公网IP或者无法进行端口映射的场景

## intun-root

中心服务器，负责注册、统计 leaf 和转发消息。部署在可访问的地址上（公网）。

## intun-leaf

需要桥接的地址，真正执行命令的节点。一般部署在不能访问的地址（内网）

## intun-cli

控制程序。
