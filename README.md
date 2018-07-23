# gosparrow
麻雀虽小，五脏俱全。基于http的golang空项目模板。

# 模块
* DEPLOY: 配置文件，启动脚本存放
* base: 基础模块，包含通用函数库，http_server库，日志库
* config: 配置文件初始化模块
* db: 数据库操作
* handle: 业务逻辑处理
* proxy: 第三方调用

# 其他
* 支持make打包
* http_server模块使用gin
* log模块使用seelog
* config使用configor，yaml格式
* 支持http请求日志统一打印
* 支持第三方调用请求和响应日志统一打印
* 支持错误日志预警上报（待实现）

# 依赖
* [github.com/cihub/seelog](http://github.com/cihub/seelog)
* [github.com/jinzhu/configor](http://github.com/jinzhu/configor)
* [github.com/gin-gonic/gin](http://github.com/gin-gonic/gin)
