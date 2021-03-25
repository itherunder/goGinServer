package main

import (
	"ginServer/dao"
	"ginServer/models"
	"ginServer/routers"
)

func main() {
	// 创建数据库
	// CREATE DATABASE ginServer;
	// 连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	// 关闭数据库连接
	defer dao.Close()
	// 绑定
	dao.DB.AutoMigrate(&models.Todo{})
	// 注册路由
	r := routers.SetupRouter()
	r.Run()
}
