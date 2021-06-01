package main

import (
	"client/services"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	c := &http.Client{}
	u := &url.URL{}
	u.Scheme = "http"
	u.Host = "localhost:8080"
	u.Path = "/check"
	urlStr := u.String()
	fmt.Println(urlStr)
	bodyStr :=   "{\"ai\":\"test-ai\",\"name\":\"陈耀兴\",\"idNum\":\"450721198908091810\"}"
	req, err := http.NewRequest("POST", urlStr, io.NopCloser(strings.NewReader(bodyStr)))
	req.Close = true
	req.Header.Set("Content-Type", "")
	req.Header.Set("appId", "test_appId")
	req.Header.Set("bizId", "test-bizId")
	rsp, err := c.Do(req)
	fmt.Println(rsp)

	return
	creds, err := credentials.NewClientTLSFromFile("keys/server.crt","chenyaoxing.com")
	if err != nil {
		log.Fatalln(err)
	}

	// 客户端
	// 1、不带自签证书带客户端
	//conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())

	// 2、带自签证书的客户端
	/*
	出现错误：
	2020/12/08 16:41:22 could not greet: rpc error: code = Unavailable desc = connection error: desc = "transport: authentication handshake failed: x509: certificate relies on legacy Common Name field, use SANs or temporarily enable Common Name matching with GODEBUG=x509ignoreCN=0"
	Process exiting with code: 1 signal: false
	解决办法：GODEBUG=x509ignoreCN=0 go run client.go
	GODEBUG=x509ignoreCN=0 加入到运行配置的环境变量中
	*/
	conn, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := services.NewProdServiceClient(conn)
	prodRes, err := client.GetProdStock(context.Background(), &services.ProdRequest{
		ProdId: 12,
		})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(prodRes.ProdStock)
}
