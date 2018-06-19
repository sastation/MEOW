# MEOW Proxy

当前版本：1.6 [CHANGELOG](CHANGELOG.md)
[![Build Status](https://travis-ci.org/netheril96/MEOW.png?branch=master)](https://travis-ci.org/netheril96/MEOW)

<pre>
       /\
   )  ( ')     MEOW 是 [COW](https://github.com/cyfdecyf/cow) 的一个派生版本
  (  /  )      MEOW 与 COW 最大的不同之处在于，COW 采用黑名单模式， 而 MEOW 采用白名单模式
   \(__)|      国内网站直接连接，其他的网站使用代理连接
</pre>

## 与netheril96版MEOW的差别
* IPv6可以选择是否走代理
* 去除内建的 ssh 代理、shadowsocks 代理、meow-ss 代理 （simple is the best)
* 所有配置文件默认统一为 meow 执行文件所在目录，所有配置文件后缀为 "*.conf"
* 在配置文件中增加 DProxy 参数，用于指定国内IP走的代理设置，用于某些特殊场合

## MEOW 可以用来
- 作为全局 HTTP 代理（支持 PAC），可以智能分流（直连国内网站、使用代理连接其他网站）
- 将 SOCKS5 等代理转换为 HTTP 代理，HTTP 代理能最大程度兼容各种软件（可以设置为程序代理）和设备（设置为系统全局代理）
- 架设在内网（或者公网），为其他设备提供智能分流代理
- 编译成一个无需任何依赖的可执行文件运行，支持各种平台（Win / Linux / OS X），甚至是树莓派（Linux ARM）

## 获取

- **从源码安装:** 安装 [Go](http://golang.org/doc/install)，然后 `go get github.com/sastation/MEOW`

## 配置

编辑 `./rc.conf`，例子：

    # 监听地址，设为0.0.0.0可以监听所有端口，共享给局域网使用
    listen = http://127.0.0.1:4411
    # 至少指定一个上级代理
    # SOCKS5 上级代理
    # proxy = socks5://127.0.0.1:1080
    # HTTP 上级代理
    # proxy = http://127.0.0.1:8087
    # HTTPS 上级代理
    # proxy = https://user:password@example.server.com:port
    # 直连所用代理
    # DProxy = https://127.0.0.1:8088

## 工作方式

当 MEOW 启动时会从配置文件加载直连列表和强制使用代理列表，详见下面两节。

当通过 MEOW 访问一个网站时，MEOW 会：

- 检查域名是否在直连列表中，如果在则直连
- 检查域名是否在强制使用代理列表中，如果在则通过代理连接
- **检查域名的 IP 是否为国内 IP**
    - 通过本地 DNS 解析域名，得到域名的 IP
    - 如果是国内 IP: 
        + 若有 DProxy 参数，则走 DProxy 指定的代理
        + 若无 DProxy 参数，则直连
    - 如果不是国内 IP，则走 "proxy" 参数指定的代理
    - 将域名加入临时的直连或者强制使用代理列表，下次可以不用 DNS 解析直接判断域名是否直连

## 直连列表

直接连接的域名列表保存在 `./direct.conf` 

匹配域名**按 . 分隔的后两部分**或者**整个域名**，例子：

-  `baidu.com` => `*.baidu.com`
-  `com.cn` => `*.com.cn`
-  `edu.cn` => `*.edu.cn`
-  `music.163.com` => `music.163.com`

一般是**确定**要直接连接的网站

## 强制使用代理列表

强制使用代理连接的域名列表保存在 `./proxy.conf`，语法格式与直连列表相同
（注意：匹配的是域名**按 . 分隔的后两部分**或者**整个域名**）

## 致谢

- @cyfdecyf - COW author
- @renzhn - Original MEOW author
- @netheril96 - MEOW v1.5 author
- Github - Github Student Pack
- https://www.pandafan.org/pac/index.html - Domain White List
- https://github.com/Leask/Flora_Pac - CN IP Data
- https://github.com/17mon/china_ip_list - CN IP Data
