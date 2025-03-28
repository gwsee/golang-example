package cdatatypes

import (
	"gorm.io/gen/field"
	"gorm.io/gorm/clause"
)

type ExprCond struct {
	clause.Expr
	field.String
}

func Cond(expr clause.Expr) *ExprCond {
	return &ExprCond{
		Expr:   expr,
		String: field.String{},
	}
}

func (c *ExprCond) BeCond() interface{} { return c.Expr }

func (c *ExprCond) CondError() error { return nil }
