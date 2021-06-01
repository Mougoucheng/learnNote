### CA证书自签
[参考网址](https://ningyu1.github.io/site/post/51-ssl-cert/)

步骤：
* 通过OpenSSL生成秘钥
```shell
  openssl genrsa -out server.key 1024
```
* 根据私钥生成证书申请文件csr
```shell
  openssl req -new -key server.key -out server.csr
```
* 这里根据命令行向导来进行信息输入：
    ![信息输入](https://ningyu1.github.io/site/img/ssl-cert/1.png)
  
<font color=red>*PS: Common name可以输入：\*.yourdomain.com, 这种方式生成通配符域名证书*</font>

* 使用私钥对证书申请进行签名从而生成证书
```shell
  openssl x509 -req -in server.csr -out server.crt -signkey server.key -days 3650
```

这样就生成了有效期为：10年的证书文件，对于自己内网服务使用足够了