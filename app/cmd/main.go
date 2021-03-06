package main

import (
	"github.com/revel/revel"
	"github.com/revel/cmd/harness"
	"fmt"
)

func main() {
	// go run main.go
	// 生成routes.go, main.go
	revel.Init("", "github.com/leanote/leanote", "")
	_, err := harness.Build() // ok, err
	if err != nil {
		panic(err)
	}
	fmt.Println("Ok")
//	panicOnError(reverr, "Failed to build")
}