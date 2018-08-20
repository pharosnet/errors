package main

import (
	"fmt"
	"github.com/pharosnet/errors"
)

func main() {
	e1 := errors.New("error 1")
	e2 := errors.With(e1, "error 2")
	fmt.Println(e2)
	// output
	//
	// error 2
	fmt.Println("======================")
	fmt.Printf("%+v", e2)
	// output
	//
	// error 2
	// 		[T] 2018-08-20 07:29:07.2631465 +0800 CST
	// 		[F] main.main
	// 		[P] E:/golang/workspace
	// 		[X] github.com/pharosnet/errors/example/stack.go:10
	// error 1
	// 		[T] 2018-08-20 07:29:07.2631465 +0800 CST
	//		[F] main.main
	//		[P] E:/golang/workspace
	//		[X] github.com/pharosnet/errors/example/stack.go:9

}
