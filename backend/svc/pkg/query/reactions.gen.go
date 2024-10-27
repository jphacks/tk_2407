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

func newReaction(db *gorm.DB, opts ...gen.DOOption) reaction {
	_reaction := reaction{}

	_reaction.reactionDo.UseDB(db, opts...)
	_reaction.reactionDo.UseModel(&domain.Reaction{})

	tableName := _reaction.reactionDo.TableName()
	_reaction.ALL = field.NewAsterisk(tableName)
	_reaction.ID = field.NewString(tableName, "id")
	_reaction.UserID = field.NewString(tableName, "user_id")
	_reaction.MessageID = field.NewString(tableName, "message_id")
	_reaction.CreatedAt = field.NewTime(tableName, "created_at")
	_reaction.StampType = field.NewString(tableName, "stamp_type")

	_reaction.fillFieldMap()

	return _reaction
}

type reaction struct {
	reactionDo

	ALL       field.Asterisk
	ID        field.String
	UserID    field.String
	MessageID field.String
	CreatedAt field.Time
	StampType field.String

	fieldMap map[string]field.Expr
}

func (r reaction) Table(newTableName string) *reaction {
	r.reactionDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r reaction) As(alias string) *reaction {
	r.reactionDo.DO = *(r.reactionDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *reaction) updateTableName(table string) *reaction {
	r.ALL = field.NewAsterisk(table)
	r.ID = field.NewString(table, "id")
	r.UserID = field.NewString(table, "user_id")
	r.MessageID = field.NewString(table, "message_id")
	r.CreatedAt = field.NewTime(table, "created_at")
	r.StampType = field.NewString(table, "stamp_type")

	r.fillFieldMap()

	return r
}

func (r *reaction) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *reaction) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 5)
	r.fieldMap["id"] = r.ID
	r.fieldMap["user_id"] = r.UserID
	r.fieldMap["message_id"] = r.MessageID
	r.fieldMap["created_at"] = r.CreatedAt
	r.fieldMap["stamp_type"] = r.StampType
}

func (r reaction) clone(db *gorm.DB) reaction {
	r.reactionDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r reaction) replaceDB(db *gorm.DB) reaction {
	r.reactionDo.ReplaceDB(db)
	return r
}

type reactionDo struct{ gen.DO }

type IReactionDo interface {
	gen.SubQuery
	Debug() IReactionDo
	WithContext(ctx context.Context) IReactionDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IReactionDo
	WriteDB() IReactionDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IReactionDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IReactionDo
	Not(conds ...gen.Condition) IReactionDo
	Or(conds ...gen.Condition) IReactionDo
	Select(conds ...field.Expr) IReactionDo
	Where(conds ...gen.Condition) IReactionDo
	Order(conds ...field.Expr) IReactionDo
	Distinct(cols ...field.Expr) IReactionDo
	Omit(cols ...field.Expr) IReactionDo
	Join(table schema.Tabler, on ...field.Expr) IReactionDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IReactionDo
	RightJoin(table schema.Tabler, on ...field.Expr) IReactionDo
	Group(cols ...field.Expr) IReactionDo
	Having(conds ...gen.Condition) IReactionDo
	Limit(limit int) IReactionDo
	Offset(offset int) IReactionDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IReactionDo
	Unscoped() IReactionDo
	Create(values ...*domain.Reaction) error
	CreateInBatches(values []*domain.Reaction, batchSize int) error
	Save(values ...*domain.Reaction) error
	First() (*domain.Reaction, error)
	Take() (*domain.Reaction, error)
	Last() (*domain.Reaction, error)
	Find() ([]*domain.Reaction, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*domain.Reaction, err error)
	FindInBatches(result *[]*domain.Reaction, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*domain.Reaction) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IReactionDo
	Assign(attrs ...field.AssignExpr) IReactionDo
	Joins(fields ...field.RelationField) IReactionDo
	Preload(fields ...field.RelationField) IReactionDo
	FirstOrInit() (*domain.Reaction, error)
	FirstOrCreate() (*domain.Reaction, error)
	FindByPage(offset int, limit int) (result []*domain.Reaction, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IReactionDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (r reactionDo) Debug() IReactionDo {
	return r.withDO(r.DO.Debug())
}

func (r reactionDo) WithContext(ctx context.Context) IReactionDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r reactionDo) ReadDB() IReactionDo {
	return r.Clauses(dbresolver.Read)
}

func (r reactionDo) WriteDB() IReactionDo {
	return r.Clauses(dbresolver.Write)
}

func (r reactionDo) Session(config *gorm.Session) IReactionDo {
	return r.withDO(r.DO.Session(config))
}

func (r reactionDo) Clauses(conds ...clause.Expression) IReactionDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r reactionDo) Returning(value interface{}, columns ...string) IReactionDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r reactionDo) Not(conds ...gen.Condition) IReactionDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r reactionDo) Or(conds ...gen.Condition) IReactionDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r reactionDo) Select(conds ...field.Expr) IReactionDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r reactionDo) Where(conds ...gen.Condition) IReactionDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r reactionDo) Order(conds ...field.Expr) IReactionDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r reactionDo) Distinct(cols ...field.Expr) IReactionDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r reactionDo) Omit(cols ...field.Expr) IReactionDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r reactionDo) Join(table schema.Tabler, on ...field.Expr) IReactionDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r reactionDo) LeftJoin(table schema.Tabler, on ...field.Expr) IReactionDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r reactionDo) RightJoin(table schema.Tabler, on ...field.Expr) IReactionDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r reactionDo) Group(cols ...field.Expr) IReactionDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r reactionDo) Having(conds ...gen.Condition) IReactionDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r reactionDo) Limit(limit int) IReactionDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r reactionDo) Offset(offset int) IReactionDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r reactionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IReactionDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r reactionDo) Unscoped() IReactionDo {
	return r.withDO(r.DO.Unscoped())
}

func (r reactionDo) Create(values ...*domain.Reaction) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r reactionDo) CreateInBatches(values []*domain.Reaction, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r reactionDo) Save(values ...*domain.Reaction) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r reactionDo) First() (*domain.Reaction, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*domain.Reaction), nil
	}
}

func (r reactionDo) Take() (*domain.Reaction, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*domain.Reaction), nil
	}
}

func (r reactionDo) Last() (*domain.Reaction, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*domain.Reaction), nil
	}
}

func (r reactionDo) Find() ([]*domain.Reaction, error) {
	result, err := r.DO.Find()
	return result.([]*domain.Reaction), err
}

func (r reactionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*domain.Reaction, err error) {
	buf := make([]*domain.Reaction, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r reactionDo) FindInBatches(result *[]*domain.Reaction, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r reactionDo) Attrs(attrs ...field.AssignExpr) IReactionDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r reactionDo) Assign(attrs ...field.AssignExpr) IReactionDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r reactionDo) Joins(fields ...field.RelationField) IReactionDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r reactionDo) Preload(fields ...field.RelationField) IReactionDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r reactionDo) FirstOrInit() (*domain.Reaction, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*domain.Reaction), nil
	}
}

func (r reactionDo) FirstOrCreate() (*domain.Reaction, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*domain.Reaction), nil
	}
}

func (r reactionDo) FindByPage(offset int, limit int) (result []*domain.Reaction, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r reactionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r reactionDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r reactionDo) Delete(models ...*domain.Reaction) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *reactionDo) withDO(do gen.Dao) *reactionDo {
	r.DO = *do.(*gen.DO)
	return r
}
