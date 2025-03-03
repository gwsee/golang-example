package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"golang-example/config"
	"golang-example/dao/model"
	"golang-example/dao/query"
	"golang-example/database"
	"gorm.io/datatypes"
	"gorm.io/gen"
)

var configFile = flag.String("f", "config/zero.yaml", "the config file")

func main() {
	var c config.MysqlConfig
	conf.MustLoad(*configFile, &c)
	gormDB, err := database.NewMysql(c.Mysql)
	if err != nil {
		return
	}
	_ = gormDB.AutoMigrate(&model.Kec{})
	query.SetDefault(gormDB)
	list, err := query.Kec.Where(gen.Cond(datatypes.JSONArrayQuery("area").Contains([]int{1, 2, 3}))...).Find()
	if err != nil {
		// SELECT * FROM `kec` WHERE JSON_CONTAINS (`area`, JSON_ARRAY((1,2,3))) AND `kec`.`deleted_at` = 0
		// Error 1241 (21000): Operand should contain 1 column(s)
		return
	}
	fmt.Println(len(list))
}
