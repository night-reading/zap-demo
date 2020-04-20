# zap-demo


## zap+日志分级分文件+按时间切割日志整合

### 实现功能
    info debug 级别的日志输出到 /path/log/demo.log
    warn error .... 级别的日志输出到 /path/log/demo_error.log
    日志自动按小时分割 最多保留7天的日志
    
    以及程序启动时生成 pid 文件，退出是销毁 pid 文件
    
    
### 依赖的第三方包github地址
    https://github.com/uber-go/zap
    https://github.com/lestrrat-go/file-rotatelogs
    
    zap提供日志分级，打点key-value参数等功能，而file-foratelogs则提供日志自动按时间分割的功能。
