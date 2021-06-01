package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net/http"
	"server/services"

	//"gorpc.jtthink.com/services"
)

func main() {
	// openssl自签证书
	creds, err := credentials.NewServerTLSFromFile("keys/server.crt", "keys/server.key")
	if err != nil {
		log.Fatalln(err)
	}


	// rpc服务创建
	// 1、不带证书的rpc服务
	//rpcServer := grpc.NewServer()

	// 2、带自签证书选项的rpc服务
	rpcServer := grpc.NewServer(grpc.Creds(creds))

	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))

	//lis, err := net.Listen("tcp", "localhost:8081")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//rpcServer.Serve(lis)

	// 3、提供http服务
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.Proto)
		fmt.Println(request.Header)
		fmt.Println(request)
		rpcServer.ServeHTTP(writer, request)
		})
	httpServer := http.Server{
		Addr: "localhost:8081",
		Handler: mux,
	}
	// httpServer.ListenAndServe()  // 报错：gRPC requires HTTP/2
	httpServer.ListenAndServeTLS("keys/server.crt", "keys/server.key")
}
