package dbdao

import (
	"context"
	"database/sql"

	"github.com/go-xorm/xorm"
)

type Session struct {
	*xorm.Session
	engine *Engine
}

// Context sets the context on this session
func (session *Session) Context(ctx context.Context) *Session {
	session.Session = session.Session.Context(ctx)
	return session
}

// PingContext test if database is ok
func (session *Session) PingContext(ctx context.Context) error {
	return session.Session.PingContext(ctx)
}

// Exec raw sql
func (session *Session) Exec(sqlOrArgs ...interface{}) (sql.Result, error) {
	return HookExec(func() (sql.Result, error) {
		return session.Session.Exec(sqlOrArgs...)
	})
}

// Get retrieve one record from database, bean's non-empty fields
// will be as conditions
func (session *Session) Get(bean interface{}) (bool, error) {
	return HookGet(func() (bool, error) {
		return session.Session.Get(bean)
	})
}

// Find retrieve records from table, condiBeans's non-empty fields
// are conditions. beans could be []Struct, []*Struct, map[int64]Struct
// map[int64]*Struct
func (session *Session) Find(rowsSlicePtr interface{}, condiBean ...interface{}) error {

	return HookFind(func() error {
		return session.Session.Find(rowsSlicePtr, condiBean...)
	})
}

// FindAndCount find the results and also return the counts
func (session *Session) FindAndCount(rowsSlicePtr interface{}, condiBean ...interface{}) (int64, error) {
	return session.Session.FindAndCount(rowsSlicePtr, condiBean...)
}

// LastSQL returns last query information
func (session *Session) LastSQL() (string, []interface{}) {
	return session.Session.LastSQL()
}

// Insert insert one or more beans
func (session *Session) Insert(beans ...interface{}) (int64, error) {
	return HookInsert(func() (int64, error) {
		return session.Session.Insert(beans...)
	})
}

// Delete records, bean's non-empty fields are conditions
func (session *Session) Delete(bean interface{}) (int64, error) {
	return HookDelete(func() (int64, error) {
		return session.Session.Delete(bean)
	})
}

// InsertMulti insert multiple records
func (session *Session) InsertMulti(rowsSlicePtr interface{}) (int64, error) {
	return session.Session.InsertMulti(rowsSlicePtr)
}

// InsertOne insert only one struct into database as a record.
// The in parameter bean must a struct or a point to struct. The return
// parameter is inserted and error
func (session *Session) InsertOne(bean interface{}) (int64, error) {
	return session.Session.InsertOne(bean)
}

// Query runs a raw sql and return records as []map[string][]byte
func (session *Session) Query(sqlOrArgs ...interface{}) ([]map[string][]byte, error) {
	return HookQuery(func() ([]map[string][]byte, error) {
		return session.Session.Query(sqlOrArgs...)
	})
}

// QueryString runs a raw sql and return records as []map[string]string
func (session *Session) QueryString(sqlOrArgs ...interface{}) ([]map[string]string, error) {
	return session.Session.QueryString(sqlOrArgs...)
}

// QuerySliceString runs a raw sql and return records as [][]string
func (session *Session) QuerySliceString(sqlOrArgs ...interface{}) ([][]string, error) {
	return session.Session.QuerySliceString(sqlOrArgs...)
}

// QueryInterface runs a raw sql and return records as []map[string]interface{}
func (session *Session) QueryInterface(sqlOrArgs ...interface{}) ([]map[string]interface{}, error) {
	return session.Session.QueryInterface(sqlOrArgs...)
}

// Update records, bean's non-empty fields are updated contents,
// condiBean' non-empty filds are conditions
// CAUTION:
//        1.bool will defaultly be updated content nor conditions
//         You should call UseBool if you have bool to use.
//        2.float32 & float64 may be not inexact as conditions
func (session *Session) Update(bean interface{}, condiBean ...interface{}) (int64, error) {
	return HookUpdate(func() (int64, error) {
		return session.Session.Update(bean, condiBean...)
	})
}

