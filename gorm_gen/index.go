package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"golang-example/config"
	"golang-example/dao/model"
	"golang-example/dao/query"
	"golang-example/database"
	"golang-example/gorm_gen/cdatatypes"
	"gorm.io/gorm"
)

var configFile = flag.String("f", "./config/zero.yaml", "the config file")

func main() {
	var c config.MysqlConfig
	conf.MustLoad(*configFile, &c)
	gormDB, err := database.NewMysql(c.Mysql)
	if err != nil {
		return
	}
	_ = gormDB.AutoMigrate(&model.Kec{})
	query.SetDefault(gormDB)
	var a []int
	a = append(a, 1, 2, 3)
	list, err := query.Area.Where(cdatatypes.Cond(gorm.Expr("JSON_CONTAINS (`AREA_INT`, json_array(?))", a))).Find()
	if err != nil {
		// SELECT * FROM `kec` WHERE JSON_CONTAINS (`area`, JSON_ARRAY((1,2,3))) AND `kec`.`deleted_at` = 0
		// Error 1241 (21000): Operand should contain 1 column(s)
		fmt.Println(err.Error())
	}
	list, err = query.Area.Where(cdatatypes.Cond(gorm.Expr("JSON_CONTAINS (`AREA_INT`, json_array(?))", 1))).Find()
	if err != nil {
		// SELECT * FROM `kec` WHERE JSON_CONTAINS (`area`, JSON_ARRAY((1,2,3))) AND `kec`.`deleted_at` = 0
		// Error 1241 (21000): Operand should contain 1 column(s)
		fmt.Println(err.Error())
	}
	fmt.Println(len(list))
}
