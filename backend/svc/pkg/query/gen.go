// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q               = new(Query)
	GmPlace         *gmPlace
	GmPlacePhoto    *gmPlacePhoto
	Message         *message
	Reaction        *reaction
	SchemaMigration *schemaMigration
	Stamp           *stamp
	Test            *test
	User            *user
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	GmPlace = &Q.GmPlace
	GmPlacePhoto = &Q.GmPlacePhoto
	Message = &Q.Message
	Reaction = &Q.Reaction
	SchemaMigration = &Q.SchemaMigration
	Stamp = &Q.Stamp
	Test = &Q.Test
	User = &Q.User
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:              db,
		GmPlace:         newGmPlace(db, opts...),
		GmPlacePhoto:    newGmPlacePhoto(db, opts...),
		Message:         newMessage(db, opts...),
		Reaction:        newReaction(db, opts...),
		SchemaMigration: newSchemaMigration(db, opts...),
		Stamp:           newStamp(db, opts...),
		Test:            newTest(db, opts...),
		User:            newUser(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	GmPlace         gmPlace
	GmPlacePhoto    gmPlacePhoto
	Message         message
	Reaction        reaction
	SchemaMigration schemaMigration
	Stamp           stamp
	Test            test
	User            user
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:              db,
		GmPlace:         q.GmPlace.clone(db),
		GmPlacePhoto:    q.GmPlacePhoto.clone(db),
		Message:         q.Message.clone(db),
		Reaction:        q.Reaction.clone(db),
		SchemaMigration: q.SchemaMigration.clone(db),
		Stamp:           q.Stamp.clone(db),
		Test:            q.Test.clone(db),
		User:            q.User.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:              db,
		GmPlace:         q.GmPlace.replaceDB(db),
		GmPlacePhoto:    q.GmPlacePhoto.replaceDB(db),
		Message:         q.Message.replaceDB(db),
		Reaction:        q.Reaction.replaceDB(db),
		SchemaMigration: q.SchemaMigration.replaceDB(db),
		Stamp:           q.Stamp.replaceDB(db),
		Test:            q.Test.replaceDB(db),
		User:            q.User.replaceDB(db),
	}
}

type queryCtx struct {
	GmPlace         IGmPlaceDo
	GmPlacePhoto    IGmPlacePhotoDo
	Message         IMessageDo
	Reaction        IReactionDo
	SchemaMigration ISchemaMigrationDo
	Stamp           IStampDo
	Test            ITestDo
	User            IUserDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		GmPlace:         q.GmPlace.WithContext(ctx),
		GmPlacePhoto:    q.GmPlacePhoto.WithContext(ctx),
		Message:         q.Message.WithContext(ctx),
		Reaction:        q.Reaction.WithContext(ctx),
		SchemaMigration: q.SchemaMigration.WithContext(ctx),
		Stamp:           q.Stamp.WithContext(ctx),
		Test:            q.Test.WithContext(ctx),
		User:            q.User.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
