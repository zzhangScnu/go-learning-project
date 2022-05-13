package main

import (
	"context"
	"fmt"
	"github.com/go-learning-project/gen/config"
	"github.com/go-learning-project/gen/service"
	"log"
	"testing"
)

func TestQuery(t *testing.T) {
	config.InitMysql()
	ctx := context.Background()
	res, err := service.NewUserService().QueryAll(ctx)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range res {
		//log.Println(v)
		log.Printf("%+v\n", v) // %+v会比上面的打印多出字段名，更方便对比查看
	}
}

func TestDIYMethod(t *testing.T) {
	config.InitMysql()
	res, err := service.NewUserService().QueryByName("admin", context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res["name"].(string))
}
