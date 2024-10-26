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

func newSpot(db *gorm.DB, opts ...gen.DOOption) spot {
	_spot := spot{}

	_spot.spotDo.UseDB(db, opts...)
	_spot.spotDo.UseModel(&domain.Spot{})

	tableName := _spot.spotDo.TableName()
	_spot.ALL = field.NewAsterisk(tableName)
	_spot.ID = field.NewString(tableName, "id")
	_spot.GmID = field.NewString(tableName, "gm_id")
	_spot.CreatedAt = field.NewTime(tableName, "created_at")
	_spot.UpdatedAt = field.NewTime(tableName, "updated_at")

	_spot.fillFieldMap()

	return _spot
}

type spot struct {
	spotDo

	ALL       field.Asterisk
	ID        field.String
	GmID      field.String
	CreatedAt field.Time
	UpdatedAt field.Time

	fieldMap map[string]field.Expr
}

func (s spot) Table(newTableName string) *spot {
	s.spotDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s spot) As(alias string) *spot {
	s.spotDo.DO = *(s.spotDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *spot) updateTableName(table string) *spot {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewString(table, "id")
	s.GmID = field.NewString(table, "gm_id")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")

	s.fillFieldMap()

	return s
}

func (s *spot) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *spot) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 4)
	s.fieldMap["id"] = s.ID
	s.fieldMap["gm_id"] = s.GmID
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
}

func (s spot) clone(db *gorm.DB) spot {
	s.spotDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s spot) replaceDB(db *gorm.DB) spot {
	s.spotDo.ReplaceDB(db)
	return s
}

type spotDo struct{ gen.DO }

type ISpotDo interface {
	gen.SubQuery
	Debug() ISpotDo
	WithContext(ctx context.Context) ISpotDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISpotDo
	WriteDB() ISpotDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISpotDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISpotDo
	Not(conds ...gen.Condition) ISpotDo
	Or(conds ...gen.Condition) ISpotDo
	Select(conds ...field.Expr) ISpotDo
	Where(conds ...gen.Condition) ISpotDo
	Order(conds ...field.Expr) ISpotDo
	Distinct(cols ...field.Expr) ISpotDo
	Omit(cols ...field.Expr) ISpotDo
	Join(table schema.Tabler, on ...field.Expr) ISpotDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISpotDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISpotDo
	Group(cols ...field.Expr) ISpotDo
	Having(conds ...gen.Condition) ISpotDo
	Limit(limit int) ISpotDo
	Offset(offset int) ISpotDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISpotDo
	Unscoped() ISpotDo
	Create(values ...*domain.Spot) error
	CreateInBatches(values []*domain.Spot, batchSize int) error
	Save(values ...*domain.Spot) error
	First() (*domain.Spot, error)
	Take() (*domain.Spot, error)
	Last() (*domain.Spot, error)
	Find() ([]*domain.Spot, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*domain.Spot, err error)
	FindInBatches(result *[]*domain.Spot, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*domain.Spot) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISpotDo
	Assign(attrs ...field.AssignExpr) ISpotDo
	Joins(fields ...field.RelationField) ISpotDo
	Preload(fields ...field.RelationField) ISpotDo
	FirstOrInit() (*domain.Spot, error)
	FirstOrCreate() (*domain.Spot, error)
	FindByPage(offset int, limit int) (result []*domain.Spot, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISpotDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s spotDo) Debug() ISpotDo {
	return s.withDO(s.DO.Debug())
}

func (s spotDo) WithContext(ctx context.Context) ISpotDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s spotDo) ReadDB() ISpotDo {
	return s.Clauses(dbresolver.Read)
}

func (s spotDo) WriteDB() ISpotDo {
	return s.Clauses(dbresolver.Write)
}

func (s spotDo) Session(config *gorm.Session) ISpotDo {
	return s.withDO(s.DO.Session(config))
}

func (s spotDo) Clauses(conds ...clause.Expression) ISpotDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s spotDo) Returning(value interface{}, columns ...string) ISpotDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s spotDo) Not(conds ...gen.Condition) ISpotDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s spotDo) Or(conds ...gen.Condition) ISpotDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s spotDo) Select(conds ...field.Expr) ISpotDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s spotDo) Where(conds ...gen.Condition) ISpotDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s spotDo) Order(conds ...field.Expr) ISpotDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s spotDo) Distinct(cols ...field.Expr) ISpotDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s spotDo) Omit(cols ...field.Expr) ISpotDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s spotDo) Join(table schema.Tabler, on ...field.Expr) ISpotDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s spotDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISpotDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s spotDo) RightJoin(table schema.Tabler, on ...field.Expr) ISpotDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s spotDo) Group(cols ...field.Expr) ISpotDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s spotDo) Having(conds ...gen.Condition) ISpotDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s spotDo) Limit(limit int) ISpotDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s spotDo) Offset(offset int) ISpotDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s spotDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISpotDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s spotDo) Unscoped() ISpotDo {
	return s.withDO(s.DO.Unscoped())
}

func (s spotDo) Create(values ...*domain.Spot) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s spotDo) CreateInBatches(values []*domain.Spot, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s spotDo) Save(values ...*domain.Spot) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s spotDo) First() (*domain.Spot, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*domain.Spot), nil
	}
}

func (s spotDo) Take() (*domain.Spot, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*domain.Spot), nil
	}
}

func (s spotDo) Last() (*domain.Spot, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*domain.Spot), nil
	}
}

func (s spotDo) Find() ([]*domain.Spot, error) {
	result, err := s.DO.Find()
	return result.([]*domain.Spot), err
}

func (s spotDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*domain.Spot, err error) {
	buf := make([]*domain.Spot, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s spotDo) FindInBatches(result *[]*domain.Spot, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s spotDo) Attrs(attrs ...field.AssignExpr) ISpotDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s spotDo) Assign(attrs ...field.AssignExpr) ISpotDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s spotDo) Joins(fields ...field.RelationField) ISpotDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s spotDo) Preload(fields ...field.RelationField) ISpotDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s spotDo) FirstOrInit() (*domain.Spot, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*domain.Spot), nil
	}
}

func (s spotDo) FirstOrCreate() (*domain.Spot, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*domain.Spot), nil
	}
}

func (s spotDo) FindByPage(offset int, limit int) (result []*domain.Spot, count int64, err error) {
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

func (s spotDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s spotDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s spotDo) Delete(models ...*domain.Spot) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *spotDo) withDO(do gen.Dao) *spotDo {
	s.DO = *do.(*gen.DO)
	return s
}
