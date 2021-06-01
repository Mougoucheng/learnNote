package services

import (
	"context"
	"fmt"
)

type ProdService struct {

}

func (th *ProdService) GetProdStock(ctx context.Context, request *ProdRequest) (*ProdResponse, error) {
	return &ProdResponse{
		ProdStock: 20,
	}, nil
}

func (th *ProdService) mustEmbedUnimplementedProdServiceServer() {
	// TODO：干嘛用的？
	//panic("implement me")
	fmt.Println("hello??")
}