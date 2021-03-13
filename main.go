package main

import (
	"ginProject/core"
	"ginProject/global"
	"ginProject/initialize"
)

func main() {
	//router := gin.Default()

	//http.ListenAndServe(":8080", router)

	//r.GET("/", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"Blog":"www.flysnow.org",
	//		"wechat":"flysnow_org",
	//	})
	//})
	//r.Run(":8080")
	//router := initialize.Routers()
	//http.ListenAndServe(":8080", router)
	// 初始化viper 读取配置文件 config.yaml
	global.GVA_VP = core.Viper()
	global.GVA_LOG = core.Zap()

	// 初始化数据库连接
	global.GVA_DB = initialize.Gorm()
	if global.GVA_DB != nil {
		global.GVA_LOG.Info("数据库连接成功，初始化数据表")
		initialize.MysqlTables(global.GVA_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}

	// 初始化zap 打印日志
	core.RunWindowsServer()
	//router := initialize.Routers()
	//s := &http.Server{
	//	Addr:           ":8080",
	//	Handler:        router,
	//	ReadTimeout:    10 * time.Second,
	//	WriteTimeout:   10 * time.Second,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//s.ListenAndServe()
}