// Unscoped always disable struct tag "deleted"
func (session *Session) Unscoped() *Session {
	session.Session = session.Session.Unscoped()
	return session
}

// Prepare set a flag to session that should be prepare statement before execute query
func (session *Session) Prepare() *Session {
	session.Session = session.Session.Prepare()
	return session
}

// Before Apply before Processor, affected bean is passed to closure arg
func (session *Session) Before(closures func(interface{})) *Session {
	session.Session = session.Session.Before(closures)
	return session
}

// After Apply after Processor, affected bean is passed to closure arg
func (session *Session) After(closures func(interface{})) *Session {
	session.Session = session.Session.After(closures)
	return session
}

// Table can input a string or pointer to struct for special a table to operate.
func (session *Session) Table(tableNameOrBean interface{}) *Session {
	session.Session = session.Session.Table(tableNameOrBean)
	return session
}

// Alias set the table alias
func (session *Session) Alias(alias string) *Session {
	session.Session = session.Session.Alias(alias)
	return session
}

// NoAutoCondition disable generate SQL condition from beans
func (session *Session) NoAutoCondition(no ...bool) *Session {
	session.Session = session.Session.NoAutoCondition(no...)
	return session
}

// Limit provide limit and offset query condition
func (session *Session) Limit(limit int, start ...int) *Session {
	session.Session = session.Session.Limit(limit, start...)
	return session
}

// OrderBy provide order by query condition, the input parameter is the content
// after order by on a sql statement.
func (session *Session) OrderBy(order string) *Session {
	session.Session = session.Session.OrderBy(order)
	return session
}

// Desc provide desc order by query condition, the input parameters are columns.
func (session *Session) Desc(colNames ...string) *Session {
	session.Session = session.Session.Desc(colNames...)
	return session
}

// Asc provide asc order by query condition, the input parameters are columns.
func (session *Session) Asc(colNames ...string) *Session {
	session.Session = session.Session.Asc(colNames...)
	return session
}

// StoreEngine is only avialble mysql dialect currently
func (session *Session) StoreEngine(storeEngine string) *Session {
	session.Session = session.Session.StoreEngine(storeEngine)
	return session
}

// Charset is only avialble mysql dialect currently
func (session *Session) Charset(charset string) *Session {
	session.Session = session.Session.Charset(charset)
	return session
}

// NoCascade indicate that no cascade load child object
func (session *Session) NoCascade() *Session {
	session.Session = session.Session.NoCascade()
	return session
}

// ForUpdate Set Read/Write locking for UPDATE
func (session *Session) ForUpdate() *Session {
	session.Session = session.Session.ForUpdate()
	return session
}

// Cascade indicates if loading sub Struct
func (session *Session) Cascade(trueOrFalse ...bool) *Session {
	session.Session = session.Session.Cascade(trueOrFalse...)
	return session
}

// NoCache ask this session do not retrieve data from cache system and
// get data from database directly.
func (session *Session) NoCache() *Session {
	session.Session = session.Session.NoCache()
	return session
}

// Join join_operator should be one of INNER, LEFT OUTER, CROSS etc - this will be prepended to JOIN
func (session *Session) Join(joinOperator string, tablename interface{}, condition string, args ...interface{}) *Session {
	session.Session = session.Session.Join(joinOperator, tablename, condition, args...)
	return session
}

// GroupBy Generate Group By statement
func (session *Session) GroupBy(keys string) *Session {
	session.Session = session.Session.GroupBy(keys)
	return session
}

// Having Generate Having statement
func (session *Session) Having(conditions string) *Session {
	session.Session = session.Session.Having(conditions)
	return session
}

// Incr provides a query string like "count = count + 1"
func (session *Session) Incr(column string, arg ...interface{}) *Session {
	session.Session = session.Session.Incr(column, arg...)
	return session
}

