### 微服务Go项目结构

* adaptor 负责处理与外界数据(其他服务调用、数据库、缓存等)的交互
* config 表示项目中的配置
* components 表示业务组件，我们主要的业务逻辑都在这里
* handler view层，处理参数绑定，返回值的封装逻辑
* service 主要的业务逻辑层


服务使用rpcx作为rpc通讯框架

