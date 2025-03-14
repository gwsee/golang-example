// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"golang-example/dao/model"
)

func newKec(db *gorm.DB, opts ...gen.DOOption) kec {
	_kec := kec{}

	_kec.kecDo.UseDB(db, opts...)
	_kec.kecDo.UseModel(&model.Kec{})

	tableName := _kec.kecDo.TableName()
	_kec.ALL = field.NewAsterisk(tableName)
	_kec.ID = field.NewInt(tableName, "id")
	_kec.KecID = field.NewString(tableName, "kec_id")
	_kec.Name = field.NewString(tableName, "name")
	_kec.Cover = field.NewString(tableName, "cover")
	_kec.CateIds = field.NewString(tableName, "cate_ids")
	_kec.CateShow = field.NewInt8(tableName, "cate_show")
	_kec.Area = field.NewString(tableName, "area")
	_kec.SellerID = field.NewInt(tableName, "seller_id")
	_kec.Desc = field.NewString(tableName, "desc")
	_kec.Lecturer = field.NewString(tableName, "lecturer")
	_kec.LecturerDesc = field.NewString(tableName, "lecturer_desc")
	_kec.Num = field.NewInt(tableName, "num")
	_kec.RequireNum = field.NewInt(tableName, "require_num")
	_kec.NotRequireNum = field.NewInt(tableName, "not_require_num")
	_kec.RecNum = field.NewFloat64(tableName, "rec_num")
	_kec.RequiredXNum = field.NewFloat64(tableName, "required_x_num")
	_kec.NotRequiredXNum = field.NewFloat64(tableName, "not_required_x_num")
	_kec.PubTime = field.NewInt(tableName, "pub_time")
	_kec.MinPrice = field.NewFloat64(tableName, "min_price")
	_kec.MaxPrice = field.NewFloat64(tableName, "max_price")
	_kec.RecPrice = field.NewFloat64(tableName, "rec_price")
	_kec.Level = field.NewInt(tableName, "level")
	_kec.Comment = field.NewInt8(tableName, "comment")
	_kec.Status = field.NewInt8(tableName, "status")
	_kec.IsPackage = field.NewInt8(tableName, "is_package")
	_kec.CoverFrom = field.NewInt8(tableName, "cover_from")
	_kec.Duration = field.NewInt(tableName, "duration")
	_kec.CreatedAt = field.NewTime(tableName, "created_at")
	_kec.UpdatedAt = field.NewTime(tableName, "updated_at")
	_kec.DeletedAt = field.NewField(tableName, "deleted_at")
	_kec.KecCode = field.NewString(tableName, "kec_code")

	_kec.fillFieldMap()

	return _kec
}

type kec struct {
	kecDo

	ALL             field.Asterisk
	ID              field.Int
	KecID           field.String
	Name            field.String
	Cover           field.String
	CateIds         field.String
	CateShow        field.Int8   // 是否显示
	Area            field.String // 数据地址
	SellerID        field.Int    // 课程归属商家的ID
	Desc            field.String
	Lecturer        field.String
	LecturerDesc    field.String
	Num             field.Int     // 课程数量
	RequireNum      field.Int     // 必修课程数量
	NotRequireNum   field.Int     // 选修课程数量
	RecNum          field.Float64 // 推荐学时数量 在课程包的时候等于后者之和
	RequiredXNum    field.Float64 // 必修学时
	NotRequiredXNum field.Float64 // 选修学时
	PubTime         field.Int
	MinPrice        field.Float64
	MaxPrice        field.Float64
	RecPrice        field.Float64 // 推荐价格
	Level           field.Int
	Comment         field.Int8
	Status          field.Int8 // 0未发布1已发布
	IsPackage       field.Int8 // 0课程1课程包
	CoverFrom       field.Int8 // 封面来自哪里0自定义1系统
	Duration        field.Int  // 课程时长
	CreatedAt       field.Time
	UpdatedAt       field.Time
	DeletedAt       field.Field  // 默认未删除
	KecCode         field.String // 课程编号--唯一

	fieldMap map[string]field.Expr
}

