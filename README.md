# aliyun_driver_gin
基于gin+微服务的临时文件存储服务（后端），可自由设置文件过期时间（基于redis键过期时间通知），使用个人阿里云盘当云文件服务器（免费可以搞到好几t，比对象存储便宜多了），控制中心实现了简单的服务发现和简单的负载均衡。

# aliyun_services
后台服务，负责与阿里云盘交互，使用的的阿里云盘SDK是https://github.com/jakeslee/aliyundrive
需要在配置文件中设置好控制中心的ip+host，后续阿里云的reflashtoken会自动到控制中心拉取

# Control_center
控制中心，负责服务发现，心跳检测，服务变动通知
需要在本地配置好配置阿里云盘的reflashtoken,和redis的ip+host，还有要保存文件到哪个目录的目录名

# aliyun_web、
web端，从控制中心拉取后台服务的ip+host，负责负载均衡，返回客户端服务端ip+host
需要在配置文件中设置好控制中心的ip+host



# 感谢
https://github.com/jakeslee/aliyundrive
