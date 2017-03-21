# 云中台应用部署API规划

## Namespace

一个用户对应一个namespace，用户注册时创建namespace。支持的操作如下：

- 查询 
- 创建
- 删除


## App

一个App包含一个Service和一个ReplicationController，对用户展现整个APP，不区分后者。但前端调用接口时需要区分。

- 创建/部署：调用创建Service和RC的接口，副本数为1
- 启动：将副本个数设置为1
- 停止：将副本个数设置为0
- 删除：调用Service和RC的删除接口


### Service

- 查询
- 创建
- 删除

### ReplicationController

- 查询
- 创建
- 删除

具体API接口在swagger文档中描述。