func (k kec) Table(newTableName string) *kec {
	k.kecDo.UseTable(newTableName)
	return k.updateTableName(newTableName)
}

func (k kec) As(alias string) *kec {
	k.kecDo.DO = *(k.kecDo.As(alias).(*gen.DO))
	return k.updateTableName(alias)
}

func (k *kec) updateTableName(table string) *kec {
	k.ALL = field.NewAsterisk(table)
	k.ID = field.NewInt(table, "id")
	k.KecID = field.NewString(table, "kec_id")
	k.Name = field.NewString(table, "name")
	k.Cover = field.NewString(table, "cover")
	k.CateIds = field.NewString(table, "cate_ids")
	k.CateShow = field.NewInt8(table, "cate_show")
	k.Area = field.NewString(table, "area")
	k.SellerID = field.NewInt(table, "seller_id")
	k.Desc = field.NewString(table, "desc")
	k.Lecturer = field.NewString(table, "lecturer")
	k.LecturerDesc = field.NewString(table, "lecturer_desc")
	k.Num = field.NewInt(table, "num")
	k.RequireNum = field.NewInt(table, "require_num")
	k.NotRequireNum = field.NewInt(table, "not_require_num")
	k.RecNum = field.NewFloat64(table, "rec_num")
	k.RequiredXNum = field.NewFloat64(table, "required_x_num")
	k.NotRequiredXNum = field.NewFloat64(table, "not_required_x_num")
	k.PubTime = field.NewInt(table, "pub_time")
	k.MinPrice = field.NewFloat64(table, "min_price")
	k.MaxPrice = field.NewFloat64(table, "max_price")
	k.RecPrice = field.NewFloat64(table, "rec_price")
	k.Level = field.NewInt(table, "level")
	k.Comment = field.NewInt8(table, "comment")
	k.Status = field.NewInt8(table, "status")
	k.IsPackage = field.NewInt8(table, "is_package")
	k.CoverFrom = field.NewInt8(table, "cover_from")
	k.Duration = field.NewInt(table, "duration")
	k.CreatedAt = field.NewTime(table, "created_at")
	k.UpdatedAt = field.NewTime(table, "updated_at")
	k.DeletedAt = field.NewField(table, "deleted_at")
	k.KecCode = field.NewString(table, "kec_code")

	k.fillFieldMap()

	return k
}

func (k *kec) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := k.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (k *kec) fillFieldMap() {
	k.fieldMap = make(map[string]field.Expr, 31)
	k.fieldMap["id"] = k.ID
	k.fieldMap["kec_id"] = k.KecID
	k.fieldMap["name"] = k.Name
	k.fieldMap["cover"] = k.Cover
	k.fieldMap["cate_ids"] = k.CateIds
	k.fieldMap["cate_show"] = k.CateShow
	k.fieldMap["area"] = k.Area
	k.fieldMap["seller_id"] = k.SellerID
	k.fieldMap["desc"] = k.Desc
	k.fieldMap["lecturer"] = k.Lecturer
	k.fieldMap["lecturer_desc"] = k.LecturerDesc
	k.fieldMap["num"] = k.Num
	k.fieldMap["require_num"] = k.RequireNum
	k.fieldMap["not_require_num"] = k.NotRequireNum
	k.fieldMap["rec_num"] = k.RecNum
	k.fieldMap["required_x_num"] = k.RequiredXNum
	k.fieldMap["not_required_x_num"] = k.NotRequiredXNum
	k.fieldMap["pub_time"] = k.PubTime
	k.fieldMap["min_price"] = k.MinPrice
	k.fieldMap["max_price"] = k.MaxPrice
	k.fieldMap["rec_price"] = k.RecPrice
	k.fieldMap["level"] = k.Level
	k.fieldMap["comment"] = k.Comment
	k.fieldMap["status"] = k.Status
	k.fieldMap["is_package"] = k.IsPackage
	k.fieldMap["cover_from"] = k.CoverFrom
	k.fieldMap["duration"] = k.Duration
	k.fieldMap["created_at"] = k.CreatedAt
	k.fieldMap["updated_at"] = k.UpdatedAt
	k.fieldMap["deleted_at"] = k.DeletedAt
	k.fieldMap["kec_code"] = k.KecCode
}

