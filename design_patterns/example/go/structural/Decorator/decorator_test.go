package Decorator

import (
	"chenyx/design_patterns/example/go/structural/Decorator/bad"
	"chenyx/design_patterns/example/go/structural/Decorator/good"
	"log"
	"testing"
)

func TestDecorator_bad(t *testing.T) {
	// 编译时装配
	fs := bad.NewFileStream()
	fs1 := bad.NewCryptoFileStream(*fs)       // 加密
	fs2 := bad.NewBufferFileStream(*fs)       // 缓存
	fs3 := bad.NewCryptoBufferFileStream(*fs) // 加密并缓存
	log.Println(fs1, fs2, fs3)
}

func TestDecorator_good(t *testing.T) {
	// 运行时装配
	fs := good.NewFileStream()
	s1 := good.NewCryptoStream(fs) // 加密
	s2 := good.NewBufferStream(fs) // 缓存
	s3 := good.NewBufferStream(s1) // 加密并缓存
	log.Println(s1, s2, s3)
}
