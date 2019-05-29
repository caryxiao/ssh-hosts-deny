# 监控日志变化

看到自己的服务器经常被别人扫描端口, ssh登录尝试, 所以准备写个程序直接监控系统日志, 如果出现使用root账户登录多次都密码错误, 直接封掉IP, 错误次数可以自己设置, 默认5次

```
  -cnt int 
        ssh login failed count (default 5) SSH登录出现错误的次数
  -df string
        hosts.deny file path (default "/etc/hosts.deny") hosts.deny 文件位置
  -log-level int
        log level, default:5, trace:6, debug:5, info:4, warning:3, error:2, fatal:1, panic:0 (default 5)
        日志等级
  -log-path string
        log save path, default terminal output
        日志保存的位置
  -sf string
        Please specify a file you need to monitor
        secure日志文件位置, 这里没有指定默认的
  -v    print version
```



