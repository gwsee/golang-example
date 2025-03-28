package main

import (
	//"code.byted.org/gopkg/lang/slices"
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	//"strings"
)

type SliceOfInt []int

func (c *SliceOfInt) GormValue(context.Context, *gorm.DB) clause.Expr {
	//by, _ := json.Marshal(*c)
	return clause.Expr{SQL: `1,2`}
	//return clause.Expr{SQL: strings.Join(slices.MapString(s, func(s string) string { return "'" + s + "'" }), ", ")}
}
