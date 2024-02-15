# 通过微信公众号发消息添加到trilium笔记

### 安装使用
创建conf.toml，添加如下配置
```
[app]
    runmode="test" # gin mode取值： test、release、debug
[wx]
    wx_token="xxx" # 微信公众号后台的令牌配置
    self_open_ids="xxxx,xxxx" # 发送者微信openid，只对指定openid生效，多个用英文逗号分隔
    note_prefix="note:" # 添加笔记消息前缀，用于标识
[trilium]
    eapi_token="xxx" # trilium eapi 令牌
    host="http://trilium_server:8080" # trilium server
    parent_note_id="xxxx" # 笔记节点id，创建的笔记在这个笔记目录下
```
使用docker启动
```
docker run --restart=always -d -p 1234:1234 -v /yourpath/conf.toml:/conf.toml --name wx-trilium vbmf/wx-trilium
```

- 通过nginx代理80端口，将来自微信的消息代理到docker服务上

- 添加微信公众号服务器配置，注意令牌配置和配置文件里的保持一致

- 获取openid
公众号里发送消息：
```
openid
```

- 记录笔记
目前之支持简单的文本消息，公众号里发送消息：
```
note:
标题
正文
```