func (k kec) clone(db *gorm.DB) kec {
	k.kecDo.ReplaceConnPool(db.Statement.ConnPool)
	return k
}

func (k kec) replaceDB(db *gorm.DB) kec {
	k.kecDo.ReplaceDB(db)
	return k
}

type kecDo struct{ gen.DO }

type IKecDo interface {
	gen.SubQuery
	Debug() IKecDo
	WithContext(ctx context.Context) IKecDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IKecDo
	WriteDB() IKecDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IKecDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IKecDo
	Not(conds ...gen.Condition) IKecDo
	Or(conds ...gen.Condition) IKecDo
	Select(conds ...field.Expr) IKecDo
	Where(conds ...gen.Condition) IKecDo
	Order(conds ...field.Expr) IKecDo
	Distinct(cols ...field.Expr) IKecDo
	Omit(cols ...field.Expr) IKecDo
	Join(table schema.Tabler, on ...field.Expr) IKecDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IKecDo
	RightJoin(table schema.Tabler, on ...field.Expr) IKecDo
	Group(cols ...field.Expr) IKecDo
	Having(conds ...gen.Condition) IKecDo
	Limit(limit int) IKecDo
	Offset(offset int) IKecDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IKecDo
	Unscoped() IKecDo
	Create(values ...*model.Kec) error
	CreateInBatches(values []*model.Kec, batchSize int) error
	Save(values ...*model.Kec) error
	First() (*model.Kec, error)
	Take() (*model.Kec, error)
	Last() (*model.Kec, error)
	Find() ([]*model.Kec, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Kec, err error)
	FindInBatches(result *[]*model.Kec, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Kec) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IKecDo
	Assign(attrs ...field.AssignExpr) IKecDo
	Joins(fields ...field.RelationField) IKecDo
	Preload(fields ...field.RelationField) IKecDo
	FirstOrInit() (*model.Kec, error)
	FirstOrCreate() (*model.Kec, error)
	FindByPage(offset int, limit int) (result []*model.Kec, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IKecDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (k kecDo) Debug() IKecDo {
	return k.withDO(k.DO.Debug())
}

func (k kecDo) WithContext(ctx context.Context) IKecDo {
	return k.withDO(k.DO.WithContext(ctx))
}

func (k kecDo) ReadDB() IKecDo {
	return k.Clauses(dbresolver.Read)
}

func (k kecDo) WriteDB() IKecDo {
	return k.Clauses(dbresolver.Write)
}

func (k kecDo) Session(config *gorm.Session) IKecDo {
	return k.withDO(k.DO.Session(config))
}

func (k kecDo) Clauses(conds ...clause.Expression) IKecDo {
	return k.withDO(k.DO.Clauses(conds...))
}

func (k kecDo) Returning(value interface{}, columns ...string) IKecDo {
	return k.withDO(k.DO.Returning(value, columns...))
}

func (k kecDo) Not(conds ...gen.Condition) IKecDo {
	return k.withDO(k.DO.Not(conds...))
}

func (k kecDo) Or(conds ...gen.Condition) IKecDo {
	return k.withDO(k.DO.Or(conds...))
}

func (k kecDo) Select(conds ...field.Expr) IKecDo {
	return k.withDO(k.DO.Select(conds...))
}

func (k kecDo) Where(conds ...gen.Condition) IKecDo {
	return k.withDO(k.DO.Where(conds...))
}

func (k kecDo) Order(conds ...field.Expr) IKecDo {
	return k.withDO(k.DO.Order(conds...))
}

func (k kecDo) Distinct(cols ...field.Expr) IKecDo {
	return k.withDO(k.DO.Distinct(cols...))
}

func (k kecDo) Omit(cols ...field.Expr) IKecDo {
	return k.withDO(k.DO.Omit(cols...))
}

func (k kecDo) Join(table schema.Tabler, on ...field.Expr) IKecDo {
	return k.withDO(k.DO.Join(table, on...))
}

func (k kecDo) LeftJoin(table schema.Tabler, on ...field.Expr) IKecDo {
	return k.withDO(k.DO.LeftJoin(table, on...))
}

func (k kecDo) RightJoin(table schema.Tabler, on ...field.Expr) IKecDo {
	return k.withDO(k.DO.RightJoin(table, on...))
}

func (k kecDo) Group(cols ...field.Expr) IKecDo {
	return k.withDO(k.DO.Group(cols...))
}

func (k kecDo) Having(conds ...gen.Condition) IKecDo {
	return k.withDO(k.DO.Having(conds...))
}

func (k kecDo) Limit(limit int) IKecDo {
	return k.withDO(k.DO.Limit(limit))
}

func (k kecDo) Offset(offset int) IKecDo {
	return k.withDO(k.DO.Offset(offset))
}

func (k kecDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IKecDo {
	return k.withDO(k.DO.Scopes(funcs...))
}

func (k kecDo) Unscoped() IKecDo {
	return k.withDO(k.DO.Unscoped())
}

func (k kecDo) Create(values ...*model.Kec) error {
	if len(values) == 0 {
		return nil
	}
	return k.DO.Create(values)
}

func (k kecDo) CreateInBatches(values []*model.Kec, batchSize int) error {
	return k.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (k kecDo) Save(values ...*model.Kec) error {
	if len(values) == 0 {
		return nil
	}
	return k.DO.Save(values)
}

func (k kecDo) First() (*model.Kec, error) {
	if result, err := k.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Kec), nil
	}
}

func (k kecDo) Take() (*model.Kec, error) {
	if result, err := k.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Kec), nil
	}
}