// Decr provides a query string like "count = count - 1"
func (session *Session) Decr(column string, arg ...interface{}) *Session {
	session.Session = session.Session.Decr(column, arg...)
	return session
}

// SetExpr provides a query string like "column = {expression}"
func (session *Session) SetExpr(column string, expression interface{}) *Session {
	session.Session = session.Session.SetExpr(column, expression)
	return session
}

// Select provides some columns to special
func (session *Session) Select(str string) *Session {
	session.Session = session.Session.Select(str)
	return session
}

// Cols provides some columns to special
func (session *Session) Cols(columns ...string) *Session {
	session.Session = session.Session.Cols(columns...)
	return session
}

// AllCols ask all columns
func (session *Session) AllCols() *Session {
	session.Session = session.Session.AllCols()
	return session
}

// MustCols specify some columns must use even if they are empty
func (session *Session) MustCols(columns ...string) *Session {
	session.Session = session.Session.MustCols(columns...)
	return session
}

// UseBool automatically retrieve condition according struct, but
// if struct has bool field, it will ignore them. So use UseBool
// to tell system to do not ignore them.
// If no parameters, it will use all the bool field of struct, or
// it will use parameters's columns
func (session *Session) UseBool(columns ...string) *Session {
	session.Session = session.Session.UseBool(columns...)
	return session
}

// Distinct use for distinct columns. Caution: when you are using cache,
// distinct will not be cached because cache system need id,
// but distinct will not provide id
func (session *Session) Distinct(columns ...string) *Session {
	session.Session = session.Session.Distinct(columns...)
	return session
}

// Omit Only not use the parameters as select or update columns
func (session *Session) Omit(columns ...string) *Session {
	session.Session = session.Session.Omit(columns...)
	return session
}

// Nullable Set null when column is zero-value and nullable for update
func (session *Session) Nullable(columns ...string) *Session {
	session.Session = session.Session.Nullable(columns...)
	return session
}

// NoAutoTime means do not automatically give created field and updated field
// the current time on the current session temporarily
func (session *Session) NoAutoTime() *Session {
	session.Session = session.Session.NoAutoTime()
	return session
}

// Sql provides raw sql input parameter. When you have a complex SQL statement
// and cannot use Where, Id, In and etc. Methods to describe, you can use SQL.
//
// Deprecated: use SQL instead.
func (session *Session) Sql(query string, args ...interface{}) *Session {
	session.Session = session.Session.Sql(query, args...)
	return session
}

// SQL provides raw sql input parameter. When you have a complex SQL statement
// and cannot use Where, Id, In and etc. Methods to describe, you can use SQL.
func (session *Session) SQL(query interface{}, args ...interface{}) *Session {
	session.Session = session.Session.SQL(query, args...)
	return session
}

// Where provides custom query condition.
func (session *Session) Where(query interface{}, args ...interface{}) *Session {
	session.Session = session.Session.Where(query, args...)
	return session
}

// And provides custom query condition.
func (session *Session) And(query interface{}, args ...interface{}) *Session {
	session.Session = session.Session.And(query, args...)
	return session
}

// Or provides custom query condition.
func (session *Session) Or(query interface{}, args ...interface{}) *Session {
	session.Session = session.Session.Or(query, args...)
	return session
}

// Id provides converting id as a query condition
//
// Deprecated: use ID instead
func (session *Session) Id(id interface{}) *Session {
	session.Session = session.Session.ID(id)
	return session
}

// ID provides converting id as a query condition
func (session *Session) ID(id interface{}) *Session {
	session.Session = session.Session.ID(id)
	return session
}

// In provides a query string like "id in (1, 2, 3)"
func (session *Session) In(column string, args ...interface{}) *Session {
	session.Session = session.Session.In(column, args...)
	return session
}

// NotIn provides a query string like "id in (1, 2, 3)"
func (session *Session) NotIn(column string, args ...interface{}) *Session {
	session.Session = session.Session.NotIn(column, args...)
	return session
}
