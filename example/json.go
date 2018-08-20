package main

import (
	"fmt"
	"github.com/pharosnet/errors"
	"io"
)

func main() {
	errors.Configure().SetFormatFunc(errors.JsonFormatFn)
	e1 := io.EOF
	e2 := errors.With(e1, "error 2")
	fmt.Printf("%+v", e2)

	// output:
	//
	//	{
	//		"msg": "error 2",
	//		"occurTime": "2018-08-20 08:23:01.3592428 +0800 CST",
	//		"stack": [
	//			{
	//				"fn": "main.main",
	//				"home": "E:/golang/workspace",
	//				"file": "github.com/pharosnet/errors/example/json.go",
	//				"line": 12
	//			}
	//		],
	//		"cause": {
	//			"msg": "EOF"
	//		}
	//	}
}
