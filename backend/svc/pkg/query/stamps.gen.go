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

	"backend/svc/pkg/domain"
)

func newStamp(db *gorm.DB, opts ...gen.DOOption) stamp {
	_stamp := stamp{}

	_stamp.stampDo.UseDB(db, opts...)
	_stamp.stampDo.UseModel(&domain.Stamp{})

	tableName := _stamp.stampDo.TableName()
	_stamp.ALL = field.NewAsterisk(tableName)
	_stamp.Type = field.NewString(tableName, "type")
	_stamp.CreatedAt = field.NewTime(tableName, "created_at")
	_stamp.UpdatedAt = field.NewTime(tableName, "updated_at")

	_stamp.fillFieldMap()

	return _stamp
}

type stamp struct {
	stampDo

	ALL       field.Asterisk
	Type      field.String
	CreatedAt field.Time
	UpdatedAt field.Time

	fieldMap map[string]field.Expr
}

func (s stamp) Table(newTableName string) *stamp {
	s.stampDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s stamp) As(alias string) *stamp {
	s.stampDo.DO = *(s.stampDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *stamp) updateTableName(table string) *stamp {
	s.ALL = field.NewAsterisk(table)
	s.Type = field.NewString(table, "type")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")

	s.fillFieldMap()

	return s
}

func (s *stamp) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *stamp) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 3)
	s.fieldMap["type"] = s.Type
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
}

func (s stamp) clone(db *gorm.DB) stamp {
	s.stampDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s stamp) replaceDB(db *gorm.DB) stamp {
	s.stampDo.ReplaceDB(db)
	return s
}

type stampDo struct{ gen.DO }

type IStampDo interface {
	gen.SubQuery
	Debug() IStampDo
	WithContext(ctx context.Context) IStampDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IStampDo
	WriteDB() IStampDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IStampDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IStampDo
	Not(conds ...gen.Condition) IStampDo
	Or(conds ...gen.Condition) IStampDo
	Select(conds ...field.Expr) IStampDo
	Where(conds ...gen.Condition) IStampDo
	Order(conds ...field.Expr) IStampDo
	Distinct(cols ...field.Expr) IStampDo
	Omit(cols ...field.Expr) IStampDo
	Join(table schema.Tabler, on ...field.Expr) IStampDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IStampDo
	RightJoin(table schema.Tabler, on ...field.Expr) IStampDo
	Group(cols ...field.Expr) IStampDo
	Having(conds ...gen.Condition) IStampDo
	Limit(limit int) IStampDo
	Offset(offset int) IStampDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IStampDo
	Unscoped() IStampDo
	Create(values ...*domain.Stamp) error
	CreateInBatches(values []*domain.Stamp, batchSize int) error
	Save(values ...*domain.Stamp) error
	First() (*domain.Stamp, error)
	Take() (*domain.Stamp, error)
	Last() (*domain.Stamp, error)
	Find() ([]*domain.Stamp, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*domain.Stamp, err error)
	FindInBatches(result *[]*domain.Stamp, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*domain.Stamp) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IStampDo
	Assign(attrs ...field.AssignExpr) IStampDo
	Joins(fields ...field.RelationField) IStampDo
	Preload(fields ...field.RelationField) IStampDo
	FirstOrInit() (*domain.Stamp, error)
	FirstOrCreate() (*domain.Stamp, error)
	FindByPage(offset int, limit int) (result []*domain.Stamp, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IStampDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s stampDo) Debug() IStampDo {
	return s.withDO(s.DO.Debug())
}

func (s stampDo) WithContext(ctx context.Context) IStampDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s stampDo) ReadDB() IStampDo {
	return s.Clauses(dbresolver.Read)
}

func (s stampDo) WriteDB() IStampDo {
	return s.Clauses(dbresolver.Write)
}

func (s stampDo) Session(config *gorm.Session) IStampDo {
	return s.withDO(s.DO.Session(config))
}

func (s stampDo) Clauses(conds ...clause.Expression) IStampDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s stampDo) Returning(value interface{}, columns ...string) IStampDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s stampDo) Not(conds ...gen.Condition) IStampDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s stampDo) Or(conds ...gen.Condition) IStampDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s stampDo) Select(conds ...field.Expr) IStampDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s stampDo) Where(conds ...gen.Condition) IStampDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s stampDo) Order(conds ...field.Expr) IStampDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s stampDo) Distinct(cols ...field.Expr) IStampDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s stampDo) Omit(cols ...field.Expr) IStampDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s stampDo) Join(table schema.Tabler, on ...field.Expr) IStampDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s stampDo) LeftJoin(table schema.Tabler, on ...field.Expr) IStampDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s stampDo) RightJoin(table schema.Tabler, on ...field.Expr) IStampDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s stampDo) Group(cols ...field.Expr) IStampDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s stampDo) Having(conds ...gen.Condition) IStampDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s stampDo) Limit(limit int) IStampDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s stampDo) Offset(offset int) IStampDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s stampDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IStampDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s stampDo) Unscoped() IStampDo {
	return s.withDO(s.DO.Unscoped())
}

func (s stampDo) Create(values ...*domain.Stamp) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s stampDo) CreateInBatches(values []*domain.Stamp, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s stampDo) Save(values ...*domain.Stamp) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s stampDo) First() (*domain.Stamp, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*domain.Stamp), nil
	}
}

func (s stampDo) Take() (*domain.Stamp, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*domain.Stamp), nil
	}
}

func (s stampDo) Last() (*domain.Stamp, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*domain.Stamp), nil
	}
}

func (s stampDo) Find() ([]*domain.Stamp, error) {
	result, err := s.DO.Find()
	return result.([]*domain.Stamp), err
}

func (s stampDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*domain.Stamp, err error) {
	buf := make([]*domain.Stamp, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s stampDo) FindInBatches(result *[]*domain.Stamp, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s stampDo) Attrs(attrs ...field.AssignExpr) IStampDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s stampDo) Assign(attrs ...field.AssignExpr) IStampDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s stampDo) Joins(fields ...field.RelationField) IStampDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s stampDo) Preload(fields ...field.RelationField) IStampDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s stampDo) FirstOrInit() (*domain.Stamp, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*domain.Stamp), nil
	}
}

func (s stampDo) FirstOrCreate() (*domain.Stamp, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*domain.Stamp), nil
	}
}

func (s stampDo) FindByPage(offset int, limit int) (result []*domain.Stamp, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s stampDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s stampDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s stampDo) Delete(models ...*domain.Stamp) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *stampDo) withDO(do gen.Dao) *stampDo {
	s.DO = *do.(*gen.DO)
	return s
}
