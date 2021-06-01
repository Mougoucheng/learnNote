## 双向认证

*(过程中的common name都先填写成localhost，正常情况下会写成使用的域名)*

### 1.生成根证书
> * 根证书(Root certificate)是属于根证书颁发机构（CA）的公钥证书。
>   用以验证它所签发的证书（客户端、服务端）
> * openssl genrsa -out ca.key 2048
> * openssl req -new -x509 -days 3650 -key ca.key -out ca.pem

### 2.生成服务器端证书
> * openssl genrsa -out server.key 2048
> * openssl req -new -key server.key -out server.csr
> * openssl x509 -req -sha256 -CA ca.pem -CAkey ca.key -CAcreateserial -days 3650 -in server.csr -out server.pem

### 3.生成客户端证书
> * openssl ecparam -genkey -name secp384r1 -out client.key
> * openssl req -new -key client.key -out client.csr
> * openssl x509 -req -sha256 -CA ca.pem -CAkey ca.key -CAcreateserial -days 3650 -in client.csr -out client.pem