func (k kecDo) Last() (*model.Kec, error) {
	if result, err := k.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Kec), nil
	}
}

func (k kecDo) Find() ([]*model.Kec, error) {
	result, err := k.DO.Find()
	return result.([]*model.Kec), err
}

func (k kecDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Kec, err error) {
	buf := make([]*model.Kec, 0, batchSize)
	err = k.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (k kecDo) FindInBatches(result *[]*model.Kec, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return k.DO.FindInBatches(result, batchSize, fc)
}

func (k kecDo) Attrs(attrs ...field.AssignExpr) IKecDo {
	return k.withDO(k.DO.Attrs(attrs...))
}

func (k kecDo) Assign(attrs ...field.AssignExpr) IKecDo {
	return k.withDO(k.DO.Assign(attrs...))
}

func (k kecDo) Joins(fields ...field.RelationField) IKecDo {
	for _, _f := range fields {
		k = *k.withDO(k.DO.Joins(_f))
	}
	return &k
}

func (k kecDo) Preload(fields ...field.RelationField) IKecDo {
	for _, _f := range fields {
		k = *k.withDO(k.DO.Preload(_f))
	}
	return &k
}

func (k kecDo) FirstOrInit() (*model.Kec, error) {
	if result, err := k.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Kec), nil
	}
}

func (k kecDo) FirstOrCreate() (*model.Kec, error) {
	if result, err := k.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Kec), nil
	}
}

func (k kecDo) FindByPage(offset int, limit int) (result []*model.Kec, count int64, err error) {
	result, err = k.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = k.Offset(-1).Limit(-1).Count()
	return
}

func (k kecDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = k.Count()
	if err != nil {
		return
	}

	err = k.Offset(offset).Limit(limit).Scan(result)
	return
}

func (k kecDo) Scan(result interface{}) (err error) {
	return k.DO.Scan(result)
}

func (k kecDo) Delete(models ...*model.Kec) (result gen.ResultInfo, err error) {
	return k.DO.Delete(models)
}

func (k *kecDo) withDO(do gen.Dao) *kecDo {
	k.DO = *do.(*gen.DO)
	return k
}
