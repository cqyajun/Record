package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

//写数据
var DB_Write *xorm.Engine

func init() {
	fmt.Println("db ==> init")
	//数据库链接
	var err error
	DB_Write, err = xorm.NewEngine("mysql", "root:root@/xorm?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	}
	//连接测试
	if err := DB_Write.Ping(); err != nil {
		fmt.Println(err)
		return
	}
	//日志打印SQL
	DB_Write.ShowSQL(true)
	//设置连接池的空闲数大小
	DB_Write.SetMaxIdleConns(5)
	//设置最大打开连接数
	DB_Write.SetMaxOpenConns(5)

	//名称映射规则主要负责结构体名称到表名和结构体field到表字段的名称映射
	DB_Write.SetTableMapper(core.SnakeMapper{})
}
