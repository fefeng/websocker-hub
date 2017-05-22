## websocket-hub 模块使用

#### 模块说明

websocket-hub是一个在微服务架构中，用于集中管理多个web后端与web前端之间的websocket通讯,避免了多个web后端和一个前端之间需要建立多个websocker通道，避免了前端的资源开销。



#### 结构图

端口号默认为80

#### API

```sh
curl -X POST http://192.168.14.164:8887/v1/socket/notice -d '{"module": "uflow","type": "build","action": "update","content": {},"key": "app_p1"}' 
```

数据格式

```json
{
    "module": "uflow",  // 通知模块来源，便于前端识别是那个模块的更新。
    "key": "app_p1",    // 通知资源路径，这里表示用户组"app"下的流水线"p1"。
    "type": "build",    // 通知资源类型，通知的是构建信息。
    "action": "update", // 通知更新的动作,create,delete,update。
    "content": {}      // 更新内容
}
```
