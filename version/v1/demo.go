package v1

import (
	"EasyGinLite/utils"
	"context"
	"fmt"
)

func test() {

	//这里可以直接用全局调用的数据库，仅仅是个例子，或者可以提提
	result, err := utils.G.Redis1.Get(context.Background(), "AAAA").Result()
	if err != nil {
		return
	}
	fmt.Println(result)

	all, err := utils.G.MysqlDb.Model("AAAA").All()
	if err != nil {
		return
	}
	fmt.Println(all)

}
