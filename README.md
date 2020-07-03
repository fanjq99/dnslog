# dnslog
dnslog  reverse vul-verify  反连平台 漏洞验证 

## dnslog 使用环境准备
 1.搭建并使用 DNSLog,需要两个域名,一个外网地址.一个ns域名,一个dnslog使用的域名,并将dnslog域名的ns记录使用ns域名。两个域名同时指向外网地址.
 
 2.修改配置文件 
 ``` yaml
 # dns域名
 dns_domain: xxx.xxx
 
 #服务器对外ip
 server_ip: 127.0.0.1
 
 #key save time,单位秒
 save_time: 300
 
 api_addr: "127.0.0.1:8888"
 
 # redis 配置
 redis:
   addr: 127.0.0.1:6379 #ip:port
   password: xxx
   database: 0
 ```
 
 ## 运行
 在 cmd目录下面
 
 `go bulid`
 
 `cmd -yml fixture/dev.yml`
