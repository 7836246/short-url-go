# 短链的源码Go版本

##  项目实现
    - Java 改 Go
    - Mysql 改 Sqlite
    - 前端页面改 Nuxt3 
    - UI Acro Design

### 接口地址
```
    路由前缀: api
    get参数: url
    PublicGroup.GET("/url/generate", handler.GenerateShortUrl)
	PublicGroup.GET("/url/getLong", handler.GetLongUrl)
```
### 图片
![图片](/img/img.png "图片")
## 使用方法
```
    前段只需要 cd web
    执行 nuxt generate 可能需要 pnpm i 安装一下在编译
    后端只需要 go run main.go 端口在 conf/config.yaml
```
## 优化点(难点、亮点)
 - 生成短链只需要访问一次数据库。而不是传统的先查询，在判断插入，而是直接插入，用唯一索引来判断是否hash冲突
 - 利用LRUCache，将最近生成的几千个kv放进map中，一段时间内，同一个长url会生成相同的短url
 - hash冲突后，给hash冲突值 加一个随机url，降低冲突概率
 - 选择比较优秀的murmur64 hash算法
 - get获取常链的时候，利用LRU识别热点数据，直接从map中读取，防止打挂数据库

参考: https://github.com/wenbochang888/short-url