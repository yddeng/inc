# inc

Internal Network Channel 

做内网服务器到公网的映射访问，用于解决内网服务器没有公网IP或者无法进行端口映射的场景

## 介绍

* `master` 中心（公网）服务器，负责创建隧道，转发消息。
* `slave`  被代理的内网主机。
* `client` 命令行控制程序。

## 隧道

### 公有隧道

指 `slave-master-client` 的线路，由控制程序来选择消息去往的地址。可用于在目标 `slave` 上执行一些命令等。

### 专线隧道

通过公有隧道，在目标 `slave` 上创建一条专线，用于桥接 `slave` 上的服务，再将消息投递到服务。
例如 `slave` 上在 `localhost:5432` 运行了 `PostgreSQL`, 在 slave 上开启一条专线转发消息到 `PostgreSQL`,
`client` 机上的 `psql` 就能访问内网的 `PostgreSQL` 。

## 反向代理

将内网地址代理到公网地址上，可直接访问公网地址来访问内网

